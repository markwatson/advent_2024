package main

import (
	"log"
	"os"

	"github.com/markwatson/advent_2024/pkg/days"
)


func main() {
	// check arguments
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <day> <input|example>", os.Args[0])
	}
	dayName := os.Args[1]
	input := os.Args[2]

	// construct the path
	fname := "./data/day" + dayName + "/" + input + ".txt"
	
	// Setup
	days.RegisterAllDays()

	// run the day
	days.Run(dayName, fname)
}
