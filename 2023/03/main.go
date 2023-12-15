package main

import (
	"log"
	"strconv"

	"github.com/cnnrznn/advent-of-code/util"
)

func main() {
	lines, err := util.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines = lines[:len(lines)-1]
	log.Println(len(lines))

	var result int

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			switch {
			case lines[i][j] >= '0' && lines[i][j] <= '9':
				// number starts here
				var score int
				j, score = processNum(i, j, lines)
				j -= 1
				result += score
			default:
				// symbol
			}
		}
	}

	log.Printf("Result is (%v)\n", result)
}

func processNum(i, j int, mat []string) (int, int) {
	endJ := findLast(i, j, mat)
	var score int = 0

	lastJ := findLast(i, j, mat)

	n, _ := strconv.Atoi(mat[i][j:lastJ])
	if isTouching(i, j, endJ, mat) {
		score = n
	}

	return endJ, score
}

func findLast(i, j int, lines []string) int {
	for k := j; k < len(lines[i]); k++ {
		if lines[i][k] < '0' || lines[i][k] > '9' {
			return k
		}
	}
	return len(lines[i])
}

func isTouching(i, startJ, endJ int, lines []string) bool {
	for j := startJ; j < endJ; j++ {
		if above(i, j, lines) ||
			below(i, j, lines) ||
			(j == startJ && (left(i, j, lines) ||
				symAboveOrBelow(i, j-1, lines))) ||
			(j == endJ-1 && (right(i, j, lines) ||
				symAboveOrBelow(i, j+1, lines))) {
			return true
		}
	}
	return false
}

func symAboveOrBelow(i, j int, mat []string) bool {
	if j < 0 || j >= len(mat[i]) {
		return false
	}
	return above(i, j, mat) || below(i, j, mat)
}

func above(i, j int, mat []string) bool {
	return i > 0 && isSym(i-1, j, mat)
}

func below(i, j int, mat []string) bool {
	return i < len(mat)-1 && isSym(i+1, j, mat)
}

func left(i, j int, mat []string) bool {
	return j > 0 && isSym(i, j-1, mat)
}

func right(i, j int, mat []string) bool {
	return j < len(mat[i])-1 && isSym(i, j+1, mat)
}

func isSym(i, j int, mat []string) bool {
	c := mat[i][j]
	if (c < '0' || c > '9') && c != '.' {
		return true
	}
	return false
}
