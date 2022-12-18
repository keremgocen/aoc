package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	readFile, err := os.Open("part2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalScore := 0

	scannedLines := make([]string, 0)
	scanCount := 0
	for fileScanner.Scan() {
		scannedLines = append(scannedLines, fileScanner.Text())
		scanCount += 1
		if scanCount%3 == 0 {
			totalScore += calculatePriorityPart2(scannedLines)
			fmt.Println("total score", totalScore)

			scanCount = 0
			scannedLines = make([]string, 0)
		}
	}

	fmt.Println("final total score", totalScore)

	readFile.Close()
}

func calculatePriorityPart1(input string) int {
	score := 0

	halfLen := len(input) / 2

	fmt.Println("str", input)
	fmt.Println("part1 str", input[:halfLen])
	fmt.Println("part2 str", input[halfLen:])

	part1 := make(map[rune]bool, halfLen)
	for i, r := range input {
		if i < halfLen {
			part1[r] = true
		} else {
			if _, ok := part1[r]; ok {
				score += getPriority(r)
				fmt.Printf("%c - %d\n", r, getPriority(r))
				break
			}
		}
	}

	fmt.Println("part1", part1)

	return score
}

func calculatePriorityPart2(input []string) int {
	commonElements := make(map[rune]bool, 0)

	inputMap1 := make(map[rune]bool, len(input[0]))
	for _, r := range input[0] {
		inputMap1[r] = true
	}

	for _, r := range input[1] {
		if _, ok := inputMap1[r]; ok {
			commonElements[r] = true
		}
	}

	for _, r := range input[2] {
		if _, ok := commonElements[r]; ok {
			fmt.Printf("found %c in all 3 inputs - %d\n", r, getPriority(r))
			return getPriority(r)
		}
	}

	return 0
}

func getPriority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 38
	}
	return int(r) - 96
}
