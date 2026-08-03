package optimizerobjects

const HolidayScaling = 10

type PersonRestriction interface {
	validate() (bool, string)
	badness_score(p Person) float32 // minimize this function
}

// TODO check validity of restrictions
// --------------
type Person_R_working_hours struct {
	PersonRestriction
	total_time float32
	scaling    float32
}

func (r Person_R_working_hours) validate() (bool, string) {
	if r.total_time < 0 {
		return false, "total hours < 0"
	}
	return true, ""
}

func (r Person_R_working_hours) badness_score(p Person) float32 {
	var tot float32 = 0
	for _, day := range p.days {
		tot += float32(day.hours_planned)
	}
	return tot
}

// --------------
type Person_R_total_working_days struct {
	PersonRestriction //individual days restrictions moved to days
	days_total        int
	scaling           float32
}

func (r Person_R_total_working_days) validate() (bool, string) {
	if r.days_total < 0 {
		return false, "days total < 0"
	}
	return true, ""
}

func (r Person_R_total_working_days) badness_score(p Person) float32 {

	var total_score float32 = 0

	for _, day := range p.days {
		total_score += float32(bool_to_int(day.active))
	}
	return total_score
}

// --------------
