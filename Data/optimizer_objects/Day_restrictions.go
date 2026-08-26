package optimizerobjects

/// --------------------- Restrictions ----------------

type D_Restr_Impl interface {
	validate() (bool, string)
	badness_score(d Day) float32
}

// --------- Tot hours
type Day_R_hours struct {
	should_hours float32
}

func (r Day_R_hours) validate() (bool, string) {
	if r.should_hours < 0 {
		return false, "hour limitation under 0"
	}
	return true, ""
}

func (r Day_R_hours) badness_score(day Day) float32 {
	return abs_val(r.should_hours - day.current_hours)
}

// ------ daily hours max
type Day_R_max_hours struct {
	max_hours float32
}

func (r Day_R_max_hours) validate() (bool, string) {
	if r.max_hours < 0 {
		return false, "hour limitation under 0"
	}
	return true, ""
}

func (r Day_R_max_hours) badness_score(day Day) float32 {
	return max(day.current_hours-r.max_hours, 0)
}
