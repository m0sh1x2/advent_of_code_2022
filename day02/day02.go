package day02

import (
	"fmt"
	"strings"

	"github.com/m0sh1x2/advent_of_code_2022/day01"
)

func IdentifyType(inputType string) int {
	switch inputType {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		return 0
	}
}

func IdentifyWinnerScore(inputData string) int {
	fmt.Println(inputData)
	// AX
	// BY
	// CZ
	switch inputData {
	case "A X":
		return 4
	case "B Y":
		return 5
	case "C Z":
		return 6

	case "A Y":
		return 1
	case "B Z":
		return 2
	case "C X":
		return 3

	case "A Z":
		return 7
	case "B X":
		return 8
	case "C Y":
		return 9
	}

	return 0
	// A A = 1 + 1 = 2
	// B B = 2 + 2 = 4
	// C C = 3 + 3 = 6

	// A B = 1 + 0 = 1
	// B C = 2 + 0 = 2
	// C A = 3 + 0 = 3

	// A C = 1 + 6 = 7
	// B A = 2 + 6 = 8
	// C B = 3 + 6 = 9
}

func EngageCombat(playerInputs string) int {

	currentInputs := IdentifyType(strings.Split(playerInputs, " ")[0]) + IdentifyType(strings.Split(playerInputs, " ")[1])

	switch currentInputs {

	}
	return currentInputs
	// A, X -> Rock -> 1 point
	// B, Y -> Paper -> 2 points
	// C -> Scissors -> 3 points

	// Win = 6 points
	// Draw = Element points

	// fmt.Println(currentInputs)
	// switch currentInputs {
	// case "RockScissors":
	// 	fmt.Println("Win 1+6")
	// 	return 7
	// case "RockPaper":
	// 	fmt.Println("Loose 1+0")
	// 	return 1
	// case "RockRock":
	// 	fmt.Println("Draw 1+3")
	// 	return 4

	// case "PaperRock":
	// 	fmt.Println("Win 2+6")
	// 	return 8
	// case "PaperScissors":
	// 	fmt.Println("Loose 2+0")
	// 	return 2
	// case "PaperPaper":
	// 	fmt.Println("Draw 2+3")
	// 	return 5

	// case "ScissorsPaper":
	// 	fmt.Println("Win 3+6")
	// 	return 9
	// case "ScissorsRock":
	// 	fmt.Println("Loose 3+0")
	// 	return 3
	// case "ScissorsScissors":
	// 	fmt.Println("Draw 3+3")
	// 	return 6
	// default:
	// 	fmt.Println("Unknow Input")
	// 	return 0
	// }
}

func Day02(inputFilePath string) int {
	inputFile := day01.ReadFileByLines(inputFilePath)
	score := 0

	// Map every score based on element
	points := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	// get the score based on element

	win := map[int]int{
		1: 3, // - Rock wins over Scissor
		2: 1, // - Paper wins over Rock
		3: 2, // - Scissors wins over Paper
	}

	// fmt.Println(win[1])

	// fmt.Println(points["B"])

	for inputFile.Scan() {
		elf, human := points[strings.Split(inputFile.Text(), " ")[0]], points[strings.Split(inputFile.Text(), " ")[1]]

		fmt.Println(elf)
		fmt.Println(human)
		fmt.Println(win[human])

		fmt.Println(" _---- ")

		// Wa always get the minimal score of the element that we choose
		score += human

		if win[human] == elf {
			// If we win we get max score + 6
			score += 6
		} else if elf == human {
			// Draw is element + 3
			score += 3
		}

	}

	fmt.Println(score)

	return score
}
