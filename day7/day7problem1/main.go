package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var crabPositions []int

	for scanner.Scan() {
		for _, v := range strings.Split(scanner.Text(), ",") {
			pos, _ := strconv.Atoi(v)
			crabPositions = append(crabPositions, pos)
		}
	}

	// calculate median
	// sort array
	sort.Ints(crabPositions)

	var median int
	n := len(crabPositions)
	if len(crabPositions)%2 == 0 {
		median = (crabPositions[n/2] + crabPositions[(n/2)-1]) / 2
	} else {
		median = crabPositions[(n-1)/2]
	}

	// now calculate the sum of the distance for each crab to the median
	var fuelUsed int
	for i := range crabPositions {
		pos := crabPositions[i]
		if pos < median {
			fuelUsed += median - pos
		} else {
			fuelUsed += pos - median
		}
	}

	fmt.Printf("median = %d\n", median)
	fmt.Printf("cheapest fuel used = %d\n", fuelUsed)
}
