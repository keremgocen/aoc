package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("part2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalScore := 0

	for fileScanner.Scan() {
		if isOverlappingAtAll(fileScanner.Text()) {
			totalScore += 1
		}
		fmt.Println("total score", totalScore)
	}

	fmt.Println("final total score", totalScore)

	readFile.Close()
}

func isOverlappingFully(input string) bool {
	allSections := strings.Split(input, ",")

	firstSection := strings.Split(allSections[0], "-")
	secondSection := strings.Split(allSections[1], "-")

	s11, _ := strconv.Atoi(firstSection[0])
	s12, _ := strconv.Atoi(firstSection[1])
	s21, _ := strconv.Atoi(secondSection[0])
	s22, _ := strconv.Atoi(secondSection[1])

	// pick min
	if s11 > s21 {
		fmt.Printf("s21 is smaller %d-%d, %d-%d\n", s11, s12, s21, s22)
		s11, s12, s21, s22 = swap(s11, s12, s21, s22)
	} else {
		if s11 == s21 && s12 < s22 {
			fmt.Printf("s12 is smaller %d-%d, %d-%d\n", s11, s12, s21, s22)
			s11, s12, s21, s22 = swap(s11, s12, s21, s22)
		}
	}
	// s11-s12 is the larger gap

	if s11 <= s21 && s12 >= s22 {
		fmt.Printf("overlap %d-%d, %d-%d\n", s11, s12, s21, s22)

		return true
	}

	fmt.Printf("no overlap %d-%d, %d-%d\n", s11, s12, s21, s22)

	return false
}

func swap(a, b, c, d int) (int, int, int, int) {
	t1, t2 := a, b
	a, b = c, d
	c, d = t1, t2
	return a, b, c, d
}

func isOverlappingAtAll(input string) bool {
	allSections := strings.Split(input, ",")

	firstSection := strings.Split(allSections[0], "-")
	secondSection := strings.Split(allSections[1], "-")

	f1, _ := strconv.Atoi(firstSection[0])
	f2, _ := strconv.Atoi(firstSection[1])
	s1, _ := strconv.Atoi(secondSection[0])
	s2, _ := strconv.Atoi(secondSection[1])

	// f1-f2 s1-s2

	// 2-4 3-8
	if f1 >= s1 && f1 <= s2 {
		return true
	}

	if f2 >= s1 && f2 <= s2 {
		return true
	}

	if s1 >= f1 && s1 <= f2 {
		return true
	}

	if s2 >= f1 && s2 <= f2 {
		return true
	}

	fmt.Printf("no overlap %d-%d, %d-%d\n", f1, f2, s1, s2)

	return false
}
