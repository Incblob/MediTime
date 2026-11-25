package optimizerobjects

import (
	"fmt"
	"log/slog"
	G "optimizer/cmd/globals"
)

/// --------------------- Restrictions ----------------

type D_Restr_Impl interface {
	validate() (bool, string)
	calculate_badness(d Day)
	GetScore() float32
}

var scaling_map_day = map[string](float32){
	"day_R_should_hours":  5,
	"day_R_max_hours":     5,
	"day_R_not_available": 20,
}

// --------- Tot hours
type day_R_should_hours struct {
	should_hours  float32
	badness_score float32
	scaling       float32
}

func Make_day_R_hours(should_hours float32) *day_R_should_hours {
	return &day_R_should_hours{
		should_hours:  should_hours,
		badness_score: 0,
		scaling:       scaling_map_day["day_R_should_hours"],
	}
}

func (r day_R_should_hours) validate() (bool, string) {
	if r.should_hours < 0 {
		return false, "hour limitation under 0"
	}
	return true, ""
}

func (r *day_R_should_hours) calculate_badness(day Day) {
	slog.Debug("\t\t " + G.Bold + "DAY " + G.Reset + "- calculating bad_score for restriction" + G.Underline + " hours per day")
	r.badness_score = abs_val(r.should_hours-day.current_hours) * r.scaling
	r.badness_score /= r.should_hours // normalize
	slog.Debug(fmt.Sprintf(G.Bold+G.CCyan+"Score  == %v", r.badness_score))
}

func (r day_R_should_hours) GetScore() float32 {
	return r.badness_score
}

// ------ daily hours max
type day_R_max_hours struct {
	max_hours     float32
	badness_score float32
	scaling       float32
}

func Make_day_R_max_hours(should_hours float32) *day_R_max_hours {
	return &day_R_max_hours{
		max_hours:     should_hours,
		badness_score: 0,
		scaling:       scaling_map_day["day_R_max_hours"],
	}
}

func (r day_R_max_hours) validate() (bool, string) {
	if r.max_hours < 0 {
		return false, "hour limitation under 0"
	}
	return true, ""
}

func (r *day_R_max_hours) calculate_badness(day Day) {
	slog.Debug(fmt.Sprintf("\t\t - calculating bad_score for restriction"+G.Underline+" max hours per day for day %v \n", day.nr))
	r.badness_score = max(day.current_hours-r.max_hours, 0) * r.scaling
	r.badness_score /= r.max_hours
	slog.Debug(fmt.Sprintf(G.CCyan+G.Bold+"Score  == %v", r.badness_score))
}

func (r day_R_max_hours) GetScore() float32 {
	return r.badness_score
}

type day_R_not_available struct {
	scaling       float32
	badness_score float32
}

func Make_day_R_not_available() *day_R_not_available {
	return &day_R_not_available{
		badness_score: 0,
		scaling:       scaling_map_day["day_R_not_available"],
	}
}

func (r day_R_not_available) validate() (bool, string) {
	return true, ""
}

func (r *day_R_not_available) calculate_badness(day Day) {
	slog.Debug(fmt.Sprintf("\t\t - calculating bad_score for restriction"+G.Underline+" available for day %v \n", day.nr))
	r.badness_score = bool_to_float((day.current_hours > 0)) * r.scaling
	// does not need normalization as as gives bool value, scaling does everything
	slog.Debug(fmt.Sprintf(G.CCyan+G.Bold+"Score  == %v", r.badness_score))
}

func (r day_R_not_available) GetScore() float32 {
	return r.badness_score
}
