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

type policy struct {
	min      int
	max      int
	letter   string
	password string
}

func validatePolicyA(data []policy) []policy {

	var result []policy
	for _, i := range data {
		nLetters := strings.Count(i.password, i.letter)
		if nLetters >= i.min && nLetters <= i.max {
			result = append(result, i)
		}

	}

	return result
}

func validatePolicyB(data []policy) []policy {
	var result []policy
	for _, i := range data {
		first := string([]rune(i.password)[i.min-1])
		second := string([]rune(i.password)[i.max-1])
		if strings.Count(first+second, i.letter) == 1 {
			result = append(result, i)
		}
	}

	return result
}

func main() {

	// Get current directory
	_, filename, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}

	// read in the file
	raw, err := ioutil.ReadFile(dir + "/day2_input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// turn the file into an array
	staging := strings.Split(string(raw), "\n") //list of strings
	// convert to a list of structs policy
	var data []policy
	for _, line := range staging {
		temp := strings.Split(line, " ")
		min, _ := strconv.Atoi(strings.Split(temp[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(temp[0], "-")[1])
		result := policy{
			min,
			max,
			string([]rune(temp[1])[0]),
			temp[2],
		}
		data = append(data, result)

	}

	validPolicyA := validatePolicyA(data)
	fmt.Println("Number of Valid Policies: ", len(validPolicyA))

	validPolicyB := validatePolicyB(data)
	fmt.Println("Number of Valid Policies: ", len(validPolicyB))

}
