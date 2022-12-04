package main

import (
	"fmt"

	"github.com/m0sh1x2/advent_of_code_2022/day01"
	"github.com/m0sh1x2/advent_of_code_2022/day02"
)

func main() {
	inpputFileArray := day01.ReadFileByLines("./day02/input.txt")
	fmt.Println(day02.Day02Part01(inpputFileArray))
	fmt.Println(day02.Day02Part02(inpputFileArray))
}
