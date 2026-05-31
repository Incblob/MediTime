package optimizerobjects

import (
	"golang.org/x/exp/constraints"
)

const HolidayScaling = 10

type Person struct {
	id            string
	in_schedule   bool
	restrictions  []Restriction
	hours_working float32
	badness_score float32
	days_working  [7]int
}

func (p Person) validate() {
	for i := range p.days_working {
		if i < 0 || i > 1 {
			panic("validation error in Person days working")
		}
	}
}

func (p Person) calculate_badness() float32 {
	var total float32 = 0
	for _, r := range p.restrictions {
		total += r.badness_score(p)
	}
	p.badness_score = total
	return p.badness_score
}

type Restriction interface {
	badness_score(p Person) float32 // minimize this function
}

//TODO check validity of restrictions

type Restriction_working_hours struct {
	total_time float32
	scaling    float32
}

func (r Restriction_working_hours) badness_score(p Person) float32 {
	return abs_val[float32](r.total_time-p.hours_working) * r.scaling //scaling under/over differently?
}

type Restriction_working_days struct {
	days [7]int // 1: can work, 0: should not, -1: holiday
	// is_arbitrary bool
	days_total int
	scaling    float32
}

// if sum_day_bools(p.days_working) == 0{
// TODO if no work then remove from shedule
// }

// badness score is calculated by comparing the restriction days to the days
// assigned to the person. restriction has -1 for holidays, 0 for should not work and 1 for available
// for -1 the scaling is higher
// if the days assigned exceed the wanted days then the score is raised further
// TODO this is probably going to get much more complicated with individual op sheduling
func (r Restriction_working_days) badness_score(p Person) float32 {

	if len(r.days) != len(p.days_working) {
		panic("days length not matching")
	}

	days_scaled := r.days
	for i := range len(days_scaled) { // scale holidays (r.days == -1) to -10
		if days_scaled[i] < 0 {
			days_scaled[i] = -10
		}
	}

	var total_score float32 = 0
	var assigned_days int = 0

	for i, v := range days_scaled {
		switch c := v + p.days_working[i]; {
		case c < -1: // -10 and (0 or 1)
			total_score += r.scaling * HolidayScaling
		case c == 1: // (0 and 1) or (1 and 0)
			total_score += r.scaling
		} // (1 and 1) or (0 and 0) is ok
		assigned_days += p.days_working[i] // don't exceed the number of days
	}
	return total_score + float32(assigned_days)*r.scaling
}

func abs_val[L constraints.Integer | constraints.Float](l L) L {
	return L(l)
}
