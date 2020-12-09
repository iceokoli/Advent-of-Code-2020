package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func appendNoDupes(a []string, b []string) []string {

	check := make(map[string]int)
	d := append(a, b...)
	res := make([]string, 0)
	for _, val := range d {
		check[val] = 1
	}

	for letter := range check {
		res = append(res, letter)
	}

	return res
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

func formatData(data []string) map[string][]string {

	res := make(map[string][]string)
	for _, i := range data {
		var idx = regexp.MustCompile("^\\w+\\s\\w+")
		var elm = regexp.MustCompile("[0-9]+\\s\\w+\\s\\w+\\s\\w+[^\\.|,]")
		res[idx.FindString(i)] = elm.FindAllString(i, -1)
	}
	return res
}

func findBag(color string, data map[string][]string) []string {
	// find bag recursive
	var baseColors []string
	for k, v := range data {
		for _, i := range v {
			if strings.Contains(i, color) {
				baseColors = append(baseColors, k)
			}
		}
	}

	var result []string
	if len(baseColors) == 0 {
		return baseColors
	}

	for _, i := range baseColors {
		temp := findBag(i, data)
		result = appendNoDupes(result, temp)
	}

	result = appendNoDupes(result, baseColors)
	return result
}

func findNoBag(color string, data map[string][]string) float64 {
	// count no bag recursive
	var noBag float64
	if len(data[color]) == 0 {
		return 0
	}

	for _, i := range data[color] {
		n, _ := strconv.ParseFloat(i[:1], 64)
		clr := strings.Join(strings.Split(i, " ")[1:3], " ")
		noBag = noBag + (n * findNoBag(clr, data)) + n
	}

	return noBag
}

func main() {

	raw := readInput()
	clean := formatData(raw)
	partA := findBag("shiny gold", clean)
	fmt.Println("Part A: ", len(partA))
	fmt.Println("Part B: ", findNoBag("shiny gold", clean))

}
