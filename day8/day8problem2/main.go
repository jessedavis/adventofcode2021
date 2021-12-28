package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func commonChars(s string, contains string) int {
	count := 0
	for _, v := range contains {
		if strings.ContainsRune(s, v) {
			count++
		}
	}
	return count
}

// requires digitPatterns to have been initialized to all empty strings
func isDigitPatternSet(digitPatterns [10]string, digit int) bool {
	switch digit {
	case 0:
		return len(digitPatterns[0]) == 6
	case 1:
		return len(digitPatterns[1]) == 2
	case 2:
		return len(digitPatterns[2]) == 5
	case 3:
		return len(digitPatterns[3]) == 5
	case 4:
		return len(digitPatterns[4]) == 4
	case 5:
		return len(digitPatterns[5]) == 5
	case 6:
		return len(digitPatterns[6]) == 6
	case 7:
		return len(digitPatterns[7]) == 3
	case 8:
		return len(digitPatterns[8]) == 7
	case 9:
		return len(digitPatterns[9]) == 6
	}
	return false
}

// returns int corresponding to which digit this is
// returns 1 if somehow fell through (since 1 is a two-segment number)
func FiveDigitPattern(pattern string, digitPatterns [10]string) int {
	// easy cases, if segments for 1 or 7 are in pattern, then this is a 3
	if isDigitPatternSet(digitPatterns, 7) && commonChars(pattern, digitPatterns[7]) == 3 ||
		isDigitPatternSet(digitPatterns, 1) && commonChars(pattern, digitPatterns[1]) == 2 {
		return 3
	}
	// 4 bit harder
	// 3/4 of 4 could also be 5, so only able to discern if pattern contains 1 or 7
	if isDigitPatternSet(digitPatterns, 4) && commonChars(pattern, digitPatterns[4]) == 3 {
		// 3 = 3/4 of 4 AND (2/2 of 1 OR 3/3 of 7)
		if commonChars(pattern, digitPatterns[1]) == 2 || commonChars(pattern, digitPatterns[7]) == 3 {
			return 3
		}
		// 5 = 3/4 of 4 AND (1/2 of 1 OR 2/3 of 7)
		if commonChars(pattern, digitPatterns[1]) == 1 || commonChars(pattern, digitPatterns[7]) == 2 {
			return 5
		}
	}
	// if we've gotten this far, then mostly like a 2, but check just to be sure
	// 4 pretty much required to distinguish between 2 and 5, 1 and/or 7 not enough
	// 2 is (not 3 and not 5)
	if isDigitPatternSet(digitPatterns, 4) && commonChars(pattern, digitPatterns[4]) == 2 {
		return 2
	}

	// logic problem, so return "bad" value
	return 1
}

