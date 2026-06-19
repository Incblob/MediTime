package optimizerobjects

import (
	"golang.org/x/exp/constraints"
)

const HolidayScaling = 10

type PersonRestriction interface {
	badness_score(p Person) float32 // minimize this function
	validate() (bool, string)
}

//TODO check validity of restrictions

type Restriction_working_hours struct {
	PersonRestriction
	total_time float32
	scaling    float32
}

func (r Restriction_working_hours) badness_score(p Person) float32 {
	return abs_val[float32](r.total_time-p.hours_working) * r.scaling //scaling under/over differently?
}

func (r Restriction_working_hours) validate() (bool, string) {
	if r.total_time < 0 {
		return false, "total hours < 0"
	}
	return true, ""
}

type Restriction_working_days struct {
	PersonRestriction
	days_restrictions [7]int // 1: can work, 0: should not, -1: holiday
	// is_arbitrary bool
	days_total int
	scaling    float32
}

func (r Restriction_working_days) validate() (bool, string) {
	if r.days_total < 0 {
		return false, "days total < 0"
	}
	return true, ""
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

	if len(r.days_restrictions) != len(p.days_working_sheduled) {
		panic("days length not matching")
	}

	days_scaled := r.days_restrictions
	for i := range len(days_scaled) { // scale holidays (r.days == -1) to -10
		if days_scaled[i] < 0 {
			days_scaled[i] = -10
		}
	}

	var total_score float32 = 0
	var assigned_days int = 0

	for i, v := range days_scaled {
		switch c := v + p.days_working_sheduled[i]; {
		case c < -1: // -10 and (0 or 1)
			total_score += r.scaling * HolidayScaling
		case c == 1: // (0 and 1) or (1 and 0)
			total_score += r.scaling
		} // (1 and 1) or (0 and 0) is ok
		assigned_days += p.days_working_sheduled[i] // don't exceed the number of days
	}
	return total_score + float32(assigned_days)*r.scaling
}

func abs_val[L constraints.Integer | constraints.Float](l L) L {
	return L(l)
}
