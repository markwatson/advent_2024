package day1

import (
	"fmt"
	"log"
	"math"
	"slices"

	"github.com/markwatson/advent_2024/pkg/util"
)

type Day1 struct {

}

func (d Day1) Run(inputFile string) {
	// read
	left, right := readInFile(inputFile)

	// find the difference
	diff := findDiff(left, right)
	fmt.Printf("Total difference: %d\n", diff)

	// find similarity score
	similarity := findSimilarity(left, right)
	fmt.Printf("Similarity score: %d\n", similarity)
}

func readInFile(filename string) ([]int64, []int64) {
	// Read in numbers
	contents, err := util.ReadInNumbers(filename, int64(0))
	util.Check(err)

	// Outputs
	var leftValues []int64
	var rightValues []int64

	// Parse
	for _, row := range contents {
		if len(row) != 2 {
			log.Fatalf("Invalid row: %v", row)
		}
		leftValues = append(leftValues, row[0])
		rightValues = append(rightValues, row[1])
	}

	return leftValues, rightValues
}

func findDiff(left []int64, right []int64) int64 {
	if len(left) != len(right) {
		log.Fatal("Error: left and right arrays are not the same length")
	}
	problemLen := len(left)

	slices.Sort(left)
	slices.Sort(right)

	// calculate the difference
	var totalDiff int64
	for i := 0; i < problemLen; i++ {
		diff := math.Abs(float64(left[i]) - float64(right[i]))
		totalDiff += int64(diff)
	}

	return totalDiff
}

// This is overly optimized and not necessary for this problem
func findSimilarity(left []int64, right []int64) int64 {
	// calculate occurrences of each number
	occurrences := make(map[int64]int64)
	for _, val := range right {
		occurrences[val] += 1
	}

	var totalScore int64
	for _, val := range left {
		i, ok := occurrences[val]
		if ok {
			totalScore += val * i
		} // Otherwise don't increase score
	}

	return totalScore
}
