package main

import (
	"fmt"
	"math"
)

func runPartA() {
	titanic := Ship{env: "prod", direction: "E"}
	titanic.Initialize()
	titanic.BeginJourney()
	fmt.Println("A) Current position of the ship: ", titanic.position)
	fmt.Println(
		"Answer to part a: ",
		math.Abs(titanic.position["lat"])+
			math.Abs(titanic.position["long"]),
	)
}

func runPartB() {

	titanic := ShipB{
		env:      "prod",
		waypoint: map[string]float64{"lat": 1, "long": 10},
	}
	//fmt.Println(titanic.waypoint)
	titanic.Initialize()
	titanic.BeginJourney()
	fmt.Println("B) Current position of the ship: ", titanic.position)
	fmt.Println(
		"Answer to part b: ",
		math.Abs(titanic.position["lat"])+
			math.Abs(titanic.position["long"]),
	)

}

func main() {
	runPartA()
	runPartB()
}
