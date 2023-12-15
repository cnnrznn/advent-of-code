package main

import (
	"log"
	"strconv"

	"github.com/cnnrznn/advent-of-code/util"
)

func main() {
	lines, err := util.ReadLines("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines = lines[:len(lines)-1]

	var result int

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '*' {
				result += processSym(i, j, lines)
			}
		}
	}

	log.Printf("Result is (%v)\n", result)
}

func processSym(i, j int, mat []string) int {
	ns := adjacentNums(i, j, mat)
	if len(ns) == 2 {
		return ns[0] * ns[1]
	}
	return 0
}

func adjacentNums(i, j int, mat []string) []int {
	var result []int

	if n, ok := numAt(i, j-1, mat); ok { // left
		result = append(result, n)
	}
	if n, ok := numAt(i, j+1, mat); ok { // right
		result = append(result, n)
	}

	if n, ok := numAt(i-1, j, mat); ok { // above
		result = append(result, n)
	} else {
		if n, ok := numAt(i-1, j-1, mat); ok { // above left
			result = append(result, n)
		}
		if n, ok := numAt(i-1, j+1, mat); ok { // above right
			result = append(result, n)
		}
	}

	if n, ok := numAt(i+1, j, mat); ok { // below
		result = append(result, n)
	} else {
		if n, ok := numAt(i+1, j-1, mat); ok { // below left
			result = append(result, n)
		}
		if n, ok := numAt(i+1, j+1, mat); ok { // below right
			result = append(result, n)
		}
	}

	return result
}

func numAt(i, j int, mat []string) (int, bool) {
	if i < 0 || i >= len(mat) || j < 0 || j >= len(mat[i]) {
		return 0, false
	}

	if mat[i][j] < '0' || mat[i][j] > '9' {
		return 0, false
	}

	return expand(i, j, mat), true
}

func expand(i, j int, mat []string) int {
	startJ, endJ := j, j

	for startJ > 0 && mat[i][startJ-1] >= '0' && mat[i][startJ-1] <= '9' {
		startJ--
	}

	for endJ < len(mat[i]) && mat[i][endJ] >= '0' && mat[i][endJ] <= '9' {
		endJ++
	}

	n, _ := strconv.Atoi(mat[i][startJ:endJ])

	return n
}
