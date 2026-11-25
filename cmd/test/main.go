package main

import (
	"fmt"
	"log"
	"log/slog"
	opto "optimizer/Data/optimizer_objects"
	G "optimizer/cmd/globals"
)

func main() {

	// slog.SetFlags(slog.Lshortfile)
	log.SetPrefix(G.Reset)

	slog.Info("making person")
	var personA = opto.MakePerson(
		"A",
		opto.MakeNormal5Days(),
		[]opto.P_Restr_impl{
			&opto.Person_R_total_working_hours{Total_time: 40, Scaling: 1},
			&opto.Person_R_total_working_days{Total_days: 5, Scaling: 1},
		})

	slog.Info("making personB")
	var personB = opto.MakePerson(
		"B",
		opto.MakeLimitedDays([7]bool{true, true, true, false, false, false, false}),
		[]opto.P_Restr_impl{
			&opto.Person_R_total_working_hours{Total_time: 40, Scaling: 1},
		})

	slog.Info("making Week")
	var week = opto.Week{
		Persons: []opto.Person{personA, personB},
		Week_restrictions: []opto.W_Restr_Impl{
			&opto.W_Restr_workers_per_day{Days: [7]int{2, 2, 2, 1, 1, 0, 0}},
		}}

	optimize(week)

}

func optimize(week opto.Week) {
	var base_score float32 = 0
	week.Calculate_badness()
	base_score += week.GetScore()
	for _, person := range week.Persons {
		person.Calculate_badness()
		base_score += person.GetScore()
	}
	slog.Info(fmt.Sprintf(G.CGreen+G.Bold+"Base Score calculated: %v\n", base_score))
	// for i := 0; i < 100; i++ {

	// }
}
