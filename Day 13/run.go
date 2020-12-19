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

func getData(env string) (int, []int) {
	// Get current directory
	_, filename, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}
	// read in the file
	raw, err := ioutil.ReadFile(dir + "/" + env + ".txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// turn the file into an array
	staging := strings.Split(string(raw), "\n")
	start, _ := strconv.Atoi(staging[0])
	var buses []int
	for _, bus := range strings.Split(staging[1], ",") {
		temp, _ := strconv.Atoi(bus)
		buses = append(buses, temp)
	}

	return start, buses
}

func main() {

	env := "test"
	start, buses := getData(env)
	fmt.Println(start)
	fmt.Println(buses)

}
