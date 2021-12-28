package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func pruneRatings(p map[string]bool, sampleLength int, bitToConsider int, compareMostCommon bool) map[string]bool {
	// if input is only one sample, we've already determined the "correct" sample
	if len(p) == 1 {
		return p
	}

	var bitCounts = make([][]int, sampleLength)
	for i := range bitCounts {
		bitCounts[i] = []int{0, 0}
	}

	// determine most common value
	for line, _ := range p {
		for i, c := range strings.Split(line, "") {
			switch c {
			case "1":
				bitCounts[i][1]++
			case "0":
				bitCounts[i][0]++
			}
		}
	}

	var bitValueToKeep = ""

	if bitCounts[bitToConsider][0] > bitCounts[bitToConsider][1] {
		if compareMostCommon {
			bitValueToKeep = "0"
		} else {
			bitValueToKeep = "1"
		}
	} else if bitCounts[bitToConsider][0] < bitCounts[bitToConsider][1] {
		if compareMostCommon {
			bitValueToKeep = "1"
		} else {
			bitValueToKeep = "0"
		}
	} else {
		if compareMostCommon {
			bitValueToKeep = "1"
		} else {
			bitValueToKeep = "0"
		}
	}

	for line, _ := range p {
		if string(line[bitToConsider]) != bitValueToKeep {
			// prune the value
			delete(p, line)
		}
	}

	return p
}

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// make a set
	input := make(map[string]bool)

	// assumes strings are regular, but probably ok
	var sampleLength int = 0

	for scanner.Scan() {
		line := scanner.Text()
		sampleLength = len(line)
		if !input[line] {
			input[line] = true
		}
	}

	oxygenSamples := make(map[string]bool)
	scrubberSamples := make(map[string]bool)

	for k, v := range input {
		oxygenSamples[k] = v
		scrubberSamples[k] = v
	}

	for i := 0; i < sampleLength; i++ {
		oxygenSamples = pruneRatings(oxygenSamples, sampleLength, i, true)
		scrubberSamples = pruneRatings(scrubberSamples, sampleLength, i, false)
	}

	var finalOxygenSample = ""
	var finalScrubberSample = ""
	for key, _ := range oxygenSamples {
		finalOxygenSample = key
	}
	for key, _ := range scrubberSamples {
		finalScrubberSample = key
	}
	fmt.Printf("oxygen sample = %s\n", finalOxygenSample)
	fmt.Printf("scurbber sample = %s\n", finalScrubberSample)

	var convertedOxygen, _ = strconv.ParseInt(finalOxygenSample, 2, 0)
	var convertedScrubber, _ = strconv.ParseInt(finalScrubberSample, 2, 0)
	fmt.Printf("oxygen = %s, scurbber = %s\n", finalOxygenSample, finalScrubberSample)
	fmt.Printf("converted oxygen = %d, scrubber = %d, product = %d",
		convertedOxygen, convertedScrubber, convertedOxygen*convertedScrubber)
}
