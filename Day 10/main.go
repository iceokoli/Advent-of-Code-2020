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

func (c challenge) searchOneDiff(l []int, v int) bool {

	for _, i := range l {
		if v+1 == i {
			return true
		}
	}
	return false
}

func (c challenge) searchTwoDiff(l []int, v int) bool {

	for _, i := range l {
		if v+2 == i {
			return true
		}
	}
	return false
}

func (c challenge) searchThreeDiff(l []int, v int) bool {

	for _, i := range l {
		if v+3 == i {
			return true
		}
	}
	return false
}

func (c challenge) partA(data []int) (int, []int) {

	jolt := 0
	jBuiltin := max(data) + 3
	oneDiff := 0
	twoDiff := 0
	threeDiff := 0
	var order []int

	order = append(order, 0)
	for i := 0; i < len(data); i++ {

		if c.searchOneDiff(data, jolt) {
			oneDiff++
			jolt++
		} else if c.searchTwoDiff(data, jolt) {
			twoDiff++
			jolt += 2
		} else {
			threeDiff++
			jolt += 3
		}
		order = append(order, jolt)
	}

	diff := jBuiltin - jolt
	if diff == 1 {
		oneDiff++
	} else if diff == 2 {
		twoDiff++
	} else {
		threeDiff++
	}
	order = append(order, jBuiltin)

	return oneDiff * threeDiff, order
}

var check = make(map[int]int)

func (c challenge) partB(order []int, pos int) (int, map[int]int) {

	var result int

	if pos == len(order)-1 {
		return 1, check
	}

	// if _, ok := check[pos]; ok {
	// 	return check[pos]
	// }

	for i := pos + 1; i < len(order); i++ {
		if order[i]-order[pos] <= 3 {
			t, _ := c.partB(order, i)
			result += t
		}
	}
	check[pos] = result
	return result, check
}

func main() {

	day10 := challenge{env: "prod"}
	raw := day10.readInput()
	ansA, o := day10.partA(raw)
	fmt.Println("Part A ", ansA)
	ansB, _ := day10.partB(o, 0)
	fmt.Println(ansB)
	// fmt.Println(o)
	// fmt.Println(c)

}
