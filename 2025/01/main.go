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
	lines, err := util.ReadStdio()
	if err != nil {
		panic(err)
	}

	// curr := puzzleMidpoint
	var (
		spinner = NewSpinner(puzzleMidpoint, puzzleRange)
		result  int
	)

	for _, line := range lines {
		originalLine := line
		dir := 1
		if line[0] == 'L' {
			dir = -1
		}

		line = line[1:]

		x, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		spinner.Spin(dir * x)

		if spinner.Position() == 0 {
			result++
		}

		fmt.Println(originalLine, x, dir, spinner.Position())
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

type Spinner struct {
	curr int
	size int
}

func NewSpinner(start, size int) *Spinner {
	return &Spinner{
		curr: start,
		size: size,
	}
}

func (s *Spinner) Spin(delta int) {
	new := s.curr + delta

	// correct large deltas
	for new > s.size-1 {
		new = new - s.size
	}

	// correct large deltas
	for new < 0 {
		new = new + s.size
	}
	s.curr = new
}

func (s *Spinner) Position() int {
	return s.curr
}
