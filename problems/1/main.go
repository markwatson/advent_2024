package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func readInFile(filename string) ([]int64, []int64) {
	// open file and read line by line
	file, err := os.Open(filename)
    check(err)
	defer file.Close()

	// Outputs
	var leftValues []int64
	var rightValues []int64

	// Read in
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) != 2 {
			// TODO: maybe just ignore?
			log.Fatalf("Ignoring line: %s", line)
		}

		// convert strings to int64
		left, err := strconv.ParseInt(words[0], 10, 64)
		check(err)
		right, err := strconv.ParseInt(words[1], 10, 64)
		check(err)

		// append to arrays
		leftValues = append(leftValues, left)
		rightValues = append(rightValues, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return leftValues, rightValues
}

func solve(left []int64, right []int64) int64 {
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

func main() {
	// check arguments
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
    }
	fname := os.Args[1]

	// read
	left, right := readInFile(fname)

	// solve
	result := solve(left, right)

	// print result
	fmt.Printf("Total difference: %d\n", result)
}
