package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func readInput() [][]string {
	// reads the file and structures it as an arrayß
	_, filename, _, _ := runtime.Caller(0)
	dir, _ := filepath.Abs(filepath.Dir(filename))
	raw, _ := ioutil.ReadFile(dir + "/input.txt")

	var result [][]string
	for _, i := range strings.Split(string(raw), "\n\n") {
		temp := strings.Split(strings.ReplaceAll(i, "\n", " "), " ")
		result = append(result, temp)
	}

	return result
}

func validatePassportsA(passports [][]string) [][]string {
	// checks the passports have the correct number of fields
	var valid [][]string
	for _, passport := range passports {
		c1 := strings.Contains(strings.Join(passport, " "), "cid")
		if len(passport) == 8 {
			valid = append(valid, passport)
		} else if len(passport) == 7 && c1 == false {
			valid = append(valid, passport)
		}
	}
	return valid
}

type passport struct {
	pid string
	ecl string
	hcl string
	hgt string
	eyr int
	iyr int
	byr int
}

func structurePassports(passports [][]string) []passport {
	// Turn each passport into a struct
	var result []passport
	for _, pass := range passports {
		var tempPassport passport
		for _, field := range pass {
			temp := strings.Split(field, ":")
			switch temp[0] {
			case "pid":
				tempPassport.pid = temp[1]
			case "ecl":
				tempPassport.ecl = temp[1]
			case "hcl":
				tempPassport.hcl = temp[1]
			case "hgt":
				tempPassport.hgt = temp[1]
			case "eyr":
				tempPassport.eyr, _ = strconv.Atoi(temp[1])
			case "iyr":
				tempPassport.iyr, _ = strconv.Atoi(temp[1])
			case "byr":
				tempPassport.byr, _ = strconv.Atoi(temp[1])

			}
		}
		result = append(result, tempPassport)
	}
	return result
}

func validateFields(passports []passport) []passport {
	// validate the fields
	var valid []passport
	for _, pass := range passports {
		// checks
		var c1 bool
		eclr := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, i := range eclr {
			if pass.ecl == i {
				c1 = true
			}
		}
		c2 := (pass.byr >= 1920 && pass.byr <= 2002)
		c3 := (pass.iyr >= 2010 && pass.iyr <= 2020)
		c4 := (pass.eyr >= 2020 && pass.eyr <= 2030)

		var c5 bool
		if strings.Contains(pass.hgt, "in") {
			h, _ := strconv.Atoi(string(pass.hgt[:len(pass.hgt)-2]))
			c5 = (h >= 59 && h <= 76)
		} else {
			h, _ := strconv.Atoi(string(pass.hgt[:len(pass.hgt)-2]))
			c5 = (h >= 150 && h <= 193)
		}

		var validClr = regexp.MustCompile("^#[0-9|a-f]{6}$")
		c6 := validClr.MatchString(pass.hcl)

		var validPid = regexp.MustCompile("^[0-9]{9}$")
		c7 := validPid.MatchString(pass.pid)

		if c1 && c2 && c3 && c4 && c5 && c6 && c7 {
			valid = append(valid, pass)
		}

	}
	return valid
}

func main() {

	passports := readInput()
	validPassportsA := validatePassportsA(passports)
	fmt.Println("Answer to Part 1: ", len(validPassportsA))
	sValidPassportsA := structurePassports(validPassportsA)
	fmt.Println("Answer B: ", len(validateFields(sValidPassportsA)))
}
