package main

import optimizerobjects "optimizer/Data/optimizer_objects"

func main() {

	var daysA = optimizerobjects.Make7NormalDays()
	var personA = optimizerobjects.MakePerson(
		"A",
		daysA,
		[]optimizerobjects.P_Restr_impl{
			optimizerobjects.Person_R_total_working_hours{40, 1},
		})

}
