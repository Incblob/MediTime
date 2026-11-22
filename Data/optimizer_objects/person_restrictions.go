package optimizerobjects

import (
	"log"
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
	log.Println("\t\t  calculating bad_score for " + G.Bold + "restriction " + G.Reset + G.Underline + "total working hours")
	var tot float32 = 0
	for _, day := range p.days {
		tot += float32(day.current_hours)
	}
	r.badness_score = tot
	log.Printf(G.Underline+G.Bold+"Score  == %v", r.badness_score)
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

	log.Println("\t\t  calculating bad_score for " + G.Bold + "restriction " + G.Reset + G.Underline + "total working days")
	var total_score float32 = 0

	for _, day := range p.days {
		if day.current_hours > 0 {
			total_score += 1
		}
		// total_score += float32(bool_to_int(day.active))
	}
	r.badness_score = total_score
	log.Printf(G.CCyan+G.Bold+"Score  == %v", r.badness_score)
}

func (r Person_R_total_working_days) GetScore() float32 {
	return r.badness_score
}

// --------------
