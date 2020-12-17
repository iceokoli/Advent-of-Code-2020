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

//Ship is a type
type Ship struct {
	env          string
	instructions []string
	position     map[string]float64 //latitude and logitude
	direction    string             //where ship facing at start
}

//Initialize Sets up the problem
func (s *Ship) Initialize() {
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

func (s *Ship) applyShift(a string, v float64) {

	switch a {
	case "N":
		s.position["lat"] += v
	case "S":
		s.position["lat"] -= v
	case "E":
		s.position["long"] += v
	case "W":
		s.position["long"] -= v
	}
}

func (s *Ship) goRight(v float64) {

	var newDirection string
	switch s.direction {
	case "N":
		switch v {
		case 90:
			newDirection = "E"
		case 180:
			newDirection = "S"
		case 270:
			newDirection = "W"
		}
	case "S":
		switch v {
		case 90:
			newDirection = "W"
		case 180:
			newDirection = "N"
		case 270:
			newDirection = "E"
		}
	case "E":
		switch v {
		case 90:
			newDirection = "S"
		case 180:
			newDirection = "W"
		case 270:
			newDirection = "N"
		}
	case "W":
		switch v {
		case 90:
			newDirection = "N"
		case 180:
			newDirection = "E"
		case 270:
			newDirection = "S"
		}
	}
	s.direction = newDirection
}

func (s *Ship) goLeft(v float64) {

	var newDirection string
	switch s.direction {
	case "S":
		switch v {
		case 90:
			newDirection = "E"
		case 180:
			newDirection = "N"
		case 270:
			newDirection = "W"
		}
	case "N":
		switch v {
		case 90:
			newDirection = "W"
		case 180:
			newDirection = "S"
		case 270:
			newDirection = "E"
		}
	case "W":
		switch v {
		case 90:
			newDirection = "S"
		case 180:
			newDirection = "E"
		case 270:
			newDirection = "N"
		}
	case "E":
		switch v {
		case 90:
			newDirection = "N"
		case 180:
			newDirection = "W"
		case 270:
			newDirection = "S"
		}
	}
	s.direction = newDirection
}

func (s *Ship) goForward(v float64) {

	switch s.direction {
	case "N":
		s.position["lat"] += v
	case "S":
		s.position["lat"] -= v
	case "E":
		s.position["long"] += v
	case "W":
		s.position["long"] -= v
	}
}

//BeginJourney Apply Instruction based on rules
func (s *Ship) BeginJourney() {

	for _, instr := range s.instructions {
		action := instr[:1]
		value, _ := strconv.ParseFloat(instr[1:], 64)
		switch action {
		case "F":
			s.goForward(value)
		case "L":
			s.goLeft(value)
		case "R":
			s.goRight(value)
		default:
			s.applyShift(action, value)
		}
	}
}
