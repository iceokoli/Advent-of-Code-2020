package main

import (
	"fmt"
	"strconv"
)

type timeTable struct {
	pos  int
	time []int
}

func returnSchedule(bus int) []int {
	var result []int
	temp := (int(1e6/bus) + 1) * bus
	fmt.Println(temp)
	for temp <= 1e7 {
		temp += bus
		if temp > 1e6 {
			result = append(result, temp)
		}

	}
	return result
}

func check(schedules map[string]timeTable, buses []string) (int, int) {

	loop := 0
	start := buses[0]
	var result int
	var verify int
out:
	for loop < len(schedules[start].time) {
		count := 0
		for _, bus := range buses {
			if loop == len(schedules[bus].time) {
				break out
			}
			a := schedules[start].time[loop]
			b := schedules[bus].time[loop]
			c := schedules[bus].pos

			if a+c == b {
				result = b
				count++
				verify = count
			}

			if count == len(buses) {
				break out
			}
		}
		loop++
	}

	return result, verify

}

//PartB runs second part of the challenge
func PartB(buses []string) {

	schedules := make(map[string]timeTable)
	var busClean []string
	for idx, bus := range buses {
		if bus != "x" {
			busClean = append(busClean, bus)
			temp, _ := strconv.Atoi(bus)
			var temp2 timeTable
			temp2.pos = idx
			temp2.time = returnSchedule(temp)
			schedules[bus] = temp2
		}
	}

	result, v := check(schedules, busClean)
	if v == len(busClean) {
		fmt.Println("answer to part b ", result, v)
	} else {
		fmt.Println("Try Again")
	}

}
