package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalScore := 0

	for fileScanner.Scan() {
		totalScore += calculateScore2(fileScanner.Text())
	}

	fmt.Println("total score", totalScore)

	readFile.Close()
}

// func calculateScore(input string) int {
// 	score := 0
// 	entries := strings.Split(input, " ")

// 	switch entries[0] {
// 	case "A": // rock
// 		if entries[1] == "Y" { // rock x paper (2) = win (6)
// 			score += 8
// 		} else if entries[1] == "Z" { // rock x scissors (3) = lose (0)
// 			score += 3
// 		} else {
// 			score += 4
// 		}
// 	case "B": // paper
// 		if entries[1] == "X" { // paper x rock (1) = lose (0)
// 			score += 1
// 		} else if entries[1] == "Z" { // paper x scissors (3) = win (6)
// 			score += 9
// 		} else {
// 			score += 5
// 		}
// 	case "C": // scissors
// 		if entries[1] == "X" { // scissors x rock (1) = win (6)
// 			score += 7
// 		} else if entries[1] == "Y" { // scissors x paper (2) = lose (0)
// 			score += 2
// 		} else {
// 			score += 6
// 		}
// 	}

// 	return score
// }

func calculateScore2(input string) int {
	score := 0
	entries := strings.Split(input, " ")

	switch entries[0] {
	case "A": // rock
		if entries[1] == "Y" { // draw (3) + rock (1)
			score += 4
		} else if entries[1] == "Z" { // paper win
			score += 8
		} else { // X scissor lose
			score += 3
		}
	case "B": // paper
		if entries[1] == "X" { // rock lose
			score += 1
		} else if entries[1] == "Z" { // scissor win
			score += 9
		} else { // Y paper draw
			score += 5
		}
	case "C": // scissors
		if entries[1] == "X" { // paper lose
			score += 2
		} else if entries[1] == "Y" { // scissors draw
			score += 6
		} else { // rock win
			score += 7
		}
	}

	return score
}
