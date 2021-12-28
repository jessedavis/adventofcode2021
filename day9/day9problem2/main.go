package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// TODO: this code is a mess, refactor

type Point struct {
	i        int
	j        int
	value    int
	explored bool
}

type Coord struct {
	i int
	j int
}

func AdjacentCoords(x Point, num_rows int, num_cols int) []Coord {
	possible_coords := make([]Coord, 0)

	// all the weird combos
	// top left
	if x.i == 0 && x.j == 0 {
		return append(possible_coords,
			Coord{0, 1},
			Coord{1, 0})
	}
	// top right
	if x.i == 0 && x.j == num_cols-1 {
		return append(possible_coords,
			Coord{0, x.j - 1},
			Coord{1, x.j})
	}
	// bottom left
	if x.i == num_rows-1 && x.j == 0 {
		return append(possible_coords,
			Coord{x.i - 1, 0},
			Coord{x.i, 1})
	}
	// bottom right
	if x.i == num_rows-1 && x.j == num_cols-1 {
		return append(possible_coords,
			Coord{x.i, x.j - 1},
			Coord{x.i - 1, x.j})
	}
	// top edge
	if x.i == 0 && x.j > 0 && x.j < num_cols-1 {
		return append(possible_coords,
			Coord{0, x.j - 1},
			Coord{1, x.j},
			Coord{0, x.j + 1})
	}
	// bottom edge
	if x.i == num_rows-1 && x.j > 0 && x.j < num_cols-1 {
		return append(possible_coords,
			Coord{x.i, x.j - 1},
			Coord{x.i - 1, x.j},
			Coord{x.i, x.j + 1})
	}
	// left edge
	if x.j == 0 && x.i > 0 && x.i < num_rows-1 {
		return append(possible_coords,
			Coord{x.i - 1, 0},
			Coord{x.i, 1},
			Coord{x.i + 1, 0})
	}
	// right edge
	if x.j == num_cols-1 && x.i > 0 && x.i < num_rows-1 {
		return append(possible_coords,
			Coord{x.i - 1, x.j},
			Coord{x.i, x.j - 1},
			Coord{x.i + 1, x.j})
	}
	// everything else
	return append(possible_coords,
		Coord{x.i - 1, x.j},
		Coord{x.i, x.j + 1},
		Coord{x.i + 1, x.j},
		Coord{x.i, x.j - 1})
}

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	locations := make([][]Point, 0)

	line_number := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]Point, len(line))
		for i, c := range line {
			location, _ := strconv.Atoi(string(c))
			digits[i] = Point{line_number, i, location, false}
		}
		locations = append(locations, digits)
		line_number++
	}

	lowpoints := make([]Point, 0)

	num_rows := len(locations)
	num_cols := len(locations[0])

	// top left
	value := locations[0][0].value
	if value < locations[0][1].value && value < locations[1][0].value {
		lowpoints = append(lowpoints, Point{0, 0, value, false})
	}
	// top right
	value = locations[0][num_cols-1].value
	if value < locations[0][num_cols-2].value && value < locations[1][num_cols-1].value {
		lowpoints = append(lowpoints, Point{0, num_cols - 1, value, false})
	}
	// bottom left
	value = locations[num_rows-1][0].value
	if value < locations[num_rows-2][0].value && value < locations[num_rows-1][1].value {
		lowpoints = append(lowpoints, Point{num_rows - 1, 0, value, false})
	}
	// bottom right
	value = locations[num_rows-1][num_cols-1].value
	if value < locations[num_rows-1][num_cols-2].value && value < locations[num_rows-2][num_cols-1].value {
		lowpoints = append(lowpoints, Point{num_rows - 1, num_cols - 1, value, false})
	}
	// top edge
	for j := 1; j < num_cols-1; j++ {
		value = locations[0][j].value
		if value < locations[0][j-1].value && value < locations[1][j].value && value < locations[0][j+1].value {
			lowpoints = append(lowpoints, Point{0, j, value, false})
		}
	}
	// bottom edge
	for j := 1; j < num_cols-1; j++ {
		i := num_rows - 1
		value = locations[i][j].value
		if value < locations[i][j-1].value && value < locations[i-1][j].value && value < locations[i][j+1].value {
			lowpoints = append(lowpoints, Point{i, j, value, false})
		}
	}
	// left edge
	for i := 1; i < num_rows-1; i++ {
		value = locations[i][0].value
		if value < locations[i-1][0].value && value < locations[i][1].value && value < locations[i+1][0].value {
			lowpoints = append(lowpoints, Point{i, 0, value, false})
		}
	}
	// right edge
	for i := 1; i < num_rows-1; i++ {
		j := num_cols - 1
		value = locations[i][j].value
		if value < locations[i-1][j].value && value < locations[i][j-1].value && value < locations[i+1][j].value {
			lowpoints = append(lowpoints, Point{i, j, value, false})
		}
	}

	for i := 1; i < num_rows-1; i++ {
		for j := 1; j < num_cols-1; j++ {
			value = locations[i][j].value
			// all other spots
			if value < locations[i-1][j].value &&
				value < locations[i][j+1].value &&
				value < locations[i+1][j].value &&
				value < locations[i][j-1].value {
				lowpoints = append(lowpoints, Point{i, j, value, false})
			}
		}
	}

	/*
		sum := 0
		for _, v := range lowpoints {
			sum += v.value + 1
		}
		fmt.Printf("sum of lowpoints = %d\n", sum)
	*/

	// going to assume that basins will be distinct around each low point, i.e.
	// one basin will containg one and only one lowpoint (although this is in the
	// description),
	basin_sizes := make([]int, len(lowpoints))

	// https://en.wikipedia.org/wiki/Breadth-first_search
	queue := make([]Point, 0)
	var x Point

	//LowPoints:
	for i_lp, lp := range lowpoints {
		locations[lp.i][lp.j].explored = true
		basin_sizes[i_lp] = 1
		queue = append(queue, lp)

		fmt.Printf("lp: i=%d, j=%d\n", lp.i, lp.j)
		for len(queue) > 0 {
			// pop
			if len(queue) > 500 {
				break
			}
			x, queue = queue[0], queue[1:]
			fmt.Printf("len queue = %d\n", len(queue))

			// nned to figure out how to mark as explored better
			possible_coords := AdjacentCoords(x, num_rows, num_cols)
			for _, pc := range possible_coords {
				p := &locations[pc.i][pc.j]
				// throw out 9s
				fmt.Printf("i = %d, j = %d\n", pc.i, pc.j)
				if p.value == 9 {
					p.explored = true
					continue
				}
				if !p.explored {
					p.explored = true
					basin_sizes[i_lp]++
					queue = append(queue, *p)
				}
			}

		}
	}

	for i, v := range basin_sizes {
		fmt.Printf("base size of low point %d = %d\n", i+1, v)
	}
	// reverse sort
	sort.Slice(basin_sizes, func(i, j int) bool {
		return basin_sizes[i] > basin_sizes[j]
	})
	//for _, v := range basin_sizes {
	//fmt.Printf("%d ", v)
	//}
	//fmt.Printf("\n")
	// get product of 3 largest basins
	product := 1
	for _, v := range basin_sizes[0:3] {
		product *= v
	}
	fmt.Printf("product of 3 largest basins = %d\n", product)
}
