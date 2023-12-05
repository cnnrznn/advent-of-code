package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	copies = make(map[int]int)
)

func numberOfCopies(x int) int {
	if n, ok := copies[x]; ok {
		return n
	}
	return 1
}

func incrementCopies(x int, by int) {
	copies[x] = numberOfCopies(x) + by
}

func main() {
	result := 0 // now it's number of cards

	log.SetFlags(log.Lshortfile)

	bs, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(bs, []byte("\n"))

	for i, line := range lines {
		result += score(i, string(line))
	}

	log.Printf("Result score is %v\n", result)
}

func score(i int, line string) int {
	score := numberOfCopies(i)

	winners := make(map[int]struct{})

	if len(strings.Fields(line)) == 0 {
		return 0
	}

	line = strings.Split(line, ":")[1]
	scores := strings.Split(line, "|")
	winnerStr := strings.Fields(scores[0])
	haveStr := strings.Fields(scores[1])

	for _, field := range winnerStr {
		n, _ := strconv.Atoi(field)
		winners[n] = struct{}{}
	}

	j := 1
	for _, field := range haveStr {
		n, _ := strconv.Atoi(field)
		if _, ok := winners[n]; ok {
			incrementCopies(i+j, score)
			j++
		}
	}

	return score
}
