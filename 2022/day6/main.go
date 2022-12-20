package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	readFile, err := os.Open("part2.txt")
	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		input := fileScanner.Text()

		idx := findEndOfMarkerIdx(input)
		log.Println("idx", idx)
	}

	readFile.Close()
}

func findEndOfMarkerIdx(input string) int {
	log.Println("input", input)
	start := 0
	for end := range input {
		if end == start+14 {
			// a b a d b c
			// check if last 4 chars are different
			log.Printf("check start:%d end:%d, str:%v\n", start, end, input[start:end])
			if !hasRepeatingRune(input[start:end]) {
				return end
			}
			start++
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
			log.Println("checking", i, ":", str[i], "against", j, ":", str[j])
			if str[i] == str[j] {
				return true
			}
		}
	}
	return false
}
