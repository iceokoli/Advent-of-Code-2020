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

func min(v []int) int {
	var m int
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func max(v []int) int {
	var m int
	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

type challenge struct {
	env      string
	preamble int
}

func (c challenge) readInput() []int {
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
	var result []int
	for _, i := range strings.Split(string(raw), "\n") {
		temp, _ := strconv.Atoi(i)
		result = append(result, temp)
	}
	return result
}

func (c challenge) check(v int, l []int) bool {
	for _, x := range l {
		for _, y := range l {
			if x+y == v {
				return true
			}
		}
	}
	return false
}

func (c challenge) findWeaknessA(data []int) int {

	var result int
	start := c.preamble + 1
	for i := start; i < len(data); i++ {
		if t := c.check(data[i], data[i-start:i]); !t {
			result = data[i]
		}
	}
	return result
}

func (c challenge) addUpPreamble(v int, l []int) (int, int) {

	var start int
	var end int
	for i := 0; i < len(l); i++ {
		counter := 0
		j := i + 1
		for counter != v && j <= len(l) {
			j++
			counter = 0
			for _, x := range l[i:j] {
				counter += x
			}

		}
		if counter == v {
			start = i
			end = j
			break
		}
	}
	return start, end
}

func (c challenge) findWeaknessB(data []int) int {

	var s, e int
	start := c.preamble + 1
	for i := start; i < len(data); i++ {
		if t := c.check(data[i], data[i-start:i]); !t {
			s, e = c.addUpPreamble(data[i], data)
		}
	}

	contigious := data[s:e]
	result := min(contigious) + max(contigious)

	return result
}

func main() {

	day9 := challenge{env: "prod"}
	if day9.env == "test" {
		day9.preamble = 5
	} else {
		day9.preamble = 25
	}
	raw := day9.readInput()
	w := day9.findWeaknessA(raw)
	fmt.Println("Answer Part A: ", w)
	fmt.Println(day9.findWeaknessB(raw))

}
