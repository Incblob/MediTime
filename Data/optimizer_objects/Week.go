package optimizerobjects

import (
	"log"
	"math/rand/v2"
	G "optimizer/cmd/globals"
)

type Week struct {
	Persons           []Person
	Week_restrictions []W_Restr_Impl
	Badness_score     float32
}

func (w Week) Calculate_badness() {
	log.Println("calculating bad_score for" + G.Bold + " WEEK")
	var score float32 = 0
	for _, r := range w.Week_restrictions {
		r.Calculate_badness_score(w)
		score += r.GetScore()
	}
}

// ------------------- Restrictions

type W_Restr_Impl interface {
	validate() bool
	Calculate_badness_score(w Week)
	GetScore() float32
}

// ----- workers per day
type W_Restr_workers_per_day struct {
	Days  [7]uint
	score float32
}

// func (w_r W_Restr_workers_per_day)get_score() float32{
// 	return w_r.score
// }

func (w_r W_Restr_workers_per_day) GetScore() float32 {
	return w_r.score
}

func (w_r *W_Restr_workers_per_day) Calculate_badness_score(week Week) {
	log.Println("\tcalculating bad_score for" + G.Bold + "restriction" + G.Reset + G.Underline + " workers_per_day")
	var days_workers [7]uint
	for _, p := range week.Persons {
		for i, d := range p.days {
			if d.current_hours > 0 {
				days_workers[i] += 1
			}
		}
	}
	var tot uint = 0

	for i := range days_workers {
		tot += abs_val(days_workers[i] - w_r.Days[i])
	}
	w_r.score = float32(tot)
	// return float32(tot)
}

func (w_r W_Restr_workers_per_day) validate() bool {
	for _, d := range w_r.Days {
		if d < 0 {
			return false
		}
	}
	return true
}

// optimization
func (w *Week) Make_change() {

	change_person := rand.IntN(len(w.Persons))
	log.Printf("making"+G.Underline+"change"+G.Reset+" on Person %v", w.Persons[change_person].id)
	w.Persons[change_person].Make_change()
}
