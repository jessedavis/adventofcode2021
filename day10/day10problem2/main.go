package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	//illegalCharCounts := []int{0, 0, 0, 0}
	var completionScores []int

Line:
	for scanner.Scan() {
		line := scanner.Text()
		stack = nil
		for _, c := range strings.Split(line, "") {
			switch c {
			case "(", "[", "{", "<":
				stack = append(stack, c)
			case ")":
				end := len(stack) - 1
				top := stack[end]
				stack = stack[:end]
				if top != "(" {
					// corrupted line, throw it away
					continue Line
				}
			case "]":
				end := len(stack) - 1
				top := stack[end]
				stack = stack[:end]
				if top != "[" {
					// corrupted line, throw it away
					continue Line
				}
			case "}":
				end := len(stack) - 1
				top := stack[end]
				stack = stack[:end]
				if top != "{" {
					// corrupted line, throw it away
					continue Line
				}
			case ">":
				end := len(stack) - 1
				top := stack[end]
				stack = stack[:end]
				if top != "<" {
					// corrupted line, throw it away
					continue Line
				}
			}
		}

		fmt.Printf("%+q\n", stack)
		// remaining stack, pop and create completion string
		var completionString []string
		for i := len(stack) - 1; i >= 0; i-- {
			switch stack[i] {
			case "(":
				completionString = append(completionString, ")")
			case "[":
				completionString = append(completionString, "]")
			case "{":
				completionString = append(completionString, "}")
			case "<":
				completionString = append(completionString, ">")
			}
		}

		// determine score of completion string
		var completionScore int = 0
		var charScore int = 0
		for _, c := range completionString {
			switch c {
			case ")":
				charScore = 1
			case "]":
				charScore = 2
			case "}":
				charScore = 3
			case ">":
				charScore = 4
			}
			completionScore *= 5
			completionScore += charScore
		}
		completionScores = append(completionScores, completionScore)
	}

	//syntaxScore := illegalCharCounts[0]*3 + illegalCharCounts[1]*57 + illegalCharCounts[2]*1197 + illegalCharCounts[3]*25137
	//fmt.Printf("syntax score = %d", syntaxScore)

	sort.Slice(completionScores, func(i, j int) bool {
		return completionScores[i] < completionScores[j]
	})
	fmt.Printf("%+v\n", completionScores)
	// "cheating" here, int seems to floor towards lower, also since output is stated to always have an odd
	// set of incomplete lines
	mid := len(completionScores) / 2
	fmt.Printf("middle score for completion strings = %d\n", completionScores[mid])
}
