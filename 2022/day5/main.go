package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(v T) {
	s.vals = append([]T{v}, s.vals...)
}

func (s *Stack[T]) PushBack(v T) {
	s.vals = append(s.vals, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	top := s.vals[0]
	s.vals = s.vals[1:]
	return top, true
}

func (s *Stack[T]) CutN(n int) (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	nth := s.vals[n]
	s.vals = append(s.vals[:n], s.vals[n+1:]...)
	return nth, true
}

// 	   [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2

func main() {
	readFile, err := os.Open("part2.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var stacks []Stack[string]

	for fileScanner.Scan() {
		input := fileScanner.Text()

		if stacks == nil {
			stacks = make([]Stack[string], (len(input)+1)/4)
			fmt.Println("len stacks", len(stacks))
		}

		re := regexp.MustCompile(`\[[A-Z]\]`)

		if re.MatchString(input) {
			inputs := re.FindAllStringIndex(input, -1)
			for k, v := range inputs {
				val := input[v[0]:v[1]]
				fmt.Println("k, v", k, v, val, v[0]/4)
				stackIdx := v[0] / 4
				stacks[stackIdx].PushBack(val)
			}
			fmt.Println("stacks", stacks)
		} else {
			inputs := strings.Split(input, " ")
			if inputs[0] == "move" {
				move, _ := strconv.Atoi(inputs[1])
				from, _ := strconv.Atoi(inputs[3])
				from -= 1
				to, _ := strconv.Atoi(inputs[5])
				to -= 1
				fmt.Println("move", move, "-from-", from, "-to-", to)
				for i := 0; i < move; i++ {
					var val string
					if i < move-1 {
						fmt.Println("popping N", i)
						val, _ = stacks[from].CutN(move - 1 - i)
					} else {
						fmt.Println("just pop", i)
						val, _ = stacks[from].Pop()
					}
					fmt.Println("moving", val, "from", from, "to", to)
					stacks[to].Push(val)
				}
				fmt.Println("stacks", stacks)
			} else {
				fmt.Println("skip", input)
			}
		}
	}

	fmt.Println("tops of stacks")

	for _, s := range stacks {
		v, _ := s.Pop()
		fmt.Printf("%v", string(v[1]))
	}

	readFile.Close()
}
