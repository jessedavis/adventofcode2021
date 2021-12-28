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
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var linesRead int = 0
	var prevMeasurement int = 0
	var increaseInMeasurement int = 0

	for scanner.Scan() {
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			linesRead++
			if linesRead == 1 {
				prevMeasurement = i
				continue
			}
			if i > prevMeasurement {
				increaseInMeasurement++
			}
			prevMeasurement = i
		}
	}

	fmt.Printf("number of increases : %d\n", increaseInMeasurement)
}
