package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// from container/heap example

// IntHeap will be a min-heap of Point pointers, sorted by Points.f .
type IntHeap []*Point

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].f < h[j].f }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Point))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Point struct {
	parent  *Point
	risk    int
	row     int
	col     int
	f       int // total cost of current node
	g       int // cost from start to current node
	h       int // heuristic estimated cost from current node to end
	visited bool
}

func findAdjacent(p Point, pointMap [][]Point) []*Point {
	adjacent := make([]*Point, 0, 4)

	max_rows := len(pointMap)
	max_cols := len(pointMap[0])

	// all the weird combos
	// top left
	if p.row == 0 && p.col == 0 {
		return append(adjacent,
			&pointMap[0][1],
			&pointMap[1][0])
	}
	// top right
	if p.row == 0 && p.col == max_cols-1 {
		return append(adjacent,
			&pointMap[0][p.col-1],
			&pointMap[1][p.col])
	}
	// bottom left
	if p.row == max_rows-1 && p.col == 0 {
		return append(adjacent,
			&pointMap[p.row-1][0],
			&pointMap[p.row][1])
	}
	// bottom right
	if p.row == max_rows-1 && p.col == max_cols-1 {
		return append(adjacent,
			&pointMap[p.row][p.col-1],
			&pointMap[p.row-1][p.col])
	}
	// top edge
	if p.row == 0 && p.col > 0 && p.col < max_cols-1 {
		return append(adjacent,
			&pointMap[0][p.col-1],
			&pointMap[1][p.col],
			&pointMap[0][p.col+1])
	}
	// bottom edge
	if p.row == max_rows-1 && p.col > 0 && p.col < max_cols-1 {
		return append(adjacent,
			&pointMap[p.row][p.col-1],
			&pointMap[p.row-1][p.col],
			&pointMap[p.row][p.col+1])
	}
	// left edge
	if p.col == 0 && p.row > 0 && p.row < max_rows-1 {
		return append(adjacent,
			&pointMap[p.row-1][0],
			&pointMap[p.row][1],
			&pointMap[p.row+1][0])
	}
	// right edge
	if p.col == max_cols-1 && p.row > 0 && p.row < max_rows-1 {
		return append(adjacent,
			&pointMap[p.row-1][p.col],
			&pointMap[p.row][p.col-1],
			&pointMap[p.row+1][p.col])
	}
	// everything else
	return append(adjacent,
		&pointMap[p.row-1][p.col],
		&pointMap[p.row][p.col+1],
		&pointMap[p.row+1][p.col],
		&pointMap[p.row][p.col-1])

}

func pathLocations(p Point) [][]int {
	cur := p

	result := make([][]int, 0)

	for cur.parent != nil {
		location := []int{cur.row, cur.col}
		result = append(result, location)
		cur = *cur.parent
	}

	// append the starting position
	result = append(result, []int{0, 0})
	// and reverse the array
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func pathSum(pointmap [][]Point, path [][]int) int {
	sum := 0

	// we don't enter the first point (0,0), so don't count it
	for _, pos := range path[1:] {
		row := pos[0]
		col := pos[1]
		sum += pointmap[row][col].risk
	}

	return sum
}

// https://en.wikipedia.org/wiki/A*_search_algorithm
// first attempt was ok for sample data, followed
// https://towardsdatascience.com/a-star-a-search-algorithm-eb495fb156bb
// but didn't return for real data, so had to set up min heap
// and follow Wikipedia article

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	riskPoints := make([][]Point, 0)

	// initialize f and g with really big values
	default_f := 99999999
	default_g := 99999999

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		points := make([]Point, len(line))
		for i, c := range strings.Split(line, "") {
			v, _ := strconv.Atoi(c)
			points[i] = Point{nil, v, row, i, default_f, default_g, 0, false}
		}
		riskPoints = append(riskPoints, points)
		row++
	}

	start_pos := riskPoints[0][0]
	end_pos := riskPoints[len(riskPoints)-1][len(riskPoints[0])-1]

	// initialize start node correctly
	start_pos.f = end_pos.row + end_pos.col
	start_pos.g = 0

	var locations [][]int

	to_visit := &IntHeap{&start_pos}
	heap.Init(to_visit)

	for to_visit.Len() > 0 {
		// find node to visit with lowest f score, since we're a min heap, it's the first element
		// could also just pop it too
		current := (*to_visit)[0]
		// setting visited on Point so that we don't have to look for inclusion of Point in heap
		current.visited = true
		// if we're at the end, we made it
		if current.row == end_pos.row && current.col == end_pos.col {
			locations = pathLocations(*current)
			break
		}
		// pop current, we don't really need it, if we did, it'd be
		// current = heap.Pop(to_visit).(*Point)
		// wikipedia lists as Remove
		heap.Pop(to_visit)

		// find neighbors
		adjacent := findAdjacent(*current, riskPoints)
		for _, adj := range adjacent {
			// already visited this adjacent node, move on
			if adj.visited {
				continue
			}

			// h value heuristic = manhattan (step) distance here to end, assume 1 per step
			// https://en.wikipedia.org/wiki/A*_search_algorithm
			// If the heuristic function is admissible, meaning that it never
			// overestimates the actual cost to get to the goal, A* is guaranteed
			// to return a least-cost path from start to goal.

			if current.g+adj.risk < adj.g {
				// better path, record it
				adj.parent = current
				adj.g = current.g + adj.risk
				adj.h = (end_pos.row - adj.row) + (end_pos.col - adj.col)
				adj.f = adj.g + adj.h

				heap.Push(to_visit, adj)
			}
		}
	}

	fmt.Printf("path = %+v\n", locations)
	fmt.Printf("total of lowest risk path = %d\n", pathSum(riskPoints, locations))
}
