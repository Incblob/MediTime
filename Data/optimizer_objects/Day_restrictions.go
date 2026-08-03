package optimizerobjects

/// --------------------- Restrictions ----------------

type d_restriction_int interface {
	validate() (bool, string)
	badness_score(d Day) float32
}

// --------- Tot hours
type Day_R_hours struct {
	d_restriction_int
	should_hours float32
}

func (d Day_R_hours) validate() (bool, string) {
	if d.should_hours < 0 {
		return false, "hour limitation under 0"
	}
	return true, ""
}

func (d Day_R_hours) badness_score(current_hours float32) float32 {
	return abs_val(d.should_hours - current_hours)
}

// ---------
