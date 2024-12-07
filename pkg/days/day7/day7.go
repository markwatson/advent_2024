package day7

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/markwatson/advent_2024/pkg/util"
)

// This is a little bit of a weird one. I decided to just brute force it. To do
// this I just went with generating every possibility and evaluating the formula. This
// is easy for part one since there are two operators, so we can use binary and bit
// shifting. For part two I decided to use a more traditional recursive function. This
// worked, and proved to be fast enough that I decided not to optimize it.
type Day7 struct {}

const (
	plus = iota
	multiply
	concat
)

type Calibration struct {
	testValue int
	numbers []int
}

func (c Calibration) String() string {
	return fmt.Sprintf("%d: %v", c.testValue, c.numbers)
}

func parse(data []string) []Calibration {
	calibrations := make([]Calibration, 0)
	for _, line := range data {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Printf("Invalid line: %s", line)
			os.Exit(1)
		}

		testValue := util.CheckFatal(strconv.Atoi(parts[0]))

		numbers := make([]int, 0)
		numStrings := strings.Split(parts[1], " ")
		for _, numString := range numStrings {
			if len(numString) == 0 {
				continue
			}
			numbers = append(numbers, util.CheckFatal(strconv.Atoi(numString)))
		}

		calibrations = append(calibrations, Calibration{testValue: testValue, numbers: numbers})
	}

	return calibrations
}

func (c Calibration) isValidPart1() bool {
	slots := len(c.numbers) - 1
	combinations := int(math.Pow(2, float64(slots)))
	
	
	for i := 0; i < combinations; i++ {		
		total := c.numbers[0]
		for j := 0; j < slots; j++ {
			if (i>>j)&1 == 1 {
				total = total * c.numbers[j+1]
			} else {
				total = total + c.numbers[j+1]
			}
		}

		if total == c.testValue {
			return true
		}
	}

	return false
}

func evalCombo(combo []int, numbers []int) int {
	total := numbers[0]

	for i := 0; i < len(combo); i++ {
		next := numbers[i+1]

		if combo[i] == multiply {
			total *= next
		} else if combo[i] == plus {
			total += next
		} else if combo[i] == concat {
			combined, err := strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(next))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			total = combined
		}
	}
	
	return total
}

// Maybe use dynamic programming here?
func (c Calibration) genCombosRecur(combo []int, index int) bool {
	if index == len(combo) {
		result := evalCombo(combo, c.numbers)
		if result == c.testValue {
			return true
		} else {
			return false
		}
	}

	result := false
	for i := 0; i < 3; i++ { // 3 ops
		combo[index] = i
		if c.genCombosRecur(combo, index+1) {
			result = true
		}
	}
	return result
}

func (c Calibration) isValidPart2() bool {
	slots := len(c.numbers) - 1
	combination := make([]int, slots)

	return c.genCombosRecur(combination, 0)
}

func (d Day7) Run(inputFile string) {
	data := util.CheckFatal(util.ReadLines(inputFile))
	calibrations := parse(data)

	total := 0
	totalPart2 := 0
	for _, calibration := range calibrations {
		// fmt.Println(calibration)

		if calibration.isValidPart1() {
			total += calibration.testValue
			// fmt.Println("==> is valid part 1")
		}

		if calibration.isValidPart2() {
			totalPart2 += calibration.testValue
			// fmt.Println("==> is valid part 2")
		}
	}
	fmt.Printf("\nPart 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", totalPart2)
}
