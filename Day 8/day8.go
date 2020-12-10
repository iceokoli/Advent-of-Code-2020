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

type code struct {
	instr string
	value int
}

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
	result := strings.Split(string(raw), "\n") //list of strings
	return result
}

func cleanInput(data []string) []code {

	var result []code
	for _, line := range data {
		temp := strings.Split(line, " ")
		inst := temp[0]
		val, _ := strconv.Atoi(temp[1])
		result = append(result, code{inst, val})
	}

	return result
}

func checkExecutes(n int, execs []int) bool {

	// check if there are more than n executes

	for _, i := range execs {
		if i > n {
			return true
		}
	}
	return false
}

func runTillSecond(data []code) (int, []int, bool) {

	accum := 0
	noExecutes := make([]int, len(data))
	order := make([]int, len(data))
	counter := 0
	exitLoop := false

	i := 0
	for counter < len(data) {
		i++
		temp := data[counter]
		noExecutes[counter]++
		exitLoop = checkExecutes(1, noExecutes)
		if exitLoop {
			break
		} //exit on second execute
		order[counter] = i
		if temp.instr == "acc" {
			accum += temp.value
		} else if temp.instr == "jmp" {
			counter += temp.value
			continue
		}
		counter++

	}
	return accum, order, exitLoop
}

func verifyFix(data []code) bool {

	_, _, result := runTillSecond(data)

	return !result
}

func fixCode(data []code) []code {

	toReplace := map[string]string{
		"jmp": "nop", "nop": "jmp",
	}
	var result []code
	for k, v := range toReplace {
		end := false
		for idx, cd := range data {
			if cd.instr == k {
				temp := make([]code, len(data))
				copy(temp, data)
				temp[idx].instr = v
				if verifyFix(temp) {
					result = temp
					end = true
					break
				}
			}
		}
		if end {
			break
		}

	}
	return result
}

func main() {
	raw := readInput()
	clean := cleanInput(raw)
	acc, _, _ := runTillSecond(clean)
	fmt.Println("Part A: ", acc)
	fixed := fixCode(clean)
	acc2, _, _ := runTillSecond(fixed)
	fmt.Println("Part B: ", acc2)
}
