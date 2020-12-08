package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
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
	raw, err := ioutil.ReadFile(dir + "/test.txt")
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
		var idx = regexp.MustCompile("^\\w+\\s\\w+\\s\\w+")
		var elm = regexp.MustCompile("[0-9]+\\s\\w+\\s\\w+\\s\\w+[^\\.|,]")
		res[idx.FindString(i)] = elm.FindAllString(i, -1)
	}
	return res
}

func findGoldBag(data map[string][]string) int {
	// find surface level gold bags
	var cnt int
	for _, v := range data {
		for _, i := range v {
			if strings.Contains(i, "shiny gold") {
				cnt = cnt + 1
			}
		}
	}

	// find nested bags

	return cnt
}

func main() {

	raw := readInput()
	clean := formatData(raw)
	fmt.Println(findGoldBag(clean))

}
