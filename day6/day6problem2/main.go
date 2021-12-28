package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func TotalFish(m map[int]int) int {
	total := 0

	for _, v := range m {
		total += v
	}

	return total
}

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	Fish := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	simulationLength := 256
	//simulationLength := 80
	//simulationLength := 18

	for scanner.Scan() {
		for _, v := range strings.Split(scanner.Text(), ",") {
			initialTimer, _ := strconv.Atoi(v)
			Fish[initialTimer]++
		}
	}

	// dynamic programming, exponential = cache!

	for day := 1; day <= simulationLength; day++ {
		newFishToCreate := Fish[0]
		for i := 0; i <= 7; i++ {
			Fish[i] = Fish[i+1]
		}
		Fish[6] += newFishToCreate
		Fish[8] = newFishToCreate
		fmt.Printf("Number of fish after %d days = %d\n", day, TotalFish(Fish))
	}

	fmt.Printf("Number of fish after %d days = %d\n", simulationLength, TotalFish(Fish))
}
