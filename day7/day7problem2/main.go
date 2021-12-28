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

/*
As it turns out, crab submarine engines don't burn fuel at a constant rate.
Instead, each change of 1 step in horizontal position costs 1 more unit of
fuel than the last: the first step costs 1, the second step costs 2, the
third step costs 3, and so on.
*/
func CostToMove(from int, to int) int {
	d := from - to
	if from < to {
		d = to - from
	}

	total := 0
	cost := 1

	for i := 1; i <= d; i++ {
		total += cost
		cost++
	}

	return total
}

// x must be sorted
func Median(x []int) int {
	var median int
	n := len(x)

	if len(x)%2 == 0 {
		median = (x[n/2] + x[(n/2)-1]) / 2
	} else {
		median = x[(n-1)/2]
	}

	return median
}

//func CostToMoveAllTo(x []int, to int, limit int) int {
func CostToMoveAllTo(x []int, to int) int {
	sum := 0
	for _, v := range x {
		sum += CostToMove(v, to)
		// break early if we are already over the limit, 0 max cost means
		// we're running for first time
		//if sum >= limit && limit > 0 {
		//return limit
		//}
	}
	return sum
}

func main() {
	//f, err := os.Open("../input.txt")
	f, err := os.Open("../part1exampledata.txt")
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

	// sort array
	sort.Ints(crabPositions)

	minimumFuelUsed := CostToMoveAllTo(crabPositions, 0)

	// not great here, need to find better way to do this than n^2 (at least)
	for x := 1; x <= crabPositions[len(crabPositions)-1]; x++ {
		c := CostToMoveAllTo(crabPositions, x)
		if c < minimumFuelUsed {
			minimumFuelUsed = c
		}
	}

	fmt.Printf("minimum fuel used = %d\n", minimumFuelUsed)
}
