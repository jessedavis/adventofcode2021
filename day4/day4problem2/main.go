package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CardCell struct {
	value  string
	marked bool
}

type Card struct {
	values [][]CardCell
}

func (c *Card) FillCard(numbers []string) {
	c.values = make([][]CardCell, 5)
	for row := 0; row < 5; row++ {
		c.values[row] = make([]CardCell, 5)
	}
	for i, n := range numbers {
		c.values[i/5][i%5] = CardCell{n, false}
	}
}

func (c *Card) MarkNumber(number string) {
	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			if c.values[row][column].value == number {
				c.values[row][column].marked = true
			}
		}
	}
}

// https://gobyexample.com/collection-functions
func All(values []CardCell, f func(CardCell) bool) bool {
	for _, v := range values {
		if !f(v) {
			return false
		}
	}
	return true
}

func (c *Card) IsWinner() bool {
	marked := func(cell CardCell) bool { return cell.marked }

	// check rows
	for row := 0; row < 5; row++ {
		if All(c.values[row], marked) {
			return true
		}
	}

	temp := make([]CardCell, 5)
	// check columns
	for column := 0; column < 5; column++ {
		for i := 0; i < 5; i++ {
			temp[i] = c.values[i][column]
		}
		if All(temp, marked) {
			return true
		}
	}
	// we're not dealing with diagonals
	/*
		for pos := 0; pos < 5; pos++ {
			temp[pos] = c.values[pos][pos]
			if All(temp, marked) {
				return true
			}
		}
		for pos := 0; pos < 5; pos++ {
			temp[pos] = c.values[4-pos][pos]
			if All(temp, marked) {
				return true
			}
		}
	*/

	return false
}

func (c *Card) Score(calledNumber string) int {
	sum := 0
	cn, _ := strconv.Atoi(calledNumber)

	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			if !c.values[row][column].marked {
				v, _ := strconv.Atoi(c.values[row][column].value)
				sum += v
			}
		}
	}

	return sum * cn
}

func (c *Card) PrintCard() {
	colorRed := "\033[31m"
	colorReset := "\033[0m"

	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			if c.values[row][column].marked {
				fmt.Printf("%s%s%s ", colorRed, c.values[row][column].value, colorReset)
			} else {
				fmt.Printf("%s ", c.values[row][column].value)
			}
		}
		fmt.Printf("\n")
	}
}

func numberOfWinnerCards(status []bool) int {
	result := 0
	for _, v := range status {
		if v {
			result++
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

	scanner.Scan()
	calledNumbers := strings.Split(scanner.Text(), ",")

	fmt.Printf("calledNumbers = %q\n", calledNumbers)

	var linesRead int = 0
	var cardNumbers []string
	var cards []Card

	for scanner.Scan() {
		line := scanner.Text()

		// card = 1 blank line and 5 lines
		if linesRead%6 == 0 && line == "" {
			cardNumbers = nil
		} else {
			cardNumbers = append(cardNumbers, strings.Fields(line)...)
			if linesRead%6 == 5 {
				fmt.Printf("card = %q\n", cardNumbers)
				var card Card
				card.FillCard(cardNumbers)
				cards = append(cards, card)
			}
		}

		linesRead++
	}

	var lastWinningCardIndex int = -1
	var winningCalledNumber string
	cardWonStatus := make([]bool, len(cards))

callNumbers:
	for _, calledNumber := range calledNumbers {
		for i := range cards {
			// no real reason to bother with card that's already won
			if cardWonStatus[i] {
				continue
			}
			cards[i].MarkNumber(calledNumber)
			if cards[i].IsWinner() {
				cardWonStatus[i] = true
			}
			if numberOfWinnerCards(cardWonStatus) == len(cards) {
				// found the last card that won
				lastWinningCardIndex = i
				winningCalledNumber = calledNumber
				break callNumbers
			}
		}
	}

	if lastWinningCardIndex >= 0 {
		fmt.Printf("last winning called number = %s\n", winningCalledNumber)
		fmt.Printf("winning card\n")
		cards[lastWinningCardIndex].PrintCard()
		// generate sum
		score := cards[lastWinningCardIndex].Score(winningCalledNumber)
		fmt.Printf("last winning score = %d\n", score)
	} else {
		fmt.Printf("no winning card\n")
	}
}
