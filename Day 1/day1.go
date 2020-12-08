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

func part1(data []int) (int, int) {

	var x int
	var y int

	for _, i := range data {
		for _, j := range data {
			if i+j == 2020 {
				x = i
				y = j
			}
		}
	}

	return x, y

}

func part2(data []int) (int, int, int) {

	var x int
	var y int
	var z int

	for _, i := range data {
		for _, j := range data {
			for _, k := range data {
				if i+j+k == 2020 {
					x = i
					y = j
					z = k
				}
			}
		}
	}

	return x, y, z

}

func main() {

	// Get current directory
	_, filename, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}

	// read in the file
	raw, err := ioutil.ReadFile(dir + "/day1_input.txt")

	// error if there are any issues
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// turn the file into an array
	staging := strings.Split(string(raw), "\n") //list of strings
	// convert to a list of numbers
	var data []int
	for _, s := range staging {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		data = append(data, n)
	}
	// find 2 entries that add to 2020 and multiply them
	entry1, entry2 := part1(data)
	fmt.Println("Answer to Part 1:", entry1*entry2)

	// find 3 entries that add to 2020 and multiply them
	entry1, entry2, entry3 := part2(data)
	fmt.Println("Answer to Part 2:", entry1*entry2*entry3)

}
