package main

import "fmt"

func returnEarliest(bus int, start int) int {

	result := 0
	for result <= start {
		result += bus
	}
	return result
}

//PartA runs fist part of the challenge
func PartA(start int, buses []int) {

	result := make(map[int]int)
	for _, bus := range buses {
		result[bus] = returnEarliest(bus, start)
	}

	var earliestBus int
	temp := 0
	for bus, time := range result {
		if temp == 0 || time < temp {
			temp = time
			earliestBus = bus
		}
	}
	answer := (result[earliestBus] - start) * earliestBus
	fmt.Println("ansswer to part a: ", answer)
}
