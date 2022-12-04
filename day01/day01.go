package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileByLines(inputFilePath string) []string {
	file, err := os.Open(inputFilePath)
	check(err)

	oututArray := []string{}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		oututArray = append(oututArray, fileScanner.Text())
	}
	return oututArray
}

func GetHighetstElfCalories(inputFilePath string) {

	fileScanner := ReadFileByLines(inputFilePath)

	currentElfCalories := 0
	allElfs := []int{}

	var highestElfCalories [3]int

	for _, line := range fileScanner {
		currentLine := line

		if currentLine == "" {
			// fmt.Println(currentElfCalories)
			allElfs = append(allElfs, currentElfCalories)
			currentElfCalories = 0
		} else {
			currentLine, err := strconv.Atoi(currentLine)
			check(err) // possible input text/string

			currentElfCalories += int(currentLine)

			if currentElfCalories > highestElfCalories[0] {
				highestElfCalories[0] = currentElfCalories
			}
		}
	}

	// fmt.Println("Highest Elf with Calories ", highestElfCalories)

	// Sort the elfs and get the top 3 elfs by calories
	sort.Ints(allElfs)
	topElfCalories := 0
	for i := 1; i <= 3; i++ {
		elfLength := len(allElfs)
		topElfCalories += allElfs[elfLength-i]
	}
	fmt.Println(topElfCalories)

}

func main() {
	GetHighetstElfCalories("input.txt")
}
