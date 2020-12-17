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

func deepCopy(b [][]rune) [][]rune {
	a := make([][]rune, len(b))
	for i := range b {
		a[i] = make([]rune, len(b[i]))
		copy(a[i], b[i])
	}
	return a

}

type challenge struct {
	env   string
	input [][]rune
}

func (c challenge) readInput() [][]rune {
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
	var result [][]rune
	for _, i := range strings.Split(string(raw), "\n") {
		temp := []rune(i)
		result = append(result, temp)
	}

	return result
}

func (c challenge) makeChanges(s [][]rune) [][]rune {
	result := deepCopy(s)

	for r, i := range result {
		for c, j := range i {

			up := s[r-1][c]
			down := s[r+1][c]
			d1 := s[r-1][c-1]
			d2 := s[r-1][c+1]
			left := s[r][c-1]
			right := s[r][c+1]
			
			d3 := s[r+1][c-1]
			d4 := s[r+1][c+1]
			if string(j) == "L" && 
		}
	}

}

func (c challenge) partA() {

	start := deepCopy(c.input)
	result := c.makeChanges(start)

}

func main() {

	day10 := challenge{env: "test"}
	day10.input = day10.readInput()
	day10.partA()

}
