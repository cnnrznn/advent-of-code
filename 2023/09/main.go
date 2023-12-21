package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/cnnrznn/advent-of-code/util"
)

func main() {
	ls, _ := util.ReadLines("./input.txt")

	answer := 0

	for _, line := range ls {
		answer += solve(line)
	}

	fmt.Println(answer)
}

func solve(line string) int {
	fields := strings.Fields(line)
	ns := []int{}
	for _, field := range fields {
		n, _ := strconv.Atoi(field)
		ns = append(ns, n)
	}

	// part 2
	slices.Reverse(ns)

	return do(ns)
}

func do(ns []int) int {
	//log.Println(ns)

	diffs, zeroes := calcDiffs(ns)

	if zeroes {
		return ns[len(ns)-1]
	}

	return ns[len(ns)-1] + do(diffs)
}

func calcDiffs(ns []int) ([]int, bool) {
	zeroes := true
	diffs := []int{}

	for i := 1; i < len(ns); i++ {
		diff := ns[i] - ns[i-1]
		diffs = append(diffs, diff)
		zeroes = zeroes && (diff == 0)
	}

	return diffs, zeroes
}
