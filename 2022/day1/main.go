package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	max := make([]int, 3)
	calories := 0
	totalMax := 0

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			for k, m := range max {
				if calories > m {
					max[k] = calories
					fmt.Printf("max[%d] changed to %d from %d\n", k, max[k], m)
					break
				}
			}

			calories = 0
		} else {
			cal, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println("error reading line")
				continue
			}
			calories += cal
		}
	}

	for _, v := range max {
		totalMax += v
	}

	fmt.Println("max calories", max, totalMax)

	readFile.Close()
}
