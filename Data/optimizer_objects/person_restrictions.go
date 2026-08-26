package optimizerobjects

const HolidayScaling = 10

type P_Restr_impl interface {
	validate() (bool, string)
	badness_score(p Person) float32 // minimize this function
}

// TODO check validity of restrictions
// --------------
type Person_R_total_working_hours struct {
	// P_Restr_impl
	total_time float32
	scaling    float32
}

func (r Person_R_total_working_hours) validate() (bool, string) {
	if r.total_time < 0 {
		return false, "total hours < 0"
	}
	return true, ""
}

func (r Person_R_total_working_hours) badness_score(p Person) float32 {
	var tot float32 = 0
	for _, day := range p.days {
		tot += float32(day.current_hours)
	}
	return tot
}

// --------------

type Person_R_total_working_days struct {
	// P_Restr_impl //individual days restrictions moved to days
	days_total int
	scaling    float32
}

func (r Person_R_total_working_days) validate(p Person) (bool, string) {
	if r.days_total < 0 {
		return false, "days total < 0"
	}

	var tot_active_days = 0
	for _, day := range p.days {
		if day.current_hours > 0 {
			tot_active_days += 1
		}
	}
	if tot_active_days < r.days_total {
		return false, "the active days are less than the working days"
	}

	return true, ""
}

func (r Person_R_total_working_days) badness_score(p Person) float32 {

	var total_score float32 = 0

	for _, day := range p.days {
		if day.current_hours > 0 {
			total_score += 1
		}
		// total_score += float32(bool_to_int(day.active))
	}
	return total_score
}

// --------------
