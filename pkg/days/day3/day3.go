package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/markwatson/advent_2024/pkg/util"
)

// Template for adding a new day

type Day3 struct {
	regex           *regexp.Regexp
	regexWithEnable *regexp.Regexp
}

const mulFirst = 1
const mulSecond = 2
const dont = 3
const do = 4

func NewDay3() (*Day3, error) {
	d := new(Day3)
	r, err := regexp.Compile(`mul\((\d{1,3})\,(\d{1,3})\)`)
	if err != nil {
		return nil, err
	}
	d.regex = r

	r2, err := regexp.Compile(`mul\((\d{1,3})\,(\d{1,3})\)|(don't\(\))|(do\(\))`)
	if err != nil {
		return nil, err
	}
	d.regexWithEnable = r2

	return d, nil
}

func (d Day3) sumMultiples(input string) int {
	n := 0

	matches := d.regex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		a := util.CheckFatal(strconv.Atoi(match[mulFirst]))
		b := util.CheckFatal(strconv.Atoi(match[mulSecond]))
		n += a * b
	}

	return n
}

func (d Day3) sumMultiplesPart2(input string) int {
	n := 0

	matches := d.regexWithEnable.FindAllStringSubmatch(input, -1)
	enabled := true
	for _, match := range matches {
		if match[dont] != "" {
			enabled = false
		} else if match[do] != "" {
			enabled = true
		} else if enabled {
			a := util.CheckFatal(strconv.Atoi(match[mulFirst]))
			b := util.CheckFatal(strconv.Atoi(match[mulSecond]))
			n += a * b
		}
	}

	return n
}

func (d Day3) Run(inputFile string) {
	data := util.CheckFatal(util.ReadString(inputFile))
	total := d.sumMultiples(data)
	fmt.Printf("Total part 1: %d\n", total)
	total2 := d.sumMultiplesPart2(data)
	fmt.Printf("Total part 2: %d\n", total2)
}
