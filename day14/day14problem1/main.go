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

type InsertionRule struct {
	pattern string
	insert  string
}

// https://stackoverflow.com/questions/53942222/how-to-match-all-overlapping-pattern
// since regexp funcs are non-overlapping
func findAllSubstringsIndex(s string, sub string) []int {
	idx := []int{}
	j := 0

	for {
		i := strings.Index(s[j:], sub)
		if i == -1 {
			break
		}
		idx = append(idx, j+i)
		j += i + 1
	}

	return idx
}

func countElements(s string) map[string]int {
	counts := make(map[string]int)

	for _, c := range s {
		counts[string(c)]++
	}

	return counts
}

func main() {
	//f, err := os.Open("../input.txt")
	f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var polymerTemplate string
	insertionRules := make([]InsertionRule, 0)
	steps := 10

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
			insertionRules = append(insertionRules, InsertionRule{fields[0], fields[1]})
		}
	}

	for step := 1; step <= steps; step++ {
		to_insert := make(map[int]string, len(insertionRules))
		keys := make([]int, len(insertionRules))
		for _, rule := range insertionRules {
			indexes := findAllSubstringsIndex(polymerTemplate, rule.pattern)
			for _, index := range indexes {
				to_insert[index] = rule.insert
				keys = append(keys, index)
			}
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})

		nextTemplate := make([]rune, len(polymerTemplate)+len(keys)+1)
		for i, s := range polymerTemplate {
			insert, ok := to_insert[i]
			if ok {
				// cheating, should only be one char/rune here
				nextTemplate = append(nextTemplate, s, []rune(insert)[0])
			} else {
				nextTemplate = append(nextTemplate, s)
			}
		}

		polymerTemplate = string(nextTemplate)
		// HACK: i probably need to do better about creating the string I'm sending here
		polymerTemplate = strings.Replace(polymerTemplate, "\x00", "", -1)
		//fmt.Printf("After step %d, polymerTemplate = %s\n", step, polymerTemplate)
	}

	counts := countElements(polymerTemplate)

	var countValues []int
	for k := range counts {
		countValues = append(countValues, counts[k])
	}
	sort.Slice(countValues, func(i, j int) bool {
		return countValues[i] < countValues[j]
	})

	leastCommon := countValues[0]
	mostCommon := countValues[len(countValues)-1]

	fmt.Printf("%v\n", counts)
	fmt.Printf("After %d steps, most common - least common = %d\n", steps, mostCommon-leastCommon)
}
