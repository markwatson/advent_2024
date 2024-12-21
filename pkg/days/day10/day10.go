package day10

import (
	"fmt"

	"github.com/markwatson/advent_2024/pkg/util"
)

type Day10 struct {
}

func (d Day10) Run(inputFile string) {
	data := util.CheckFatal(util.ReadString(inputFile))
	topographicMap := parseTopographicMap(data)
	trailheads := findTrailheads(topographicMap)
	totalScore := 0
	for _, trailhead := range trailheads {
		totalScore += calculateTrailheadScore(topographicMap, trailhead)
	}
	fmt.Println(totalScore)
}

func parseTopographicMap(data string) [][]int {
	lines := util.CheckFatal(util.ReadLines(data))
	topographicMap := make([][]int, len(lines))
	for i, line := range lines {
		topographicMap[i] = make([]int, len(line))
		for j, char := range line {
			topographicMap[i][j] = int(char - '0')
		}
	}
	return topographicMap
}

func findTrailheads(topographicMap [][]int) []util.Position {
	trailheads := []util.Position{}
	for r, row := range topographicMap {
		for c, height := range row {
			if height == 0 {
				trailheads = append(trailheads, util.Position{R: r, C: c})
			}
		}
	}
	return trailheads
}

func calculateTrailheadScore(topographicMap [][]int, trailhead util.Position) int {
	visited := make(map[util.Position]bool)
	return dfs(topographicMap, trailhead, visited)
}

func dfs(topographicMap [][]int, pos util.Position, visited map[util.Position]bool) int {
	if pos.R < 0 || pos.C < 0 || pos.R >= len(topographicMap) || pos.C >= len(topographicMap[0]) {
		return 0
	}
	if visited[pos] {
		return 0
	}
	visited[pos] = true
	if topographicMap[pos.R][pos.C] == 9 {
		return 1
	}
	score := 0
	directions := []util.Position{
		{R: -1, C: 0}, // up
		{R: 1, C: 0},  // down
		{R: 0, C: -1}, // left
		{R: 0, C: 1},  // right
	}
	for _, dir := range directions {
		newPos := util.Position{R: pos.R + dir.R, C: pos.C + dir.C}
		if newPos.R >= 0 && newPos.C >= 0 && newPos.R < len(topographicMap) && newPos.C < len(topographicMap[0]) {
			if topographicMap[newPos.R][newPos.C] == topographicMap[pos.R][pos.C]+1 {
				score += dfs(topographicMap, newPos, visited)
			}
		}
	}
	return score
}
