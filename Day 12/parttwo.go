package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

//ShipB is a type
type ShipB struct {
	env          string
	instructions []string
	position     map[string]float64 //latitude and logitude
	waypoint     map[string]float64 //where ship facing at start
}

//Initialize Sets up the problem
func (s *ShipB) Initialize() {
	// Get current directory
	_, filename, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}
	// read in the file
	raw, err := ioutil.ReadFile(dir + "/" + s.env + ".txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// turn the file into an array
	s.instructions = strings.Split(string(raw), "\n")
	s.position = make(map[string]float64)
	s.position["lat"] = 0
	s.position["long"] = 0
}

//BeginJourney Apply Instruction based on rules
func (s *ShipB) BeginJourney() {

	for _, instr := range s.instructions {
		action := instr[:1]
		value, _ := strconv.ParseFloat(instr[1:], 64)
		switch action {
		case "F":
			s.goForward(value)
		case "L":
			s.turn(action, value)
		case "R":
			s.turn(action, value)
		default:
			s.applyShift(action, value)
		}
		//fmt.Println(action, value)
		//fmt.Println(s.waypoint, s.position)
	}
}

func (s *ShipB) applyShift(a string, v float64) {

	switch a {
	case "N":
		s.waypoint["lat"] += v
	case "S":
		s.waypoint["lat"] -= v
	case "E":
		s.waypoint["long"] += v
	case "W":
		s.waypoint["long"] -= v
	}

}

func (s *ShipB) goForward(v float64) {

	s.position["lat"] += (v * s.waypoint["lat"])
	s.position["long"] += (v * s.waypoint["long"])
}

func (s *ShipB) turn(a string, v float64) {

	var newLong, newLat float64
	if a == "L" {
		i := v
		for i > 0 {
			newLong = -s.waypoint["lat"]
			newLat = s.waypoint["long"]
			i -= 90
			s.waypoint["lat"] = newLat
			s.waypoint["long"] = newLong
		}
	} else if a == "R" {
		i := v
		for i > 0 {
			newLong = s.waypoint["lat"]
			newLat = -s.waypoint["long"]
			i -= 90
			s.waypoint["lat"] = newLat
			s.waypoint["long"] = newLong
		}
	}

}
