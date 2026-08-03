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

type Day struct {
	nr            int
	active        bool
	hours_planned int
	restrictions  []d_restriction_int
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
