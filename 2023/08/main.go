package main

import (
	"fmt"
	"regexp"

	"github.com/cnnrznn/advent-of-code/util"
)

/* part 1
func main() {
	lines, _ := util.ReadLines("./input.txt")
	steps := parseSteps(lines[0])
	network := parseNodes(lines[1:])

	curr := "AAA"
	numSteps := 0

	for curr != "ZZZ" {
		numSteps++
		switch steps.Next() {
		case LEFT:
			curr = network[curr][0]
		case RIGHT:
			curr = network[curr][1]
		}
	}

	fmt.Println(numSteps)
}*/

/* A cool idea I had for brute-forcing the solution
func RunGhost(start string, network Network, ins Instructions, output chan int) {
	curr := start
	steps := 0
	for {
		steps++
		curr = network[curr][ins.Next()]
		if curr[2] == 'Z' {
			output <- steps
		}
	}
}*/

func main() {
	lines, _ := util.ReadLines("./input.txt")
	network, starts := parseNodesPart2(lines[1:])
	ins := parseSteps(lines[0])

	//log.Println(ins)
	//log.Println(starts)
	//log.Println(network)

	frequencies := []int{}

	for _, start := range starts {
		curr := start
		steps := 0
		ins.curr = 0
		for curr[2] != 'Z' {
			steps++
			dir := ins.Next()
			curr = network[curr][dir]
		}

		frequencies = append(frequencies, steps)
	}

	fmt.Println(frequencies)

	fmt.Println(LCM(frequencies[0], frequencies[1], frequencies[2:]...))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func allZs(nodes []string) bool {
	for _, node := range nodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func parseSteps(s string) Instructions {
	dirs := []Dir{}
	for _, r := range s {
		if r == 'L' {
			dirs = append(dirs, LEFT)
		} else {
			dirs = append(dirs, RIGHT)
		}
	}

	return Instructions{
		Dirs: dirs,
	}
}

func parseNodesPart2(lines []string) (Network, []string) {
	net := make(Network)
	starts := []string{}

	re, _ := regexp.Compile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)
	for _, line := range lines {
		match := re.FindSubmatch([]byte(line))
		net[string(match[1])] = []string{
			string(match[2]),
			string(match[3]),
		}

		if match[1][2] == 'A' {
			starts = append(starts, string(match[1]))
		}
	}

	return net, starts
}

func parseNodes(lines []string) Network {
	net := make(Network)

	re, _ := regexp.Compile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)
	for _, line := range lines {
		match := re.FindSubmatch([]byte(line))
		net[string(match[1])] = []string{
			string(match[2]),
			string(match[3]),
		}
	}

	return net
}

type Network map[string][]string

type Instructions struct {
	Dirs []Dir
	curr int
}

func (i *Instructions) Next() Dir {
	result := i.Dirs[i.curr]
	i.curr = (i.curr + 1) % len(i.Dirs)
	return result
}

type Dir int

const (
	LEFT  = 0
	RIGHT = 1
)