// returns int corresponding to which digit this is
// returns 1 if somehow fell through (since 1 is a two-segment number)
func SixDigitPattern(pattern string, digitPatterns [10]string) int {
	// first pass: use 5 digit segments if set
	if isDigitPatternSet(digitPatterns, 5) && commonChars(pattern, digitPatterns[5]) == 5 {
		// 6 = 5/5 of 5 and (1/2 of 1 OR 2/3 of 7)
		// 9 = 5/5 of 5 and (2/2 of 1 OR 3/3 of 7)
		if isDigitPatternSet(digitPatterns, 1) {
			if commonChars(pattern, digitPatterns[1]) == 1 {
				return 6
			}
			if commonChars(pattern, digitPatterns[1]) == 2 {
				return 9
			}
		}
		if isDigitPatternSet(digitPatterns, 7) {
			if commonChars(pattern, digitPatterns[7]) == 2 {
				return 6
			}
			if commonChars(pattern, digitPatterns[7]) == 3 {
				return 9
			}
		}
	}
	// if 3 set, and completely contained, then we're 9
	if isDigitPatternSet(digitPatterns, 3) && commonChars(pattern, digitPatterns[3]) == 5 {
		return 9
	}

	// now start with the "unique" segments

	// 4 is pretty much required
	// 1 and 7 not enough to distinguish 0 from 9
	// if 4 not set, then we have to rely on 3 and 5 having been set set as above
	if isDigitPatternSet(digitPatterns, 4) {
		// 9 = 4/4 of 4
		// OR 3/4 of 4 AND (2/2 of 1 or 3/3 of 7)
		// 6 = 3/4 of 4 AND (1/2 of 1 or 2/3 of 7)
		// 0 = 3/4 of 4 AND (2/2 of 1 or 3/3 of 7) (with caveat above)
		if commonChars(pattern, digitPatterns[4]) == 4 {
			return 9
		}
		if commonChars(pattern, digitPatterns[4]) == 3 {
			if isDigitPatternSet(digitPatterns, 1) {
				if commonChars(pattern, digitPatterns[1]) == 1 {
					return 6
				}
				if commonChars(pattern, digitPatterns[1]) == 2 {
					return 0
				}
			}
			if isDigitPatternSet(digitPatterns, 7) {
				if commonChars(pattern, digitPatterns[7]) == 2 {
					return 6
				}
				if commonChars(pattern, digitPatterns[7]) == 3 {
					return 0
				}
			}
		}
	}
	// logic problem, so return "bad" value
	return 1
}

func main() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../part1exampledatalarge.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		digitOutputs := []int{0, 0, 0, 0}
		signalPatterns := strings.Fields(line[0])
		digits := strings.Fields(line[1])

		digitPatterns := [...]string{"", "", "", "", "", "", "", "", "", ""}

		// unique patterns, will probably have to have at least one
		// of these in any line
		for _, v := range signalPatterns {
			switch len(v) {
			case 2:
				digitPatterns[1] = v
			case 3:
				digitPatterns[7] = v
			case 4:
				digitPatterns[4] = v
			case 7:
				digitPatterns[8] = v
			}
		}

		// 5 and 6 segment digits
		// 5 first, in case they can be used to distinguish 6 segment digits
		for _, v := range signalPatterns {
			switch len(v) {
			case 5:
				// 2, 3 or 5
				x := FiveDigitPattern(v, digitPatterns)
				if x != 1 {
					digitPatterns[x] = v
				} else {
					fmt.Printf("LOGIC ERROR for segment = 5, %s\n", v)
				}
			}
		}

		for _, v := range signalPatterns {
			switch len(v) {
			case 6:
				x := SixDigitPattern(v, digitPatterns)
				if x != 1 {
					digitPatterns[x] = v
				} else {
					fmt.Printf("LOGIC ERROR for segment = 6, %s\n", v)
				}
			}
		}

		for i, v := range digits {
			switch len(v) {
			case 2:
				digitOutputs[i] = 1
			case 3:
				digitOutputs[i] = 7
			case 4:
				digitOutputs[i] = 4
			case 7:
				digitOutputs[i] = 8
			case 5:
				// 2, 3 or 5
				if commonChars(v, digitPatterns[2]) == 5 {
					digitOutputs[i] = 2
				}
				if commonChars(v, digitPatterns[3]) == 5 {
					digitOutputs[i] = 3
				}
				if commonChars(v, digitPatterns[5]) == 5 {
					digitOutputs[i] = 5
				}
			case 6:
				// 0, 6 or 9
				if commonChars(v, digitPatterns[0]) == 6 {
					digitOutputs[i] = 0
				}
				if commonChars(v, digitPatterns[6]) == 6 {
					digitOutputs[i] = 6
				}
				if commonChars(v, digitPatterns[9]) == 6 {
					digitOutputs[i] = 9
				}
			}
		}
		outputValue := (digitOutputs[0] * 1000) + (digitOutputs[1] * 100) + (digitOutputs[2] * 10) + digitOutputs[3]
		sum += outputValue
		//fmt.Printf("output value of line = %d\n", outputValue)
	}

	fmt.Printf("final sum of output values = %d\n", sum)
}
