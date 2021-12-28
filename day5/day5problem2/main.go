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
	return ln.p1.x == ln.p2.x
}

func (ln *Line) IsVertical() bool {
	return ln.p1.y == ln.p2.y
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func (ln *Line) IsDiagonal() bool {
	return abs(ln.p1.x, ln.p2.x) == abs(ln.p1.y, ln.p2.y)
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

	for _, xv := range v.points {
		for _, count := range xv {
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

		// only horizontal, vertical or diagonal lines
		if !(line.IsHorizontal() || line.IsVertical() || line.IsDiagonal()) {
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
		}
		if line.IsVertical() {
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
		if line.IsDiagonal() {
			var xs []int
			var ys []int
			if line.p1.x < line.p2.x {
				for v := line.p1.x; v <= line.p2.x; v++ {
					xs = append(xs, v)
				}
			} else {
				for v := line.p1.x; v >= line.p2.x; v-- {
					xs = append(xs, v)
				}
			}
			if line.p1.y < line.p2.y {
				for v := line.p1.y; v <= line.p2.y; v++ {
					ys = append(ys, v)
				}
			} else {
				for v := line.p1.y; v >= line.p2.y; v-- {
					ys = append(ys, v)
				}
			}
			for i, x := range xs {
				ventFieldMap.AddPoint(Point{x, ys[i]})
			}
		}
	}

	fmt.Printf("overlap count = %d\n", ventFieldMap.OverlapCount(2))
}
