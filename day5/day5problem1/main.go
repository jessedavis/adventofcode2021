package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

type VentFields struct {
	points map[int]map[int]int
}

func (ln *Line) Print() {
	fmt.Printf("%d,%d  %d,%d\n", ln.p1.x, ln.p1.y, ln.p2.x, ln.p2.y)
}

func (ln *Line) IsHorizontal() bool {
	if ln.p1.x == ln.p2.x {
		return true
	} else {
		return false
	}
}

func (ln *Line) IsVertical() bool {
	if ln.p1.y == ln.p2.y {
		return true
	} else {
		return false
	}
}

func (v *VentFields) AddPoint(p Point) {
	if v.points == nil {
		v.points = make(map[int]map[int]int)
	}
	_, xok := v.points[p.x]
	if !xok {
		// need to add key
		vy := map[int]int{p.y: 1}
		v.points[p.x] = vy
	} else {
		v.points[p.x][p.y]++
	}
}

func (v *VentFields) OverlapCount(atleast int) int {
	result := 0

	for x, xv := range v.points {
		for y, count := range xv {
			fmt.Printf("x = %d, y = %d, c = %d\n", x, y, count)
			if count >= atleast {
				result++
			}
		}
	}

	return result
}

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	r := regexp.MustCompile(`(\d+)`)

	var ventFieldMap VentFields

	for scanner.Scan() {
		numbers := r.FindAllString(scanner.Text(), -1)
		p1x, _ := strconv.Atoi(numbers[0])
		p1y, _ := strconv.Atoi(numbers[1])
		p2x, _ := strconv.Atoi(numbers[2])
		p2y, _ := strconv.Atoi(numbers[3])
		line := Line{Point{p1x, p1y}, Point{p2x, p2y}}

		// only horizontal or vertical lines
		if !(line.IsHorizontal() || line.IsVertical()) {
			continue
		}

		//line.Print()

		// mark the points
		if line.IsHorizontal() {
			x := line.p1.x
			start := line.p1.y
			end := line.p2.y
			// if "backwards", then flip start and end
			if line.p1.y > line.p2.y {
				start = line.p2.y
				end = line.p1.y
			}
			for y := start; y <= end; y++ {
				ventFieldMap.AddPoint(Point{x, y})
			}
		} else {
			y := line.p1.y
			start := line.p1.x
			end := line.p2.x
			// if "backwards", then flip start and end
			if line.p1.x > line.p2.x {
				start = line.p2.x
				end = line.p1.x
			}
			for x := start; x <= end; x++ {
				ventFieldMap.AddPoint(Point{x, y})
			}
		}
	}

	fmt.Printf("overlap count = %d\n", ventFieldMap.OverlapCount(2))
}
