package optimizerobjects

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func Make_day(num int, hours float32)  Day{
	var active bool
	var restriction []D_Restr_Impl{}

	if hours>0 {
		active = true
		restriction = []D_Restr_Impl{Day_R_hours{hours}}
	} else {
		active = false
	}

	return Day{
		nr: num,
		active: active,
		current_hours: 0,
		restrictions: restriction,
		badness_score: 0,
	}
}

func Make7NormalDays() [7]Day{
	var days = [7]Day{}
	for i := range 5 {
		days = append(days, Make_day(i, 8))
	}
	for i := 5; i < 7; i++ {
		days = append(days, Make_day(i, 0))
	}
	return days
}

type Day struct {
	nr            int
	active        bool    // Is the day available for sheduling (not set to 0), does *not* mean that the day has any hours sheduled.
	current_hours float32 // used to get whether the day is currently set in the shedule, don't use 'active'
	restrictions  []D_Restr_Impl
	badness_score float32
}

func (d Day) calculate_badness() {
	var tot float32 = 0
	for _, restriction := range d.restrictions {
		tot += restriction.badness_score(d)
	}
	d.badness_score = tot
}

func (d Day) make_change() {
	//TODO
	panic("Not implemented")
	d.calculate_badness()
}
