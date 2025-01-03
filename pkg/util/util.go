package util

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

// Default helpers for slices
func GetOrDefault[T any](m []T, key int, def T) T {
	if key < 0 || key >= len(m) {
		return def
	}
	return m[key]
}

func GetOrDefault2d[T any](m [][]T, r int, c int, def T) T {
	if r < 0 || c < 0 || r >= len(m) || c >= len(m[r]) {
		return def
	}
	return m[r][c]
}

// Row, column position
type Position struct {
	R int
	C int
}

// Simple byte matrix
type Matrix struct {
	Data [][]byte
	Width int
	Height int
}

func NewMatrix(data []string) *Matrix {
	m := new(Matrix)
	m.Data = make([][]byte, len(data))
	m.Width = len(data[0])
	m.Height = len(data)
	for i, row := range data {
		if len(row) != m.Width {
			panic("Rows must be the same length")
		}
		m.Data[i] = []byte(row)
	}
	return m
}

// Gets a value, defaulting to 0 (nul) if out of bounds
func (m *Matrix) Get(r int, c int) byte {
	if r < 0 || c < 0 || r >= len(m.Data) || c >= len(m.Data[r]) {
		return 0
	}
	return m.Data[r][c]
}



// func (m *Matrix) GetUp(r int, c int) byte {
// 	return m.Get(r - 1, c)
// }

// func UpIdx(r int, c int) (int, int) {
// 	return r - 1, c
// }

// func LeftIdx(r int, c int) (int, int) {
// 	return r, c - 1
// }

// func RightIdx(r int, c int) (int, int) {
// 	return r, c + 1
// }

// func DownIdx(r int, c int) (int, int) {
// 	return r + 1, c
// }



func (m *Matrix) String() string {
	var sb strings.Builder
	for _, row := range m.Data {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	return sb.String()
}

// If an error is not nil, log and quit.
func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CheckFatal[T any](v T, e error) T {
	if e != nil {
		log.Fatal(e)
	}
	return v
}

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	output := make([]string, 0)
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			output = append(output, line)
		}
	}

	return output, nil
}

func ReadGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	output := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			output = append(output, []rune(line))
		}
	}

	return output, nil
}

func ReadString(filename string) (string, error) {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
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
func ReadNumbers[V int64 | float64](filename string, v V) ([][]V, error) {
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
