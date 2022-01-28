package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// giving up on tree representation after hint on reddit
// to just manipulate strings

func Explode(s string) (string, bool) {
	pairOpenCount := 0
	i := 0

out:
	for _, c := range s {
		switch string(c) {
		case "[":
			pairOpenCount++
			if pairOpenCount == 5 {
				break out
			}
		case "]":
			pairOpenCount--
		}
		i++
	}

	if pairOpenCount < 5 {
		return s, false
	}

	// else we need to split around "[d,d]"
	pairStart := i
	pair_re := regexp.MustCompile(`\[\d+,\d+\]`)
	loc := pair_re.FindStringIndex(s[pairStart:])

	left := s[0:pairStart]
	right := s[pairStart+loc[1]:]
	pairToExplode := strings.Split(s[pairStart+1:pairStart+loc[1]-1], ",")
	lval, _ := strconv.Atoi(pairToExplode[0])
	rval, _ := strconv.Atoi(pairToExplode[1])

	regnum_re := regexp.MustCompile(`\d+`)
	// find index of last number in the "left" remainder of the number
	leftNumberIndices := regnum_re.FindAllStringIndex(left, -1)
	rightNumberIndices := regnum_re.FindAllStringIndex(right, 1)

	var result string
	if leftNumberIndices != nil {
		x := leftNumberIndices[len(leftNumberIndices)-1]
		t, _ := strconv.Atoi(left[x[0]:x[1]])
		newLeft := strconv.Itoa(t + lval)
		result = left[0:x[0]] + newLeft + left[x[1]:]
	} else {
		result = left
	}

	result = result + "0"

	if rightNumberIndices != nil {
		x := rightNumberIndices[0]
		t, _ := strconv.Atoi(right[x[0]:x[1]])
		newRight := strconv.Itoa(t + rval)
		result = result + right[0:x[0]] + newRight + right[x[1]:]
	} else {
		result = result + right
	}

	return result, true
}

func Split(s string) (string, bool) {
	digitPattern := regexp.MustCompile(`\d{2,}`)
	tooBig := digitPattern.FindStringIndex(s)

	if tooBig != nil {
		left := s[0:tooBig[0]]
		right := s[tooBig[1]:]

		n, _ := strconv.Atoi(s[tooBig[0]:tooBig[1]])
		l := n / 2
		r := 0
		if n%2 == 0 {
			r = n / 2
		} else {
			r = n/2 + 1
		}
		pair := "[" + string(strconv.Itoa(l)) + "," + string(strconv.Itoa(r)) + "]"

		return string(left) + pair + string(right), true
	} else {
		return s, false
	}
}

func Add(s1, s2 string) string {
	result := "[" + s1 + "," + s2 + "]"

	for {
		var exploded bool
		var splitted bool
		result, exploded = Explode(result)
		if exploded {
			continue
		}
		result, splitted = Split(result)

		if !exploded && !splitted {
			break
		}
	}

	return result
}

func Magnitude(s string) int {
	pair_re := regexp.MustCompile(`\[\d+,\d+\]`)
	loc := pair_re.FindStringIndex(s)
	before_pair := s[0:loc[0]]
	after_pair := s[loc[1]:]

	t := strings.Split(s[loc[0]+1:loc[1]-1], ",")
	l, _ := strconv.Atoi(t[0])
	r, _ := strconv.Atoi(t[1])

	x := 3*l + 2*r

	if len(before_pair) == 0 && len(after_pair) == 0 {
		// we're done
		return x
	} else {
		return Magnitude(string(before_pair) + strconv.Itoa(x) + string(after_pair))
	}
}

func Sum(fish []string) string {
	sum := ""

	for i, f := range fish {
		if i == 0 {
			sum = fish[i]
		} else {
			sum = Add(sum, f)
		}
	}

	return sum
}

func LargestMagnitude(fish []string) int {
	largest := 0

	for i, f1 := range fish {
		for j, f2 := range fish {
			if i == j {
				// don't add against ourselves
				continue
			}
			sum := Add(f1, f2)
			mag := Magnitude(sum)
			if mag > largest {
				largest = mag
			}
		}
	}

	return largest
}

func main() {
	f, err := os.Open("input.txt")
	//f, err := os.Open("part1exampledata.txt")
	//f, err := os.Open("part1exampledata2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var fish []string

	for scanner.Scan() {
		line := scanner.Text()
		fish = append(fish, line)
	}

	snailfish_sum := Sum(fish)
	fmt.Printf("sum of snailfish numbers = %s\n", snailfish_sum)
	fmt.Printf("magnitude of snailfish final number = %d\n", Magnitude(snailfish_sum))
	fmt.Printf("largest magnitude of 2 snailfish = %d\n", LargestMagnitude(fish))
}
