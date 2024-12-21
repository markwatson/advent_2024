package day10

import (
	"fmt"

	"github.com/markwatson/advent_2024/pkg/util"
)

type Day10 struct {

}

func (d Day10) Run(inputFile string) {
	data := util.CheckFatal(util.ReadString(inputFile))
	fmt.Println(data)
}
