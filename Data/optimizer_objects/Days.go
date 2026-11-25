package optimizerobjects

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	G "optimizer/cmd/globals"
)

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func Make_day(num int, should_hours float32, max_hours float32) Day {
	var available bool
	var restriction []D_Restr_Impl

	restriction = append(restriction, Make_day_R_max_hours(max_hours))

	if should_hours > 0 {
		available = true
		restriction = append(restriction, Make_day_R_hours(should_hours))
	} else {
		available = false
		restriction = append(restriction, Make_day_R_not_available())

	}

	var day = Day{
		nr:            num,
		available:     available,
		current_hours: 0,
		restrictions:  restriction,
		badness_score: 0,
	}
	if r, s := day.validate(); r != true {
		panic(s)
	}
	return day
}

func MakeLimitedDays(available [7]bool) [7]Day {
	var days = [7]Day{}
	for i, d := range available {
		if d {
			days[i] = Make_day(i, 8, 8)
		} else {
			days[i] = Make_day(i, 0, 8)
		}
	}
	return days
}

func MakeNormal5Days() [7]Day {
	var days = [7]Day{}
	for i := range 5 {
		days[i] = Make_day(i, 8, 8)
	}
	for i := 5; i < 7; i++ {
		days[i] = Make_day(i, 0, 8)
	}
	return days
}

type Day struct {
	nr            int
	available     bool    // Is the day available for sheduling (not set to 0), does *not* mean that the day has any hours sheduled.
	current_hours float32 // used to get whether the day is currently set in the shedule, don't use 'active'
	restrictions  []D_Restr_Impl
	badness_score float32
}

func (d Day) validate() (bool, string) {
	slog.Debug(fmt.Sprintf("validating Day %v", d.nr))
	for _, r := range d.restrictions {
		if e, s := r.validate(); e != true {
			panic(fmt.Sprintf("!Error for %v: %v", d.nr, s))
		}
	}

	return true, ""
}

func (d *Day) Calculate_badness() {
	slog.Debug(fmt.Sprintf("\tcalculating bad_score for "+G.Bold+" DAY %v", d.nr))
	var tot float32 = 0
	for _, restriction := range d.restrictions {
		restriction.calculate_badness(*d)
		tot += restriction.GetScore()
	}
	d.badness_score = tot
	slog.Info(fmt.Sprintf("\tScore for "+G.CYellow+"Day %v:"+G.Reset+G.CCyan+"%v", d.nr, d.badness_score))
}
func (d Day) GetScore() float32 {
	return d.badness_score
}

func (d *Day) Make_change() {
	if !d.available {
		panic(fmt.Sprintf("trying to change NOT AVAILABLE Day %v", d.nr))
	}
	// selected_change := rand.Int32N(int32(len(available_changes)))
	change_nr := rand.IntN(len(day_available_changes))
	day_available_changes[change_nr](d)
	slog.Debug(fmt.Sprintf("making"+G.Underline+"change %v"+G.Reset+" on day %v", change_nr, d.nr))
	d.Calculate_badness()
}

func (d *Day) change_hours() {
	const amount float32 = 2
	//figure out direction of change_up
	var change_up bool = false
	if rand.Float32() > 0.5 {
		change_up = true
	}

	if change_up {
		d.current_hours -= amount
	} else {
		d.current_hours += amount
	}
}

var day_available_changes = []func(d *Day){
	(*Day).change_hours,
}
