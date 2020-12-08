package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

var totalRows int = 128
var totalCols int = 8

type boardingPass struct {
	row int
	col int
	id  int
}

type byID []boardingPass

func (i byID) Len() int           { return len(i) }
func (i byID) Swap(x, y int)      { i[x], i[y] = i[y], i[x] }
func (i byID) Less(x, y int) bool { return i[x].id < i[y].id }

func getInput() []string {
	// Get current directory
	_, filename, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}

	// read in the file
	raw, err := ioutil.ReadFile(dir + "/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// turn the file into an array
	result := strings.Split(string(raw), "\n") //list of strings

	return result
}

func structurePasses(data []string) []boardingPass {

	var tempPass []boardingPass
	for _, i := range data {

		// figure out the row
		tempRow := make(map[string]int)
		tempRow["min"] = 0
		tempRow["max"] = totalRows - 1
		for _, x := range i[:7] {
			if string(x) == "F" {
				tempRow["max"] = int((tempRow["max"]-tempRow["min"])/2) + tempRow["min"]
			} else {
				tempRow["min"] = int((tempRow["max"]-tempRow["min"])/2) + 1 + tempRow["min"]
			}

		}
		// figure out the columns
		tempCol := make(map[string]int)
		tempCol["min"] = 0
		tempCol["max"] = totalCols - 1
		for _, y := range i[7:] {
			if string(y) == "L" {
				tempCol["max"] = int((tempCol["max"]-tempCol["min"])/2) +
					tempCol["min"]
			} else {
				tempCol["min"] = int((tempCol["max"]-tempCol["min"])/2) + 1 +
					tempCol["min"]
			}

		}
		tempPass = append(
			tempPass,
			boardingPass{
				tempRow["min"],
				tempCol["min"],
				(tempRow["min"] * 8) + tempCol["min"],
			},
		)

	}

	return tempPass
}

func findMaxID(data []boardingPass) int {

	var m int
	for i, pass := range data {
		if i == 0 || pass.id > m {
			m = pass.id
		}
	}
	return m
}

func findMySeat(data []boardingPass) int {

	sort.Sort(byID(data))
	fmt.Println(data[:5])

	var save int
	for i, seat := range data {
		if (i + 1) >= len(data) {
			continue
		} else if seat.id+1 != data[i+1].id {
			save = seat.id
		}
	}

	return save + 1
}

func main() {
	rawData := getInput()
	partA := structurePasses(rawData)
	fmt.Println(findMaxID(partA))
	fmt.Println(findMySeat(partA))
}
