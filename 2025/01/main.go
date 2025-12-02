package main

import (
	"fmt"
	"strconv"

	"github.com/cnnrznn/advent-of-code/util"
)

const (
	puzzleRange    = 100
	puzzleMidpoint = 50
)

func main() {
	lines, err := util.ReadInput()
	if err != nil {
		panic(err)
	}

	var (
		spinner = util.NewSpinner(puzzleMidpoint, puzzleRange)
		result  int
	)

	for _, line := range lines {
		dir := 1
		if line[0] == 'L' {
			dir = -1
		}

		line = line[1:]

		x, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		result += spinner.Spin(dir * x)
	}

	fmt.Println(result)
}



type ProcessorFunc[I, O any] func(I) (O, error)

type Processor[I, O any] struct {
	proc ProcessorFunc[I, O]
}

func NewProcessor[I, O any](f ProcessorFunc[I, O]) *Processor[I, O] {
	return &Processor[I, O]{
		proc: f,
	}
}

func (p *Processor[I, O]) Exec(i I) O {
	o, err := p.proc(i)
	if err != nil {
		panic(err)
	}
	return o
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
