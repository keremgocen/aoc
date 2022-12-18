package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("part2.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		input := fileScanner.Text()

		idx := findEndOfMarkerIdx(input)
		fmt.Println("idx", idx)
	}

	readFile.Close()
}

func findEndOfMarkerIdx(input string) int {
	fmt.Println("input", input)
	start := 0
	for end := range input {
		if end == start+14 {
			// a b a d b c
			// check if last 4 chars are different
			fmt.Printf("check start:%d end:%d, str:%v\n", start, end, input[start:end])
			if !hasRepeatingRune(input[start:end]) {
				return end
			}
			start += 1
		}
	}
	return 0
}

func hasRepeatingRune(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		for j := 1; j < len(str); j++ {
			if i == j {
				continue
			}
			fmt.Println("checking", i, ":", str[i], "against", j, ":", str[j])
			if str[i] == str[j] {
				return true
			}
		}
	}
	return false
}
