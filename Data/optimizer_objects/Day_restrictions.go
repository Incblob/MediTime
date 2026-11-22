package optimizerobjects

import (
	"log"
	G "optimizer/cmd/globals"
)

/// --------------------- Restrictions ----------------

type D_Restr_Impl interface {
	validate() (bool, string)
	calculate_badness(d Day)
	GetScore() float32
}

// --------- Tot hours
type Day_R_hours struct {
	should_hours  float32
	badness_score float32
}

func (r Day_R_hours) validate() (bool, string) {
	if r.should_hours < 0 {
		return false, "hour limitation under 0"
	}
	return true, ""
}

func (r *Day_R_hours) calculate_badness(day Day) {
	log.Println("\t\t " + G.Bold + "DAY " + G.Reset + "- calculating bad_score for restriction" + G.Underline + " hours per day")
	r.badness_score = abs_val(r.should_hours - day.current_hours)
	log.Printf(G.Bold+G.CCyan+"Score  == %v", r.badness_score)
}

func (r Day_R_hours) GetScore() float32 {
	return r.badness_score
}

// ------ daily hours max
type Day_R_max_hours struct {
	max_hours     float32
	badness_score float32
}

func (r Day_R_max_hours) validate() (bool, string) {
	if r.max_hours < 0 {
		return false, "hour limitation under 0"
	}
	return true, ""
}

func (r *Day_R_max_hours) calculate_badness(day Day) {
	log.Printf("\t\t - calculating bad_score for restriction"+G.Underline+" max hours per day for day %v \n", day.nr)
	r.badness_score = max(day.current_hours-r.max_hours, 0)
	log.Printf(G.CCyan+G.Bold+"Score  == %v", r.badness_score)
}

func (r Day_R_max_hours) GetScore() float32 {
	return r.badness_score
}
