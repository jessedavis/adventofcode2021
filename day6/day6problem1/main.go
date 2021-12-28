package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Fish struct {
	timer int
}

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var FishSchool []Fish
	simulationLength := 80
	//simulationLength := 18

	for scanner.Scan() {
		for _, v := range strings.Split(scanner.Text(), ",") {
			initialTimer, _ := strconv.Atoi(v)
			FishSchool = append(FishSchool, Fish{initialTimer})
		}
	}

	for day := 1; day <= simulationLength; day++ {
		//fmt.Printf("After %d days: ", day)
		for i := range FishSchool {
			if FishSchool[i].timer == 0 {
				FishSchool = append(FishSchool, Fish{8})
				FishSchool[i].timer = 6
			} else {
				FishSchool[i].timer--
			}
		}

		/*
			fmt.Printf("After %d days: ", day)
			for i := range FishSchool {
				fmt.Printf("%d,", FishSchool[i].timer)
			}
			fmt.Printf("\n")
		*/
	}

	fmt.Printf("Number of fish after %d days = %d\n", simulationLength, len(FishSchool))
}
