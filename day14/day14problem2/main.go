package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// rule turns 2 chars into a pair of 2 chars
// had to look up a hint on this, the point is that we just need to count
// the occurances, not create and parse the string every time

type InsertionRule struct {
	pattern     string
	insertPair1 string
	insertPair2 string
}

// guess since pairs turn into 2 pairs, only count first char of rule
// https://work.njae.me.uk/2021/12/16/advent-of-code-2021-day-14/ for explanation
func countCharactersInKeys(m map[string]int) map[string]int {
	counts := make(map[string]int)

	for k := range m {
		firstChar := k[0:1]
		counts[firstChar] += m[k]
	}

	return counts
}

// had to get a hint for memoization on this one
// https://www.reddit.com/r/adventofcode/comments/rfzq6f/comment/hqxs9jf/?utm_source=reddit&utm_medium=web2x&context=3
// https://work.njae.me.uk/2021/12/16/advent-of-code-2021-day-14/ for explanation

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var polymerTemplate string
	insertionRules := make(map[string]InsertionRule)
	steps := 40

	isTemplateLine, _ := regexp.Compile("^[[:upper:]]+$")
	isBlankLine, _ := regexp.Compile("^$")
	isRuleLine, _ := regexp.Compile("^[[:upper:]]+ -> [[:upper:]]+$")

	for scanner.Scan() {
		line := scanner.Text()

		if isBlankLine.MatchString(line) {
			continue
		}
		if isTemplateLine.MatchString(line) {
			polymerTemplate = line
		}
		if isRuleLine.MatchString(line) {
			fields := strings.Split(line, " -> ")
			c1 := string(fields[0][0])
			c2 := string(fields[0][1])
			insertionRules[fields[0]] = InsertionRule{fields[0], c1 + fields[1], fields[1] + c2}
		}
	}

	//fmt.Printf("rules = %+v\n", insertionRules)

	pairs := make(map[string]int)

	// initialize our pairs from the initial polymer template
	for i := 0; i < len(polymerTemplate)-1; i++ {
		pairs[polymerTemplate[i:i+2]]++
	}

	for step := 1; step <= steps; step++ {
		newPairs := make(map[string]int)
		for k, v := range pairs {
			newPairs[insertionRules[k].insertPair1] += v
			newPairs[insertionRules[k].insertPair2] += v
		}
		pairs = newPairs
	}

	keyCounts := countCharactersInKeys(pairs)
	// last character in template gets another count
	// TODO: should understand why
	keyCounts[polymerTemplate[len(polymerTemplate)-1:]]++
	fmt.Printf("key counts = %v\n", keyCounts)

	var countValues []int
	for k := range keyCounts {
		countValues = append(countValues, keyCounts[k])
	}
	sort.Slice(countValues, func(i, j int) bool {
		return countValues[i] < countValues[j]
	})

	leastCommon := countValues[0]
	mostCommon := countValues[len(countValues)-1]

	fmt.Printf("After %d steps, most common - least common = %d\n", steps, mostCommon-leastCommon)
}
