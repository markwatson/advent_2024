package day9

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/markwatson/advent_2024/pkg/util"
)

// Kind of a hacky / verbose solution. The linked list solution
// isn't optimal, and more of an excuse to try out the stdlib
// linked list implementation. Seems to work okay?
type Day9 struct {}

const EMPTY = -1

func convertToSparse(data string) []int {
	var output = make([]int, 0)
	fileId := 0
	inFile := true
	for _, c := range data {
		blocks := util.CheckFatal(strconv.Atoi(string(c)))

		if inFile {
			for i := 0; i < blocks; i++ {
				output = append(output, fileId)
			}
			fileId++
			inFile = false
		} else {
			for i := 0; i < blocks; i++ {
				output = append(output, EMPTY)
			}
			inFile = true
		}
	}
	return output
}

type file struct {
	id      int
	blocks  int
}

func convertToIndexed(data string) (*list.List, int) {
	l := list.New()
	fileId := 0
	inFile := true
	for _, c := range data {
		blocks := util.CheckFatal(strconv.Atoi(string(c)))

		if inFile {
			l.PushBack(file{id: fileId, blocks: blocks})
			fileId++
			inFile = false
		} else {
			l.PushBack(file{id: EMPTY, blocks: blocks})
			inFile = true
		}
	}
	return l, fileId - 1
}

func compress(sparse []int) []int {
	writePtr := 0
	readPtr := len(sparse) - 1

	for readPtr > writePtr {
		if sparse[writePtr] != EMPTY {
			writePtr++
		} else if sparse[readPtr] == EMPTY {
			readPtr--
		} else {
			sparse[writePtr] = sparse[readPtr]
			sparse[readPtr] = EMPTY
			readPtr--
			writePtr++
		}
	}

	return sparse
}

func defragment(l *list.List, maxId int) *list.List {
	for id := maxId; id >= 0; id-- {
		id1 := 0
		id2 := l.Len() - 1
		pt1 := l.Front()
		pt2 := l.Back() // Slightly slower, but maybe needed. TODO move out of loop.
		// Devance pt2 to the file
		for pt2.Value.(file).id != id {
			pt2 = pt2.Prev()
			id2--
		}

		for pt1 != nil && id1 < id2 {
			// Advance pt1 to first free block
			for pt1.Value.(file).id != EMPTY {
				pt1 = pt1.Next()
				id1++
			}

			// Copy file
			source := pt2.Value.(file)
			dest := pt1.Value.(file)
			if source.blocks <= dest.blocks {
				bookmark := pt1.Prev()
				f1 := file{id: source.id, blocks: source.blocks}
				f2 := file{id: EMPTY, blocks: dest.blocks - source.blocks}
				// Remove file in destination and insert
				l.Remove(pt1)
				bookmark = l.InsertAfter(f1, bookmark)
				l.InsertAfter(f2, bookmark)
				// Replace source with empty
				pt2.Value = file{id: EMPTY, blocks: source.blocks}
				break
			} else {
				// advance to next free block
				pt1 = pt1.Next()
				id1++
			}
		}
	}
	return l
}

// Used for debugging
func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		f := e.Value.(file)
		for i := 0; i < f.blocks; i++ {
			if f.id == EMPTY {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", f.id)
			}
		}
	}
	fmt.Println()
}

func listToArray(l *list.List) []int {
	output := make([]int, 0)
	for e := l.Front(); e != nil; e = e.Next() {
		value := e.Value.(file)
		for i := 0; i < value.blocks; i++ {
			output = append(output, value.id)
		}
	}
	return output
}

func calcChecksum(data []int) int64 {
	var check int64 = 0
	for i, v := range data {
		if v != EMPTY {
			check += int64(i * v)
		}
	}
	return check
}

func (d Day9) Run(inputFile string) {
	data := util.CheckFatal(util.ReadString(inputFile))
	data = strings.TrimSpace(data)

	// Part 1
	sparse := convertToSparse(data)
	compressed := compress(sparse)
	checksum := calcChecksum(compressed)
	fmt.Printf("Part 1: %d\n", checksum)

	// Part 2
	indexed, maxIdx := convertToIndexed(data)
	defrag := defragment(indexed, maxIdx)
	checksum = calcChecksum(listToArray(defrag))
	fmt.Printf("Part 2: %d\n", checksum)

}
