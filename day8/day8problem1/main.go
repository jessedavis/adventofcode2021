package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledatalarge.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var numberEasyDigits int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		//patterns := strings.Fields(line[0])
		digits := strings.Fields(line[1])

		for _, v := range digits {
			l := len(v)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				numberEasyDigits++
			}
		}
	}
	fmt.Printf("number of 1, 4, 7 or 8s = %d\n", numberEasyDigits)
}
