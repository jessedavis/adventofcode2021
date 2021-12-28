package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part2exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var horizontalPosition int = 0
	var depth int = 0
	var aim int = 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var direction string = line[0]
		var magnitude int
		magnitude, _ = strconv.Atoi(line[1])
		fmt.Printf("%s %d\n", direction, magnitude)

		switch direction {
		case "forward":
			horizontalPosition += magnitude
			depth += aim * magnitude
		case "down":
			aim += magnitude
		case "up":
			aim -= magnitude
		}
	}
	fmt.Printf("horizontal = %d, depth = %d, final product = %d\n", horizontalPosition, depth, horizontalPosition*depth)
}
