package day8

import (
	"fmt"
	"slices"

	"github.com/markwatson/advent_2024/pkg/util"
)

// This day was kinda tricky - I'm bad at geometry. I ended up using
// a very hacky solution that got hackier in the second section. Instead
// of using a formula I just used the rise / run and kept adding it to
// itself. There are simpler ways to do this, so maybe I'll revisit.
type Day8 struct {}

type ApxPoint struct {
	X float64
	Y float64
}

type Point struct {
	X int
	Y int
}

func findAllPoints(data [][]rune) map[rune][]Point {
	points := make(map[rune][]Point)
	for y, row := range data {
		for x, cell := range row {
			if cell != '.' && cell != '#' {
				points[cell] = append(points[cell], Point{X: x, Y: y})
			}
		}
	}
	return points
}

func calcAntinodes(p1 Point, p2 Point) []Point {
	rise := p2.Y - p1.Y
	run := p2.X - p1.X
	m := float64(rise) / float64(run)
	
	//fmt.Printf("Rise: %d, Run: %d, M: %f\n", rise, run, m)
	if m == 0 {
		first := Point{
			X: p1.X - run,
			Y: p1.Y,
		}
		second := Point{
			X: p2.X + run,
			Y: p2.Y,
		}
		return []Point{first, second}
	} else if m > 0 {
		first := Point{
			X: max(p1.X, p2.X) + run,
			Y: max(p1.Y, p2.Y) + rise,
		}
		second := Point{
			X: min(p1.X, p2.X) - run,
			Y: min(p1.Y, p2.Y) - rise,
		}
		return []Point{first, second}
	} else if (m < 0) {
		first := Point{
			X: min(p1.X, p2.X) + run,
			Y: max(p1.Y, p2.Y) + rise,
		}
		second := Point{
			X: max(p1.X, p2.X) - run,
			Y: min(p1.Y, p2.Y) - rise,
		}
		return []Point{first, second}
	} else {// inf
		first := Point{
			X: p1.X,
			Y: p1.Y + rise,
		}
		second := Point{
			X: p2.X,
			Y: p2.Y - rise,
		}
		return []Point{first, second}
	}
}

func calcResAntinodes(p1 Point, p2 Point, yMax int, xMax int) []Point {
	rise := p2.Y - p1.Y
	run := p2.X - p1.X
	m := float64(rise) / float64(run)

	points := make([]Point, 0)

	minX := min(p1.X, p2.X)
	minY := min(p1.Y, p2.Y)
	maxX := max(p1.X, p2.X)
	maxY := max(p1.Y, p2.Y)

	//fmt.Printf("Rise: %d, Run: %d, M: %f\n", rise, run, m)
	if m == 0 {
		for x := p1.X - run; x > 0; x -= run {
			points = append(points, Point{X: x, Y: p1.Y})
		}
		for x := p2.X + run; x < xMax; x += run {
			points = append(points, Point{X: x, Y: p2.Y})
		}
	} else if m > 0 {
		for {
			maxX += run
			maxY += rise
			points = append(points, Point{X: maxX, Y: maxY})
			if maxX >= xMax || maxY >= yMax {
				break
			}
		}
		for {
			minX -= run
			minY -= rise
			points = append(points, Point{X: minX, Y: minY})
			if minX <= 0 || minY <= 0 {
				break
			}
		}
	} else if (m < 0) {
		for {
			minX += run
			maxY += rise
			points = append(points, Point{X: minX, Y: maxY})
			if minX >= xMax || maxY >= yMax {
				break
			}
		}
		for {
			maxX -= run
			minY -= rise
			points = append(points, Point{X: maxX, Y: minY})
			if maxX <= 0 || minY <= 0 {
				break
			}
		}
	} else {// inf
		for y := p1.Y + rise; y < yMax; y += rise {
			points = append(points, Point{X: p1.X, Y: y})
		}
		for y := p2.Y - rise; y > 0; y -= rise {
			points = append(points, Point{X: p2.X, Y: y})
		}
	}

	return points
}

func filterOutOfRange(allAntinodes []Point, data [][]rune) []Point {
	filteredAntinodes := make([]Point, 0)
	for i := 0; i < len(allAntinodes); i++ {
		if isOutOfBounds(allAntinodes[i], data) {
			continue
		} else {
			filteredAntinodes = append(filteredAntinodes, allAntinodes[i])
		}
	}
	return filteredAntinodes
}

func isOutOfBounds(p Point, data [][]rune) bool {
	if p.X < 0 || p.X >= len(data[0]) {
		return true
	} else if p.Y < 0 || p.Y >= len(data) {
		return true
	} else {
		return false
	}
}

func filterDuplicates(allAntinodes []Point) []Point {
	set := make(map[Point]bool)
	for _, v := range allAntinodes {
		set[v] = true
	}

	uniqueAntinodes := make([]Point, 0)
	for k := range set {
		uniqueAntinodes = append(uniqueAntinodes, k)
	}
	return uniqueAntinodes
}

func findAntinodes(data [][]rune) []Point {
	points := findAllPoints(data)
	
	allAntinodes := make([]Point, 0)
	for _, v := range points {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				antinodes := calcAntinodes(v[i], v[j])
				allAntinodes = append(allAntinodes, antinodes...)
			}
		}
	}

	filteredAntinodes := filterOutOfRange(allAntinodes, data)
	return filterDuplicates(filteredAntinodes)
}

func findResonantAntinodes(data [][]rune) []Point {
	yMax := len(data)
	xMax := len(data[0])
	points := findAllPoints(data)
	resonantAntinodes := make([]Point, 0)
	for _, v := range points {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				antinodes := calcResAntinodes(v[i], v[j], yMax, xMax)
				resonantAntinodes = append(resonantAntinodes, antinodes...)
			}
		}
		// Add in "normal" points
		resonantAntinodes = append(resonantAntinodes, v...)
	}

	resonantAntinodes = filterOutOfRange(resonantAntinodes, data)
	return filterDuplicates(resonantAntinodes)
}

func printGrid(data [][]rune) {
	for _, row := range data {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func (d Day8) Run(inputFile string) {
	data := util.CheckFatal(util.ReadGrid(inputFile))
	fmt.Println("Day 8. Grid:")
	printGrid(data)
	slices.Reverse(data) // Some of the math is based on a reversed grid

	// Part 1
	antinodes := findAntinodes(data)
	fmt.Printf("Part 1: %d\n", len(antinodes))

	// Part 2
	resonantAntinodes := findResonantAntinodes(data)
	fmt.Printf("Part 2: %d\n", len(resonantAntinodes))
}
