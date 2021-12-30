package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	row    int
	col    int
	energy int
}

func findAdjacentPoints(p *Point, octopuses [][]Point) []*Point {
	adjacent_points := make([]*Point, 0)

	rows := len(octopuses)
	cols := len(octopuses[0])

	// from day 9, second part, but with diagonals
	// top left
	if p.row == 0 && p.col == 0 {
		return append(adjacent_points,
			&octopuses[0][1],
			&octopuses[1][1],
			&octopuses[1][0])
	}
	// top right
	if p.row == 0 && p.col == cols-1 {
		return append(adjacent_points,
			&octopuses[0][p.col-1],
			&octopuses[1][p.col-1],
			&octopuses[1][p.col])
	}
	// bottom left
	if p.row == rows-1 && p.col == 0 {
		return append(adjacent_points,
			&octopuses[p.row-1][0],
			&octopuses[p.row-1][1],
			&octopuses[p.row][1])
	}
	// bottom right
	if p.row == rows-1 && p.col == cols-1 {
		return append(adjacent_points,
			&octopuses[p.row][p.col-1],
			&octopuses[p.row-1][p.col-1],
			&octopuses[p.row-1][p.col])
	}
	// top edge
	if p.row == 0 && p.col > 0 && p.col < cols-1 {
		return append(adjacent_points,
			&octopuses[0][p.col-1],
			&octopuses[1][p.col-1],
			&octopuses[1][p.col],
			&octopuses[1][p.col+1],
			&octopuses[0][p.col+1])
	}
	// bottom edge
	if p.row == rows-1 && p.col > 0 && p.col < cols-1 {
		return append(adjacent_points,
			&octopuses[p.row][p.col-1],
			&octopuses[p.row-1][p.col-1],
			&octopuses[p.row-1][p.col],
			&octopuses[p.row-1][p.col+1],
			&octopuses[p.row][p.col+1])
	}
	// left edge
	if p.col == 0 && p.row > 0 && p.row < rows-1 {
		return append(adjacent_points,
			&octopuses[p.row-1][0],
			&octopuses[p.row-1][1],
			&octopuses[p.row][1],
			&octopuses[p.row+1][1],
			&octopuses[p.row+1][0])
	}
	// right edge
	if p.col == cols-1 && p.row > 0 && p.row < rows-1 {
		return append(adjacent_points,
			&octopuses[p.row-1][p.col],
			&octopuses[p.row-1][p.col-1],
			&octopuses[p.row][p.col-1],
			&octopuses[p.row+1][p.col-1],
			&octopuses[p.row+1][p.col])
	}
	// everything else
	return append(adjacent_points,
		&octopuses[p.row-1][p.col-1],
		&octopuses[p.row-1][p.col],
		&octopuses[p.row-1][p.col+1],
		&octopuses[p.row][p.col+1],
		&octopuses[p.row+1][p.col+1],
		&octopuses[p.row+1][p.col],
		&octopuses[p.row+1][p.col-1],
		&octopuses[p.row][p.col-1])
}

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	octopuses := make([][]Point, 0)

	row := 0
	for scanner.Scan() {
		points := make([]Point, 0)
		for col, c := range strings.Split(scanner.Text(), "") {
			v, _ := strconv.Atoi(c)
			points = append(points, Point{row, col, v})
		}
		octopuses = append(octopuses, points)
		row++
	}

	steps := 100
	flashes := 0
	flash_queue := make([]*Point, 0)
	for step := 0; step < steps; step++ {
		// first, the energy level of each octopus increases by 1
		for row := 0; row < len(octopuses); row++ {
			for col := 0; col < len(octopuses[0]); col++ {
				pos := &octopuses[row][col]
				(*pos).energy++
				// if we're going to flash, add to queue
				if (*pos).energy > 9 {
					flash_queue = append(flash_queue, pos)
				}
			}
		}

		// next, flash if energy > 9
		var to_flash *Point
		for len(flash_queue) > 0 {
			to_flash, flash_queue = flash_queue[0], flash_queue[1:]
			if (*to_flash).energy > 9 {
				flashes++
				// reset energy to 0, since we flashed
				(*to_flash).energy = 0
				// find adjacent octopuses
				adjacent_octopuses := findAdjacentPoints(to_flash, octopuses)
				for _, octo := range adjacent_octopuses {
					// bump energy of adjacent, unless it's 0
					// if it's zero, that means it already flashed this round
					// and we only flash once per step
					if (*octo).energy > 0 {
						(*octo).energy++
					}
					// if adjacent octopus needs to flash, add it to queue
					if (*octo).energy > 9 {
						flash_queue = append(flash_queue, octo)
					}
				}
			}

		}
		fmt.Printf("after %d steps, number of flashes = %d\n", step+1, flashes)
	}

}
