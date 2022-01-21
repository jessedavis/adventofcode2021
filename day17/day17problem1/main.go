package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Velocity struct {
	xvel int
	yvel int
	xpos int
	ypos int
}

func (v *Velocity) adjustTrajectory() {
	v.xpos += v.xvel
	v.ypos += v.yvel

	if v.xvel != 0 {
		if v.xvel > 0 {
			v.xvel -= 1
		} else {
			v.xvel += 1
		}
	}
	v.yvel -= 1
}

func (v *Velocity) withinTargetArea(xmin, xmax, ymin, ymax int) bool {
	if v.xpos >= xmin && v.xpos <= xmax && v.ypos >= ymin && v.ypos <= ymax {
		return true
	}
	return false
}

func (v *Velocity) pastTargetArea(xmin, xmax, ymin, ymax int) bool {
	// case 1, x velocity is 0 and we're not between the sides of the area
	// which means we're going to undershoot or overshoot, regardless of where
	// we lie in the y range
	if v.xvel == 0 && (v.xpos < xmin || v.xpos > xmax) {
		return true
	}
	// ymin is bottom edge if ymin is negative, bit confusing
	if v.ypos < ymin {
		return true
	}
	// we're past the right edge, so even if x slowed down, we'd still be past
	if v.xpos > xmax {
		return true
	}

	// we're not past it (yet)
	return false
}

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	re := regexp.MustCompile(`target area: x=(?P<xone>-?\d+)..(?P<xtwo>-?\d+), y=(?P<yone>-?\d+)..(?P<ytwo>-?\d+)`)
	match := re.FindStringSubmatch(line)

	xmin, _ := strconv.Atoi(match[1])
	xmax, _ := strconv.Atoi(match[2])
	// ymin is "bottom" edge if negative, kinda confusing
	ymin, _ := strconv.Atoi(match[3])
	ymax, _ := strconv.Atoi(match[4])

	fmt.Printf("%d %d %d %d\n", xmin, xmax, ymin, ymax)

	// needed hint, tried deteremining formula by self
	// need to work on induction by just writing down values and determining pattern
	// but, height will be basically triangle
	// velocity = n, n-1...0, and then back down
	// so triangular number, which hints pointed out
	// height will want to land in the box
	// so the diff between that and 0,0 needs to be added to the height
	// at most can be botton edge
	// basically shifting the triangle "down"
	// t(n) = n(n+1)/2 so probably n-1 here since 0-based?
	// explained in reddit thread (-ymin - 1)
	// so this is slightly wrong (but right)

	var n int
	if ymin < 0 {
		n = -ymin
	}
	max_height := n * (n - 1) / 2
	fmt.Printf("max height = %d\n", max_height)

	// second part will require more thought

	// xvel slows down by 1, each time, so n, n-1, n-2
	// so min xvel needs to be able to hit xmin
	// max xvel is xmax in order to pass through
	min_x_vel := 0
	max_x_vel := xmax
	for min_x_vel*(min_x_vel+1)/2 < xmin {
		min_x_vel++
	}
	// range of yvel is -ymin to ymin, since ymin negative here, starting with ymin, not -ymin, bit confusing tbh
	possible_velocities := make([]Velocity, 0)
	for x := min_x_vel; x <= max_x_vel; x++ {
		for y := ymin; y <= -ymin; y++ {
			possible_velocities = append(possible_velocities, Velocity{xvel: x, yvel: y})
		}
	}

	num_distinct_values := 0

velocities:
	for _, pv := range possible_velocities {
		for {
			if pv.withinTargetArea(xmin, xmax, ymin, ymax) {
				num_distinct_values++
				continue velocities
			}
			if pv.pastTargetArea(xmin, xmax, ymin, ymax) {
				continue velocities
			}
			pv.adjustTrajectory()
		}
	}

	fmt.Printf("number of distinct initial velocity values = %d\n", num_distinct_values)
}
