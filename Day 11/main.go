package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

func deepCopy(b [][]string) [][]string {
	a := make([][]string, len(b))
	for i := range b {
		a[i] = make([]string, len(b[i]))
		copy(a[i], b[i])
	}
	return a

}

type challenge struct {
	env   string
	input [][]string
}

func (c challenge) readInput() [][]string {
	// Get current directory
	_, filename, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}
	// read in the file
	raw, err := ioutil.ReadFile(dir + "/" + c.env + ".txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// turn the file into an array
	var result [][]string
	for _, i := range strings.Split(string(raw), "\n") {
		var temp []string
		for _, j := range []rune(i) {
			temp = append(temp, string(j))
		}
		result = append(result, temp)
	}

	return result
}

func (c challenge) getAdjacent(data [][]string, x int, y int) []string {

	var dxStart, dyStart, dxEnd, dyEnd int
	if x > 0 {
		dxStart = -1
	} else {
		dxStart = 0
	}

	if y > 0 {
		dyStart = -1
	} else {
		dyStart = 0
	}

	if x < len(data)-1 {
		dxEnd = 1
	} else {
		dxEnd = 0
	}

	if y < len(data[0])-1 {
		dyEnd = 1
	} else {
		dyEnd = 0
	}

	var result []string
	for dx := dxStart; dx <= dxEnd; dx++ {
		for dy := dyStart; dy <= dyEnd; dy++ {
			if dx != 0 || dy != 0 {
				result = append(result, data[x+dx][y+dy])
			}
		}
	}

	return result

}

func (c challenge) partA(s [][]string) [][]string {

	alter := deepCopy(s)

	for x, i := range s {
		for y, j := range i {
			adj := c.getAdjacent(s, x, y)
			isEmpty := j == "L"
			isOccupied := j == string(rune(35))
			allAdjFree := !stringInSlice(string(rune(35)), adj)
			fourAdjOccupied := (numInSlice(string(rune(35)), adj) >= 4)

			if isEmpty && allAdjFree {
				alter[x][y] = string(rune(35))
			}
			if isOccupied && fourAdjOccupied {
				alter[x][y] = "L"
			}

		}
	}

	if !reflect.DeepEqual(s, alter) {
		alter = c.partA(alter)
	}

	return alter

}
func numInSlice(a string, list []string) int {

	var count int
	for _, b := range list {
		if b == a {
			count++
		}
	}
	return count
}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (c challenge) partB(s [][]string) [][]string {

}

func main() {

	day10 := challenge{env: "prod"}
	day10.input = day10.readInput()
	ansA := day10.partA(deepCopy(day10.input))

	var count int
	for _, i := range ansA {
		for _, j := range i {
			if j == string(rune(35)) {
				count++
			}
		}
	}

	fmt.Println("Answer to Part A: ", count)

}
