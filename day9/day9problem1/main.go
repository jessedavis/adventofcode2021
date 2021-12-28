package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	locations := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]int, len(line))
		for i, c := range line {
			location, _ := strconv.Atoi(string(c))
			digits[i] = location
		}
		locations = append(locations, digits)
	}

	lowpoints := make([]int, 0)

	num_rows := len(locations)
	num_cols := len(locations[0])

	// top left
	currentPoint := locations[0][0]
	if currentPoint < locations[0][1] && currentPoint < locations[1][0] {
		lowpoints = append(lowpoints, currentPoint)
	}
	// top right
	currentPoint = locations[0][num_cols-1]
	if currentPoint < locations[0][num_cols-2] && currentPoint < locations[1][num_cols-1] {
		lowpoints = append(lowpoints, currentPoint)
	}
	// bottom left
	currentPoint = locations[num_rows-1][0]
	if currentPoint < locations[num_rows-2][0] && currentPoint < locations[num_rows-1][1] {
		lowpoints = append(lowpoints, currentPoint)
	}
	// bottom right
	currentPoint = locations[num_rows-1][num_cols-1]
	if currentPoint < locations[num_rows-1][num_cols-2] && currentPoint < locations[num_rows-2][num_cols-1] {
		lowpoints = append(lowpoints, currentPoint)
	}
	// top edge
	for j := 1; j < num_cols-1; j++ {
		currentPoint = locations[0][j]
		if currentPoint < locations[0][j-1] && currentPoint < locations[1][j] && currentPoint < locations[0][j+1] {
			lowpoints = append(lowpoints, currentPoint)
		}
	}
	// bottom edge
	for j := 1; j < num_cols-1; j++ {
		i := num_rows - 1
		currentPoint = locations[i][j]
		if currentPoint < locations[i][j-1] && currentPoint < locations[i-1][j] && currentPoint < locations[i][j+1] {
			lowpoints = append(lowpoints, currentPoint)
		}
	}
	// left edge
	for i := 1; i < num_rows-1; i++ {
		currentPoint = locations[i][0]
		if currentPoint < locations[i-1][0] && currentPoint < locations[i][1] && currentPoint < locations[i+1][0] {
			lowpoints = append(lowpoints, currentPoint)
		}
	}
	// right edge
	for i := 1; i < num_rows-1; i++ {
		j := num_cols - 1
		currentPoint = locations[i][j]
		if currentPoint < locations[i-1][j] && currentPoint < locations[i][j-1] && currentPoint < locations[i+1][j] {
			lowpoints = append(lowpoints, currentPoint)
		}
	}

	for i := 1; i < num_rows-1; i++ {
		for j := 1; j < num_cols-1; j++ {
			currentPoint = locations[i][j]
			// all other spots
			if currentPoint < locations[i-1][j] &&
				currentPoint < locations[i][j+1] &&
				currentPoint < locations[i+1][j] &&
				currentPoint < locations[i][j-1] {
				lowpoints = append(lowpoints, currentPoint)
			}
		}
	}

	sum := 0
	for _, v := range lowpoints {
		sum += v + 1
	}
	//fmt.Printf("lowpoints = %q\n", lowpoints)
	fmt.Printf("sum of lowpoints = %d\n", sum)
}
