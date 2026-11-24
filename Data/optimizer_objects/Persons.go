package optimizerobjects

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	G "optimizer/cmd/globals"
)

type ValReturn int

const (
	val_ok     ValReturn = 1
	val_error  ValReturn = 0
	val_remove ValReturn = -1
)

// type Schedule struct {
// 	People []Person
// }

func MakePerson(name string, days [7]Day, restrictions []P_Restr_impl) Person {
	p := Person{
		id:             name,
		days:           days,
		in_schedule:    true,
		restrictions:   restrictions,
		badness_score:  0,
		available_days: []*Day{},
	}
	p.populate_available_days()
	p.validate()
	return p
}

type Person struct {
	id             string
	days           [7]Day
	in_schedule    bool
	restrictions   []P_Restr_impl
	badness_score  float32 // only update on change
	available_days []*Day
	// days_working_sheduled [7]int // Restriction_working_days
}

func (p Person) validate() (bool, string) {
	slog.Info(fmt.Sprintf("validating Person %v", p.id))
	for _, r := range p.restrictions {
		if e, s := r.validate(); e != true {
			panic(fmt.Sprintf("!Error for %v: %v", p.id, s))
		}
	}
	for _, d := range p.days {
		if e, s := d.validate(); e != true {
			panic(fmt.Sprintf("!Error for %v: %v", p.id, s))
		}
	}

	return true, ""
}

func (p *Person) populate_available_days() {
	p.available_days = []*Day{}
	for i := range p.days {
		if p.days[i].available {
			p.available_days = append(p.available_days, &p.days[i])
		}
	}
}

// returns the Person restriction score + the individual day scores
func (p Person) GetScore() float32 {
	var tot_badness_score float32 = 0
	tot_badness_score += p.badness_score
	for _, day := range p.days {
		if day.available {
			p.badness_score += day.GetScore() // each object ony carries it's own tot_bad_score
		}
	}
	return tot_badness_score
}

func (p *Person) Calculate_badness() {
	slog.Debug(fmt.Sprintf(G.CPurple+"calculating bad_score for"+G.Bold+" person %v", p.id))
	// calc for person restrictions
	p.badness_score = 0
	for _, r := range p.restrictions {
		r.calculate_badness(*p)
		p.badness_score += r.GetScore()
	}
	slog.Debug(G.Bold+G.CCyan+"Score person %v == %v", p.id, p.badness_score)
	// also calc for days subobjects
	for _, day := range p.days {
		if day.available {
			day.Calculate_badness()
			p.badness_score += day.GetScore() // each object ony carries it's own bad_score
		}
	}
}

// could turn into hours per day later
func (p Person) get_working_days() int {
	var tot = 0
	for _, day := range p.days {
		if day.current_hours > 0 {
			tot += 1 //bool_to_int(day.active)
		}
	}
	return tot
}

// MOVES

func (p *Person) Make_change() {
	change_nr := rand.IntN(len(person_available_changes))
	person_available_changes[change_nr](p)
	slog.Debug(fmt.Sprintf("making"+G.Underline+"change %v"+G.Reset+" on Person %v", change_nr, p.id))
	p.Calculate_badness()
}

func (p *Person) change_day() {

	if len(p.available_days) < 1 {
		panic(fmt.Sprintf("Person %v has no available days!", p.id))
	}

	day_change := p.available_days[rand.IntN(len(p.days))]
	slog.Debug(fmt.Sprintf(G.Bold+"Changing day %v"+G.Reset+" for person %v", day_change.nr, p.id))

	day_change.Make_change()

}
func (p *Person) change_day_availability() {
	day := &p.days[rand.IntN(len(p.days))]
	slog.Debug(fmt.Sprintf(G.Bold+"Changing day %v"+G.Reset+" for person %v to AVAILABLE", day.nr, p.id))
	day.available = true
	p.populate_available_days()
}

var person_available_changes = []func(p *Person){
	(*Person).change_day,
	(*Person).change_day_availability,
}
