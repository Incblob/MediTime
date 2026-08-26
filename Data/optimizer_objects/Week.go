package optimizerobjects

type Week struct {
	persons           []Person
	week_restrictions []W_Restr_Impl
}

// ------------------- Restrictions

type W_Restr_Impl interface {
	validate() bool
	badness_score(w Week) float32
}

// ----- workers per day
type W_Restr_workers_per_day struct {
	W_Restr_Impl
	days [7]uint
}

func (w_r W_Restr_workers_per_day) badness_score(week Week) float32 {
	var days_workers [7]uint
	for _, p := range week.persons {
		for i, d := range p.days {
			if d.current_hours > 0 {
				days_workers[i] += 1
			}
		}
	}
	var tot uint = 0

	for i, _ := range days_workers {
		tot += abs_val(days_workers[i] - w_r.days[i])
	}
	return float32(tot)
}

// -----
