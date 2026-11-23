package optimizerobjects

import (
	"log"
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

func Make_day(num int, hours float32) Day {
	var available bool
	var restriction []D_Restr_Impl

	if hours > 0 {
		available = true
		restriction = []D_Restr_Impl{&Day_R_hours{should_hours: hours}}
	} else {
		available = false
	}

	return Day{
		nr:            num,
		available:     available,
		current_hours: 0,
		restrictions:  restriction,
		badness_score: 0,
	}
}

func MakeLimitedDays(available [7]bool) [7]Day {
	var days = [7]Day{}
	for i, d := range available {
		if d {
			days[i] = Make_day(i, 8)
		} else {
			days[i] = Make_day(i, 0)
		}
	}
	return days
}

func MakeNormal5Days() [7]Day {
	var days = [7]Day{}
	for i := range 5 {
		days[i] = Make_day(i, 8)
	}
	for i := 5; i < 7; i++ {
		days[i] = Make_day(i, 0)
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

func (d *Day) Calculate_badness() {
	log.Printf("\tcalculating bad_score for "+G.Bold+" DAY %v", d.nr)
	var tot float32 = 0
	for _, restriction := range d.restrictions {
		restriction.calculate_badness(*d)
		tot += restriction.GetScore()
	}
	d.badness_score = tot
}
func (d Day) GetScore() float32 {
	return d.badness_score
}

func (d *Day) Make_change() {
	if !d.available {
		log.Panicf("trying to change NOT AVAILABLE Day %v", d.nr)
	}
	// selected_change := rand.Int32N(int32(len(available_changes)))
	change_nr := rand.IntN(len(day_available_changes))
	day_available_changes[change_nr](d)
	log.Printf("making"+G.Underline+"change %v"+G.Reset+" on day %v", change_nr, d.nr)
	d.Calculate_badness()
}

func (d *Day) change_hours() {
	const amount float32 = 2
	//figure out direction of change
	direction := rand.Float32()
	var change bool = true
	if direction > 0.5 {
		change = false
	}
	if change {
		d.current_hours -= amount
	} else {
		d.current_hours += amount
	}
}

var day_available_changes = []func(d *Day){
	(*Day).change_hours,
}
