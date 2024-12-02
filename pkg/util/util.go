package util

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

// If an error is not nil, log and quit.
func Check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

// Reads a table of numbers from a file. For example:
// 1 2 3
// 4 5 6
//
// Would return:
// [[1, 2, 3], [4, 5, 6]]
// 
// The type of the numbers read in is determined by the type of v.
// The type of v must be int64 or float64, unless we extend this.
func ReadInNumbers[V int64 | float64](filename string, v V) ([][]V, error) {
	// open file and read line by line
	file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
	defer file.Close()

	// Outputs
	var values [][]V

	// Read in
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) > 0 {
			var row []V
			for _, word := range words {
				switch any(v).(type) {
				case int64:
					value, err := strconv.ParseInt(word, 10, 64)
					if err != nil {
						return nil, err
					}
					row = append(row, V(value))
				case float64:
					value, err := strconv.ParseFloat(word, 64)
					if err != nil {
						return nil, err
					}
					row = append(row, V(value))
				default:
					return nil, errors.New("unsupported type")
				}
			}
			values = append(values, row)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return values, nil
}
