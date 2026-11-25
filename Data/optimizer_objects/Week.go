package optimizerobjects

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	G "optimizer/cmd/globals"
)

type Week struct {
	Persons           []Person
	Week_restrictions []W_Restr_Impl
	badness_score     float32
}

func (w Week) validate() (bool, string) {
	if len(w.Persons) < 1 {
		return false, "no persons assigned!"
	}
	return true, ""
}

func (w *Week) Calculate_badness() {
	slog.Info(G.CPurple + "calculating bad_score for" + G.Bold + " WEEK")
	var score float32 = 0
	for _, r := range w.Week_restrictions {
		r.Calculate_badness_score(*w)
		score += r.GetScore()
	}
	w.badness_score = score
	slog.Info(fmt.Sprintf("Base Score for "+G.CBlue+"Week:"+G.Reset+G.CCyan+"%v", w.badness_score))
}

func (w Week) GetScore() float32 {
	return w.badness_score
}

// ------------------- Restrictions

type W_Restr_Impl interface {
	validate() bool
	Calculate_badness_score(w Week)
	GetScore() float32
}

// ----- workers per day
type W_Restr_workers_per_day struct {
	Days          [7]int
	badness_score float32
}

// func (w_r W_Restr_workers_per_day)get_score() float32{
// 	return w_r.score
// }

func (w_r W_Restr_workers_per_day) GetScore() float32 {
	return w_r.badness_score
}

func (w_r *W_Restr_workers_per_day) Calculate_badness_score(week Week) {
	slog.Debug("\tcalculating bad_score for " + G.Bold + "restriction" + G.Reset + G.Underline + " workers_per_day")
	var days_workers [7]int
	for _, p := range week.Persons {
		for i, d := range p.days {
			if d.current_hours > 0 {
				days_workers[i] += 1
			}
		}
	}
	var tot float32 = 0

	for i := range days_workers {
		tot += float32(abs_val(days_workers[i] - w_r.Days[i]))
	}
	fmt.Printf("tot workers per day = %v", tot)
	// normalize against number of people.
	// The more people the less relevant is the deviation per day? maybe?
	// TODO test if this is fine
	tot /= float32(len(week.Persons))

	w_r.badness_score = float32(tot)
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
	slog.Debug(fmt.Sprintf("making"+G.Underline+"change"+G.Reset+" on Person %v", w.Persons[change_person].ID))
	w.Persons[change_person].Make_change()
}
