package optimizerobjects

import (
	"fmt"
	"log"
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
	return Person{
		id:            name,
		days:          days,
		in_schedule:   true,
		restrictions:  restrictions,
		badness_score: 0,
	}

}

type Person struct {
	id            string
	days          [7]Day
	in_schedule   bool
	restrictions  []P_Restr_impl
	badness_score float32 // only update on change
	// days_working_sheduled [7]int // Restriction_working_days
}

func (p Person) validate() (ValReturn, string) {
	for _, r := range p.restrictions {
		if e, s := r.validate(); e != true {
			panic(fmt.Sprintf("!Error for %v: %v", p.id, s))
		}
	}

	return val_ok, ""
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
	log.Printf(G.CPurple+"calculating bad_score for"+G.Bold+" person %v", p.id)
	// calc for person restrictions
	p.badness_score = 0
	for _, r := range p.restrictions {
		r.calculate_badness(*p)
		p.badness_score += r.GetScore()
	}
	log.Printf(G.Bold+G.CCyan+"Score person %v == %v", p.id, p.badness_score)
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

func (p Person) change_day() {
	log.Printf(G.Bold+"Changing day"+G.Reset+" for person %v", p.id)
	panic("Not implemented")
}
