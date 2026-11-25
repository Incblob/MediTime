package optimizerobjects

import (
	"fmt"
	"log/slog"
	G "optimizer/cmd/globals"
)

const HolidayScaling = 10

type P_Restr_impl interface {
	validate() (bool, string)
	calculate_badness(p Person) // minimize this function
	GetScore() float32
}

// TODO check validity of restrictions
// --------------

var scaling_map_person = map[string](float32){
	"Person_R_total_working_hours": 5,
	"Person_R_total_working_days":  10,
}

type Person_R_total_working_hours struct {
	// P_Restr_impl
	Total_time    float32
	Scaling       float32
	badness_score float32
}

func (r Person_R_total_working_hours) validate() (bool, string) {
	if r.Total_time < 0 {
		return false, "total hours < 0"
	}
	return true, ""
}

func (r *Person_R_total_working_hours) calculate_badness(p Person) {
	slog.Debug("\t\t  calculating bad_score for " + G.Bold + "restriction " + G.Reset + G.Underline + "total working hours")
	var tot float32 = 0
	for _, day := range p.days {
		tot += float32(day.current_hours)
	}
	// normalized to max wanted
	r.badness_score = abs_val(tot - r.Total_time)
	r.badness_score /= r.Total_time
	slog.Debug(fmt.Sprintf(G.Underline+G.Bold+"Score  == %v", r.badness_score))
}

func (r Person_R_total_working_hours) GetScore() float32 {
	return r.badness_score
}

// --------------

type Person_R_total_working_days struct {
	// P_Restr_impl //individual days restrictions moved to days
	Total_days    int
	Scaling       float32
	badness_score float32
}

func (r Person_R_total_working_days) validate() (bool, string) {
	if r.Total_days < 0 {
		return false, "days total < 0"
	}
	if r.Total_days > 7 {
		return false, "days total > 7"
	}

	// validate that for the person the active days are less than the total days . ?
	// BREAKS interface
	// var tot_active_days = 0
	// for _, day := range p.days {
	// 	if day.current_hours > 0 {
	// 		tot_active_days += 1
	// 	}
	// }
	// if tot_active_days < r.Total_days {
	// 	return false, "the active days are less than the working days"
	// }

	return true, ""
}

func (r *Person_R_total_working_days) calculate_badness(p Person) {
	slog.Info("\t\t  calculating bad_score for " + G.Bold + "restriction " + G.Reset + G.Underline + "total working days")
	var total_score float32 = 0

	for _, day := range p.days {
		if day.current_hours > 0 {
			total_score += 1
		}
	}
	// r.badness_score = total_score
	// normalize to total days
	r.badness_score = abs_val(total_score - float32(r.Total_days))
	r.badness_score /= float32(r.Total_days)

	slog.Info(fmt.Sprintf(G.CCyan+G.Bold+"Score  == %v", r.badness_score))
}

func (r Person_R_total_working_days) GetScore() float32 {
	return r.badness_score
}

// --------------
