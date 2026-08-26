package optimizerobjects

import (
	"fmt"
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

func (p Person) calculate_badness() float32 {
	p.badness_score = 0
	for _, r := range p.restrictions {
		p.badness_score += r.badness_score(p)
	}
	for _, day := range p.days {
		if day.active {
			p.badness_score += day.badness_score
		}
	}
	return p.badness_score
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
