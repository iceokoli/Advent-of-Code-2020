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

func readInput() []string {
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
	result := strings.Split(string(raw), "\n\n") //list of strings

	return result
}

func returnUnique(data string) []string {
	check := make(map[string]int)
	for _, val := range data {
		check[string(val)] = 1
	}
	var res []string
	for l, _ := range check {
		res = append(res, l)
	}
	return res

}

func countQuestionsA(q []string) (int, [][]string) {

	// create an array of sets
	var set [][]string
	for _, i := range q {
		temp := strings.ReplaceAll(i, "\n", "")

		res := returnUnique(temp)

		set = append(set, res)
	}

	// count the questions
	cnt := 0
	for _, i := range set {
		cnt = cnt + len(i)
	}

	return cnt, set

}

func countQuestionsB(q []string, s [][]string) int {

	// format my data
	var staging [][]string
	for _, i := range q {
		staging = append(staging, strings.Split(i, "\n"))
	}

	//remove duplicates
	var data [][]string
	for _, i := range staging {
		var group []string
		for _, j := range i {
			line := strings.Join(returnUnique(j), "")
			group = append(group, line)
		}
		data = append(data, group)
	}

	// count the questions
	cnt := 0
	for idx, i := range data {

		temp := strings.Join(i, "")
		for _, j := range s[idx] {
			if len(i) == strings.Count(temp, string(j)) {
				cnt = cnt + 1
			}
		}
	}

	return cnt

}

func main() {
	raw := readInput()
	c, s := countQuestionsA(raw)
	fmt.Println("Number of elements ", c)
	fmt.Println("Number of elements ", countQuestionsB(raw, s))
}
