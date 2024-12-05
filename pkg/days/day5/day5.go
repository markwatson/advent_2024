package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/markwatson/advent_2024/pkg/util"
)

// This problem was fun. I started by writing a super slow implementation, despite the warnings in
// the problem text. It ended up being fast enough though:
// time go run main.go 5 input
// Total: 5108
// Total part 2: 7380
//
// real	0m0.464s
// user	0m0.163s
// sys	0m0.221s
//
// Because I wrote the first part in such an odd way, part 2 was a bit harder. I ended up
// using a simple topological sort function to reorder based on the filtered rules for
// each list of page numbers. This would need to be modified a bit if there are any pages
// not included in the rules.
//
// My input parsing code is kinda garbage. Go feels a lot more verbose, plus using
// copilot results in generating a lot of code without realizing it.

type Day5 struct {

}

type Pair struct {
	left int
	right int
}


func readInput(inputFile string) ([]Pair, [][]int) {
	orderingRules := make([]Pair, 0)
	printables := make([][]int, 0)

	lines := util.CheckFatal(util.ReadLines(inputFile))
	for _, line := range lines {
		ordering := strings.Split(line, "|")
		if len(ordering) == 2 {
			left := util.CheckFatal(strconv.Atoi(ordering[0]))
			right := util.CheckFatal(strconv.Atoi(ordering[1]))

			orderingRules = append(orderingRules, Pair{left, right})
		} else {
			pages := strings.Split(line, ",")
			printable := make([]int, 0)

			if len(pages) > 0 {
				for _, page := range pages {
					printable = append(printable, util.CheckFatal(strconv.Atoi(page)))
				}
				printables = append(printables, printable)
			}
		}
	}

	return orderingRules, printables
}

func filterValidRules(orderingRules []Pair, printable []int) []Pair {
	validRules := make([]Pair, 0)
	for _, rule := range orderingRules {
		if slices.Contains(printable, rule.left) && slices.Contains(printable, rule.right) {
			validRules = append(validRules, rule)
		}
	}
	return validRules
}

// Slow way (maybe using topological sort and then comparing would be faster?)
// I didn't write the graph based sort until I started part 2,
// so that's why this just walks up and down the list for every rule to check it.
func isOrdered(orderingRules []Pair, printable []int) bool {
	validRules := filterValidRules(orderingRules, printable)

	for page := 0; page < len(printable); page++ {
		for _, rule := range validRules {
			if rule.right == printable[page] {
				good := false
				for prev := page - 1; prev >= 0; prev-- {
					if printable[prev] == rule.left {
						good = true
						break
					}
				}
				if !good {
					return false
				}
			}
		}
		
		
	}
	
	return true
}

// Order items - ignore the initial list and repopulate it based on the filtered rules.
func topologicalSort(orderingRules []Pair, printable []int) []int {
	validRules := filterValidRules(orderingRules, printable)

	// build a dependency graph to make things easier
	graph := make(map[int][]int)
	for _, rule := range validRules {
		graph[rule.left] = append(graph[rule.left], rule.right)
	}

	inDegree := make(map[int]int)
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	result := make([]int, 0)
	queue := make([]int, 0)
	for node := range graph {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	for {
		if len(queue) == 0 {
			break
		}

		node := queue[0]
		queue = queue[1:]

		result = append(result, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}


func findCenter(printable []int) int {
	return printable[len(printable) / 2]
}

func (d Day5) Run(inputFile string) {
	orderingRules, printables := readInput(inputFile)

	validPrintables := make([][]int, 0)
	invalidPrintables := make([][]int, 0)
	for _, printable := range printables {
		if isOrdered(orderingRules, printable) {
			validPrintables = append(validPrintables, printable)
		} else {
			invalidPrintables = append(invalidPrintables, printable)
		}
	}

	var total int
	for _, printable := range validPrintables {
		total += findCenter(printable)
	}
	fmt.Printf("Total: %d\n", total)

	// Part 2
	var total2 int
	for _, printable := range invalidPrintables {
		sorted := topologicalSort(orderingRules, printable)
		total2 += findCenter(sorted)
	}
	fmt.Printf("Total part 2: %d\n", total2)
}

