package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var linesCounted int = 0
	var bitCounts [][]int

	for scanner.Scan() {
		line := scanner.Text()
		linesCounted++
		// once we've read first line, init the array
		// https://go.dev/doc/effective_go#slices
		if linesCounted == 1 {
			bitCounts = make([][]int, len(line))
			for i := range bitCounts {
				bitCounts[i] = []int{0, 0}
			}
		}
		for i, c := range strings.Split(line, "") {
			switch c {
			case "1":
				bitCounts[i][1]++
			case "0":
				bitCounts[i][0]++
			}
		}
	}

	var gamma = ""
	var epsilon = ""

	for i := range bitCounts {
		if bitCounts[i][0] > bitCounts[i][1] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	var convertedGamma, _ = strconv.ParseInt(gamma, 2, 0)
	var convertedEpsilon, _ = strconv.ParseInt(epsilon, 2, 0)
	fmt.Printf("gamma = %s, epsilon = %s\n", gamma, epsilon)
	fmt.Printf("converted gamma = %d, epsilon = %d, product = %d",
		convertedGamma, convertedEpsilon, convertedGamma*convertedEpsilon)
}
