package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/markwatson/advent_2024/pkg/util"
)

type Level int64
type Report []Level

func readInFile(filename string) ([]Report) {
	contents, err := util.ReadInNumbers(filename, int64(0))
	util.Check(err)

	var reports []Report
	for _, row := range contents {
		var report Report
		for _, val := range row {
			report = append(report, Level(val))
		}
		reports = append(reports, report)
	}

	return reports
}

// A report is safe if both of the following are true:
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.
func isSafe(report Report) bool {
	log.Printf("## Checking report: %v\n", report)
	// Check if increasing
	increasing := true
	for i := 1; i < len(report); i++ {
		if report[i] < report[i-1] {
			increasing = false
			break
		}
	}

	// Check if decreasing
	decreasing := true
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			decreasing = false
			break
		}
	}

	// Check if adjacent levels differ by at least one and at most three
	differ := true
	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i]) - float64(report[i-1]))
		if diff < 1 || diff > 3 {
			differ = false
			break
		}
	}

	log.Printf("Increasing: %v, Decreasing: %v, Differ: %v\n", increasing, decreasing, differ)

	return (increasing || decreasing) && differ
}

func remove(slice []Level, s int) []Level {
    return append(slice[:s], slice[s+1:]...)
}

// A report is safe if both of the following are true:
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.
// But we also "tolerate a single bad level".
func isSafeAllowance(report Report) bool {
	// is safe
	if isSafe(report) {
		return true
	}

	// allowance
	for i := 0; i < len(report); i++ {
		// TODO: this does two copies - inefficient
		reportPrime := make(Report, len(report))
		copy(reportPrime, report)
		reportPrime = remove(reportPrime, i)

		if isSafe(reportPrime) {
			return true
		}
	}

	return false
}

func countSafe(reports []Report) (int, int) {
	count := 0
	countAllowance := 0
	for _, report := range reports {
		if isSafe(report) {
			log.Printf("=> Safe report: %v\n", report)
			count++
		} else if isSafeAllowance(report) {
			log.Printf("=> Safe report with allowance: %v\n", report)
			countAllowance++
		} else {
			log.Printf("=> Unsafe report: %v\n", report)
		}
	}
	return count, countAllowance + count
}

func main() {
	// check arguments
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
    }
	fname := os.Args[1]

	// read
	reports := readInFile(fname)

	// part 1
	count, countAllowance := countSafe(reports)
	fmt.Printf("\n\nNumber of safe reports: %d\n", count)

	// part 2
	fmt.Printf("Number of safe reports with allowance: %d\n", countAllowance)
}

