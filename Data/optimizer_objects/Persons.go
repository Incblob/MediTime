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

type Schedule struct {
	People []Person
}

type Person struct {
	id            string
	days          [7]Day
	in_schedule   bool
	restrictions  []PersonRestriction
	badness_score float32 // only update on change
	// days_working_sheduled [7]int // Restriction_working_days
}

func (p Person) validate() (ValReturn, string) {
	for _, r := range p.restrictions {
		if e, s := r.validate(); e != true {
			panic(fmt.Sprintf("!Error for %v: %v", p.id, s))
		}
	}

	// switch hours := p.hours_working; { // move to hours
	// case hours < 0:
	// 	return val_error, "hours < 0"
	// case hours == 0:
	// 	return val_remove, "0 hours"
	// case hours > 40:
	// 	return val_error, "hours over 40"
	// }

	return val_ok, ""
}

func (p Person) calculate_badness() float32 {
	p.badness_score = 0
	for _, r := range p.restrictions {
		p.badness_score += r.badness_score(p)
	}
	for _, day := range p.days {
		p.badness_score += day.badness_score
	}
	return p.badness_score
}

// func (p Person) total_days() int {
// 	total := 0
// 	for d := range p.days_working_sheduled {
// 		total += d
// 	}
// 	return total
// }

// func (p Person) Action_add_day(day int) { // only in 8 hour increments
// 	p.days_working_sheduled[day] = 1
// 	p.hours_working += 8
// 	p.in_schedule = true
// }

// func (p Person) Action_remove_day(day int) {
// 	p.days_working_sheduled[day] = 0
// 	p.hours_working -= 8
// 	if p.total_days() == 0 {
// 		p.in_schedule = false
// 	}
// }
