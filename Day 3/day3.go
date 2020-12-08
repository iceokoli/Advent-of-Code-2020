package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

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

func partA(data []string) int {

	var route string
	for i, r := range data {
		var right int
		if i*3 < len(r) {
			right = i * 3
		} else {
			right = (i * 3) % len(r)
		}
		if i != 0 {
			route += string(r[right])
		}
	}
	return strings.Count(route, string(rune(35)))
}

func partB(data []string, right int, down int) int {

	var route string
	for i, r := range data {
		var move int
		if (i*right)/down < len(r) {
			move = (i * right) / down
		} else {
			move = ((i * right) / down) % len(r)
		}

		if i != 0 && i%down == 0 {
			route += string(r[move])
		}
	}
	return strings.Count(route, string(rune(35)))
}

func main() {

	rawData := getInput()

	fmt.Println("Part 1 Answer: ", partA(rawData))

	ansB := partB(rawData, 1, 1) *
		partB(rawData, 3, 1) *
		partB(rawData, 5, 1) *
		partB(rawData, 7, 1) *
		partB(rawData, 1, 2)
	fmt.Println("Part 2 Answer: ", ansB)

}
