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

type Point struct {
	x int
	y int
}

type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

func contains(points []Point, p Point) bool {
	for _, v := range points {
		if v.x == p.x && v.y == p.y {
			return true
		}
	}

	return false
}

func FoldPaper(points []Point, d Direction, foldline int) []Point {
	if d == Horizontal {
		// sort points around y-axis
		sort.Slice(points, func(i, j int) bool {
			return points[i].y < points[j].y
		})
		// split points into 2 halves, y < foldlline  and y > foldline
		// so find index of first point with y > foldline
		split_index := 0
		for i, v := range points {
			if v.y > foldline {
				split_index = i
				break
			}
		}
		points_to_transform := points[split_index:]
		// update points to be the "original" first half, i.e. the points being "folded onto"
		points = points[:split_index]
		// translate the y-coordinate of the points we're putting on "top" of the "original"
		// first half
		for _, v := range points_to_transform {
			v.y -= 2 * (v.y - foldline)
			// if transformed point already exists in "original" points, they merge
			// else add the point to the list of "original" points
			if !contains(points, v) {
				points = append(points, v)
			}
		}
	} else {
		// sort points around x-axis
		sort.Slice(points, func(i, j int) bool {
			return points[i].x < points[j].x
		})
		// split points into 2 halves, x < foldlline  and x > foldline
		// so find index of first point with x > foldline
		split_index := 0
		for i, v := range points {
			if v.x > foldline {
				split_index = i
				break
			}
		}
		points_to_transform := points[split_index:]
		// update points to be the "original" first half, i.e. the points being "folded onto"
		points = points[:split_index]
		// translate the y-coordinate of the points we're putting on "top" of the "original"
		// first half
		for _, v := range points_to_transform {
			v.x -= 2 * (v.x - foldline)
			// if transformed point already exists in "original" points, they merge
			// else add the point to the list of "original" points
			if !contains(points, v) {
				points = append(points, v)
			}
		}
	}

	return points
}

func PrintAfterFolds(points []Point) {
	// find largest coordinates
	cols := 0
	rows := 0
	for _, p := range points {
		if p.x > cols {
			cols = p.x
		}
		if p.y > rows {
			rows = p.y
		}
	}

	// since coords are 0-based
	page := make([][]string, rows+1)
	for i := range page {
		page[i] = make([]string, cols+1)
	}

	for r := 0; r < len(page); r++ {
		for c := 0; c < len(page[r]); c++ {
			page[r][c] = "."
		}
	}

	for _, p := range points {
		page[p.y][p.x] = "#"
	}

	for i := 0; i < len(page); i++ {
		fmt.Printf("%s\n", page[i])
	}

}

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var points []Point

	for scanner.Scan() {
		line := scanner.Text()
		// skip blank line between points and instructions
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold along") {
			instructions := strings.Fields(line)
			// instructions[0] and [1] = "fold along"
			temp := strings.Split(instructions[2], "=")
			foldline, _ := strconv.Atoi(temp[1])
			switch temp[0] {
			case "x":
				points = FoldPaper(points, Vertical, foldline)
			case "y":
				points = FoldPaper(points, Horizontal, foldline)
			}
		} else {
			coords := strings.Split(line, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			points = append(points, Point{x, y})
		}
	}

	fmt.Printf("number of visible dots after 1 fold = %d\n", len(points))

	PrintAfterFolds(points)
}
