package day6

import (
	"fmt"

	"github.com/markwatson/advent_2024/pkg/util"
)

type Day6 struct {

}

const iconGuard = '^' // I think it always starts with going up
const iconWall = '#'
const iconFloor = '.'
const iconVisited = 'X'
const (
	up = iota
	right
	down
	left
)

func findGuard(data *util.Matrix) util.Position {
	var guard util.Position
	for r, row := range data.Data {
		for c, cell := range row {
			if cell == iconGuard {
				guard = util.Position{R: r, C: c}
				break
			}
		}
	}
	return guard
}

// Walks the path, returns true if it loops.
func walkPath(data *util.Matrix) bool {
	// find guard
	guard := findGuard(data)

	// walking
	direction := up
	pos := guard
	revisited := 0 // for now a simple counter
walkabout:
	for {
		toGo := util.Position{}
		switch direction {
		case up:
			toGo = util.Position{R: pos.R - 1, C: pos.C}
		case right:
			toGo = util.Position{R: pos.R, C: pos.C + 1}
		case down:
			toGo = util.Position{R: pos.R + 1, C: pos.C}
		case left:
			toGo = util.Position{R: pos.R, C: pos.C - 1}
		}
		toGoItem := data.Get(toGo.R, toGo.C)

		switch toGoItem {
		case iconFloor:
			data.Data[pos.R][pos.C] = iconVisited
			pos = toGo
		case iconVisited:
			revisited++
			data.Data[pos.R][pos.C] = iconVisited
			pos = toGo
		case iconWall:
			direction = (direction + 1) % 4
		case 0: // Out of the map
			data.Data[pos.R][pos.C] = iconVisited
			break walkabout
		}

		// For part 2
		if revisited > 1000 {
			return true
		}
	}

	return false
}

func positionsVisited(data *util.Matrix) int {
	//fmt.Println("Debugging, visited: ")
	//fmt.Println(data)
	positionsVisited := 0
	for _, row := range data.Data {
		for _, cell := range row {
			if cell == iconVisited {
				positionsVisited++
			}
		}
	}
	return positionsVisited
}

func part1(data []string) {
	matrix := util.NewMatrix(data)
	walkPath(matrix)
	visitedCells := positionsVisited(matrix)
	fmt.Printf("Visited cells: %d\n", visitedCells)
}

// This is not a super satisfying solution. It relies on a modification
// to part one to calculate loops by just counting the number of squares
// revisited, and then breaking on a threshold. Then we just add an obstruction
// in each floor tile, walk the path, and count it if it loops.
//
// A better way would probably be to build a graph/tree of the visited cells
// instead of doing graph traversal, and then just look for cycles. It would
// probably be more compact, but I started with matrix walking in a byte
// array, so I kinda stuck with it.
func part2BruteForce(data []string) {
	loopedPositions := make([]util.Position, 0)
	for r, row := range data {
		for c, cell := range row {
			if cell == iconFloor {
				matrix := util.NewMatrix(data)
				matrix.Data[r][c] = iconWall
				if walkPath(matrix) {
					loopedPositions = append(loopedPositions, util.Position{R: r, C: c})
				}
			}
		}
	}

	fmt.Printf("Looped positions: %d\n", len(loopedPositions))
}



func (d Day6) Run(inputFile string) {
	data := util.CheckFatal(util.ReadLines(inputFile))

	// Part 1
	part1(data)

	// Part 2
	part2BruteForce(data)
}
