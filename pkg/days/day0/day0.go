package day0

import (
	"fmt"

	"github.com/markwatson/advent_2024/pkg/util"
)

type Day0 struct {

}

func (d Day0) Run(inputFile string) {
	data := util.CheckFatal(util.ReadString(inputFile))
	fmt.Println(data)
}
