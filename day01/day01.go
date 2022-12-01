package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	currentElfCalories := 0
	highestElfCalories := 0

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		if currentLine == "" {
			fmt.Println(currentElfCalories)
			fmt.Println("------")
			currentElfCalories = 0
		} else {
			currentLine, err := strconv.Atoi(currentLine)
			check(err) // possible input text/string

			currentElfCalories += int(currentLine)
			// fmt.Println(currentElfCalories)

			if currentElfCalories > highestElfCalories {
				highestElfCalories = currentElfCalories
			}
		}
	}

	fmt.Println("Highest Elf with Calories ", highestElfCalories)
}
