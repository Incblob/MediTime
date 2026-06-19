package optimizerobjects

import "testing"

func TestStructs(t *testing.T) {
	p := Person{
		id:          "test",
		in_schedule: false,
		restrictions: []PersonRestriction{
			Restriction_working_days{
				days_restrictions: [7]int{0, 1, 1, 1, 0, 1, 0},
				days_total:        4,
				scaling:           1,
			},
		},
		hours_working:         40,
		badness_score:         0,
		days_working_sheduled: [7]int{1, 1, 1, 1, 1, 1, 1},
	}

}
