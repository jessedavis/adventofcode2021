package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var stack []string
	// 0 = parens
	// 1 = square brackets
	// 2 = curly brackets
	// 3 = carets
	illegalCharCounts := []int{0, 0, 0, 0}

	for scanner.Scan() {
		line := scanner.Text()
	SplitUpLine:
		for _, c := range strings.Split(line, "") {
			switch c {
			case "(", "[", "{", "<":
				stack = append(stack, c)
			case ")":
				end := len(stack) - 1
				top := stack[end]
				stack = stack[:end]
				if top != "(" {
					illegalCharCounts[0]++
					break SplitUpLine
				}
			case "]":
				end := len(stack) - 1
				top := stack[end]
				stack = stack[:end]
				if top != "[" {
					illegalCharCounts[1]++
					break SplitUpLine
				}
			case "}":
				end := len(stack) - 1
				top := stack[end]
				stack = stack[:end]
				if top != "{" {
					illegalCharCounts[2]++
					break SplitUpLine
				}
			case ">":
				end := len(stack) - 1
				top := stack[end]
				stack = stack[:end]
				if top != "<" {
					illegalCharCounts[3]++
					break SplitUpLine
				}
			}
		}
	}

	syntaxScore := illegalCharCounts[0]*3 + illegalCharCounts[1]*57 + illegalCharCounts[2]*1197 + illegalCharCounts[3]*25137
	fmt.Printf("syntax score = %d", syntaxScore)
}
