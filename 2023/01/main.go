//--- Day 1: Trebuchet?! ---
//Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.
//
//You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.
//
//Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!
//
//You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").
//
//As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.
//
//The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.
//
//For example:
//
//1abc2
//pqr3stu8vwx
//a1b2c3d4e5f
//treb7uchet
//In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.
//
//Consider your entire calibration document. What is the sum of all of the calibration values?

package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile)

	bs, _ := os.ReadFile("./input.txt")
	lines := bytes.Split(bs, []byte("\n"))

	log.Println(part1(lines))
}

func part1(input [][]byte) int {
	result := 0

	for _, line := range input {
		result += score(string(line))
	}

	return result
}

// There is no longer a record of part 1... woops
func score(s string) int {
	score := 0

	for i := 0; i < len(s); i++ {
		if n, ok := isDigit(s[i]); ok {
			score += 10 * n
			break
		} else if n, ok := isNumWord(s[i:]); ok {
			score += 10 * n
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if n, ok := isDigit(s[i]); ok {
			score += n
			break
		} else if n, ok := isNumWord(s[i:]); ok {
			score += n
			break
		}
	}

	return score
}

func isDigit(b byte) (int, bool) {
	if b >= '0' && b <= '9' {
		digit, _ := strconv.Atoi(string(b))
		return digit, true
	} else {
		return -1, false
	}
}

var (
	words = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func isNumWord(s string) (int, bool) {
	for word, n := range words {
		if strings.HasPrefix(s, word) {
			return n, true
		}
	}

	return 0, false
}
