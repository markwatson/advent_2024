package days

import (
	"log"

	"github.com/markwatson/advent_2024/pkg/days/day1"
	"github.com/markwatson/advent_2024/pkg/days/day2"
)

type RunnableDay interface {
	Run(inputFile string)
}

type Day struct {
	Name string
	Runner RunnableDay
}

var Days = make(map[string]Day)

func Register(name string, day RunnableDay) {
	Days[name] = Day{Name: name, Runner: day}
}

func Run(day string, inputFile string) {
	d, ok := Days[day]
	if !ok {
		log.Fatalf("Day %s not found", day)
	}
	d.Runner.Run(inputFile)
}

func RegisterAllDays() {
	Register("1", &day1.Day1{})
	Register("2", &day2.Day2{})
}