package io

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInFile(filename string) ([]int64, []int64) {
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
