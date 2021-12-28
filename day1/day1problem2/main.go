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
	//f, err := os.Open("../part2exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var linesRead int = 0
	var numWindowIncreases = 0

	var window [4]int

	for scanner.Scan() {
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			linesRead++
			if linesRead <= 3 {
				window[linesRead-1] = i
				continue
			}

			window[3] = i

			var windowOneSum = window[0] + window[1] + window[2]
			var windowTwoSum = window[1] + window[2] + window[3]

			if windowTwoSum > windowOneSum {
				numWindowIncreases++
			}
			fmt.Printf("Sum 1 = %d, Sum 2 = %d, incs = %d\n", windowOneSum, windowTwoSum, numWindowIncreases)
			// and shift down
			for j := 0; j < 3; j++ {
				window[j] = window[j+1]
			}
			window[3] = 0
		}
	}

	fmt.Printf("number of increases : %d\n", numWindowIncreases)
}
