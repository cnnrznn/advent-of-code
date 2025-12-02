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

		result += spinner.Spin(dir * x)
		fmt.Println(originalLine, result)
	}

	fmt.Println(result)
}

type Spinner struct {
	curr int
	size int
	max  int
}

func NewSpinner(start, size int) *Spinner {
	return &Spinner{
		curr: start,
		size: size,
		max:  size - 1,
	}
}

// Spin moves the spinner by delta. It returns the number
// of times 0 was clicked.
func (s *Spinner) Spin(delta int) int {
	var (
		clicks int
	)

	for ; delta > 0; delta-- {
		s.curr++
		if s.curr > s.max {
			s.curr = 0
			clicks++
		}
	}

	for ; delta < 0; delta++ {
		s.curr--
		if s.curr < 0 {
			s.curr += s.size
		}
		if s.curr == 0 {
			clicks++
		}
	}

	return clicks
}

func (s *Spinner) Position() int {
	return s.curr
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
