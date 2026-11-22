package main

import (
	"log"
	opto "optimizer/Data/optimizer_objects"
	G "optimizer/cmd/globals"
)

func main() {

	// log.SetFlags(log.Lshortfile)
	log.SetPrefix(G.Reset)

	log.Println("making person")
	var personA = opto.MakePerson(
		"A",
		opto.MakeNormal5Days(),
		[]opto.P_Restr_impl{
			&opto.Person_R_total_working_hours{Total_time: 40, Scaling: 1},
			&opto.Person_R_total_working_days{Total_days: 5, Scaling: 1},
		})

	log.Println("making personB")
	var personB = opto.MakePerson(
		"B",
		opto.MakeLimitedDays([7]bool{true, true, true, false, false, false, false}),
		[]opto.P_Restr_impl{
			&opto.Person_R_total_working_hours{Total_time: 40, Scaling: 1},
		})

	log.Println("making Week")
	var week = opto.Week{
		Persons: []opto.Person{personA, personB},
		Week_restrictions: []opto.W_Restr_Impl{
			&opto.W_Restr_workers_per_day{Days: [7]uint{2, 2, 2, 1, 1, 0, 0}},
		}}

	optimize(week)

}

func optimize(week opto.Week) {
	var base_score float32 = 0
	week.Calculate_badness()
	base_score += week.Badness_score
	for _, person := range week.Persons {
		person.Calculate_badness()
		base_score += person.GetScore()
	}
	log.Printf(G.CGreen+G.Bold+"Base Score calculated: %v\n", base_score)
	// for i := 0; i < 100; i++ {

	// }
}
