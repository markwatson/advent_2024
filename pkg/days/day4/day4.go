package day4

import (
	"fmt"

	"github.com/markwatson/advent_2024/pkg/util"
)

type Day4 struct {}

// XMAS
func part1(data []string) int {
	total := 0

	m := util.NewMatrix(data)
	search := "XMAS"
	
	// There's gotta be a better way - so many off by one errors :(
	for r := 0; r < m.Height; r++ {
		for c := 0; c < m.Width; c++ {
			// right
			if string([]byte{m.Get(r, c), m.Get(r, c+1), m.Get(r, c+2), m.Get(r, c+3)}) == search {
				total += 1
			}
			// down
			if string([]byte{m.Get(r, c), m.Get(r+1, c), m.Get(r+2, c), m.Get(r+3, c)}) == search {
				total += 1
			}
			// left
			if string([]byte{m.Get(r, c), m.Get(r, c-1), m.Get(r, c-2), m.Get(r, c-3)}) == search {
				total += 1
			}
			// up
			if string([]byte{m.Get(r, c), m.Get(r-1, c), m.Get(r-2, c), m.Get(r-3, c)}) == search {
				total += 1
			}
			// down right diagonal
			if string([]byte{m.Get(r, c), m.Get(r+1, c+1), m.Get(r+2, c+2), m.Get(r+3, c+3)}) == search {
				total += 1
			}
			// down left diagonal
			if string([]byte{m.Get(r, c), m.Get(r+1, c-1), m.Get(r+2, c-2), m.Get(r+3, c-3)}) == search {
				total += 1
			}
			// up right diagonal
			if string([]byte{m.Get(r, c), m.Get(r-1, c+1), m.Get(r-2, c+2), m.Get(r-3, c+3)}) == search {
				total += 1
			}
			// up left diagonal
			if string([]byte{m.Get(r, c), m.Get(r-1, c-1), m.Get(r-2, c-2), m.Get(r-3, c-3)}) == search {
				total += 1
			}
			
		}
	}

	return total
}

// X-MAS
func part2(data []string) int {
	total := 0

	m := util.NewMatrix(data)
	for r := 0; r < m.Height; r++ {
		for c := 0; c < m.Width; c++ {
			if m.Data[r][c] == 'A' {
				ul := m.Get(r-1, c-1)
				ur := m.Get(r-1, c+1)
				dl := m.Get(r+1, c-1)
				dr := m.Get(r+1, c+1)

				first := (ul == 'M' && dr == 'S') || (ul == 'S' && dr == 'M')
				second := (ur == 'M' && dl == 'S') || (ur == 'S' && dl == 'M')

				if first && second {
					total += 1
				}
			}
		}
	}

	return total
}

func (d Day4) Run(inputFile string) {
	data := util.CheckFatal(util.ReadLines(inputFile))

	total := part1(data)
	fmt.Printf("Part 1: %d\n", total)

	total2 := part2(data)
	fmt.Printf("Part 2: %d\n", total2)

}
