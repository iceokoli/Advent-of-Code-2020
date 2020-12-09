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
	raw, err := ioutil.ReadFile(dir + "/test.txt")
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
		fmt.Println(inst, val)
	}

	return result
}

func main() {
	raw := readInput()
	cleanInput(raw)

}
