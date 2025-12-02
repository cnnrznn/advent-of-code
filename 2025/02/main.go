package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cnnrznn/advent-of-code/util"
)

func main() {
	lines, err := util.ReadInput()
	if err != nil {
		panic(err)
	}

	var result int

	for _, l := range lines {
		n := processLine(l)
		result += n
	}

	fmt.Println(result)
}

func processLine(l string) int {
	var result int
	ranges := strings.Split(l, ",")
	for _, r := range ranges {
		first, _ := strconv.Atoi(strings.Split(r, "-")[0])
		last, _ := strconv.Atoi(strings.Split(r, "-")[1])
		result += countInvalidIDs(first, last)
	}
	return result
}

func countInvalidIDs(first, last int) int {
	var result int
	for i := first; i <= last; i++ {
		if isInvalidIDpart2(i) {
			result += i
		}
	}
	return result
}

func isInvalidID(x int) bool {
	s := strconv.Itoa(x)
	
	if len(s) % 2 != 0 {
		return false
	}

	mid := len(s) / 2
	
	return s[:mid] == s[mid:]
}

func isInvalidIDpart2(x int) bool {
	s := strconv.Itoa(x)
	
	// start with a window
	for w := 1; w <= len(s); w++ {
		substr := s[:w]
		if isRepeated(s[w:], substr) {
			return true
		}
	}

	return false
}

func isRepeated(s, sub string) bool {
	// see if that window is repeated for the rest of the string
	var (
		isRepeated bool
		size = len(sub)
	)

	// sub cannot be repeated in s if it is not divisible
	if len(s) % len(sub) != 0 {
		return false
	}

	// check each window for equality
	for len(s) >= len(sub) {
		if s[:size] == sub {
			isRepeated = true
			s = s[size:]
		} else {
			isRepeated = false
			break
		}
	}

	return isRepeated
}