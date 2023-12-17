package main

import (
	"fmt"

	"github.com/cnnrznn/advent-of-code/util"
)

type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
)

type Cell struct {
	Sym  rune
	Dirs map[Dir]struct{} // outgoing directions
}

func (c *Cell) Energized() bool {
	return len(c.Dirs) > 0
}

/* part 1
func main() {
	lines, _ := util.ReadLines("./input.txt")
	mat := makeByteMat(lines)

	bounce(mat, 0, 0, RIGHT)

	fmt.Printf("Energized: %v\n", numEnergized(mat))
}*/

func main() {
	lines, _ := util.ReadLines("./input.txt")
	var mat [][]*Cell
	var max, score int

	// top + bot
	for j := 0; j < len(lines[0]); j++ {
		mat = makeByteMat(lines)
		bounce(mat, 0, j, DOWN)
		score = numEnergized(mat)
		if score > max {
			max = score
		}

		mat = makeByteMat(lines)
		bounce(mat, len(mat)-1, j, UP)
		score = numEnergized(mat)
		if score > max {
			max = score
		}
	}

	// left + right
	for i := 0; i < len(lines); i++ {
		mat = makeByteMat(lines)
		bounce(mat, i, 0, RIGHT)
		score = numEnergized(mat)
		if score > max {
			max = score
		}

		mat = makeByteMat(lines)
		bounce(mat, i, len(mat[0])-1, LEFT)
		score = numEnergized(mat)
		if score > max {
			max = score
		}
	}

	fmt.Println(max)
}

func bounce(mat [][]*Cell, i, j int, d Dir) {
	if i < 0 || i >= len(mat) || j < 0 || j >= len(mat[0]) {
		return
	}

	cell := mat[i][j]
	if _, ok := cell.Dirs[d]; ok {
		return
	}

	cell.Dirs[d] = struct{}{}

	switch d {
	case UP:
		switch mat[i][j].Sym {
		case '.', '|':
			bounce(mat, i-1, j, UP)
		case '\\':
			bounce(mat, i, j-1, LEFT)
		case '/':
			bounce(mat, i, j+1, RIGHT)
		case '-':
			bounce(mat, i, j-1, LEFT)
			bounce(mat, i, j+1, RIGHT)
		}
	case DOWN:
		switch mat[i][j].Sym {
		case '.', '|':
			bounce(mat, i+1, j, DOWN)
		case '\\':
			bounce(mat, i, j+1, RIGHT)
		case '/':
			bounce(mat, i, j-1, LEFT)
		case '-':
			bounce(mat, i, j-1, LEFT)
			bounce(mat, i, j+1, RIGHT)
		}
	case LEFT:
		switch mat[i][j].Sym {
		case '.', '-':
			bounce(mat, i, j-1, LEFT)
		case '\\':
			bounce(mat, i-1, j, UP)
		case '/':
			bounce(mat, i+1, j, DOWN)
		case '|':
			bounce(mat, i-1, j, UP)
			bounce(mat, i+1, j, DOWN)
		}
	case RIGHT:
		switch mat[i][j].Sym {
		case '.', '-':
			bounce(mat, i, j+1, RIGHT)
		case '\\':
			bounce(mat, i+1, j, DOWN)
		case '/':
			bounce(mat, i-1, j, UP)
		case '|':
			bounce(mat, i-1, j, UP)
			bounce(mat, i+1, j, DOWN)
		}
	}
}

func makeByteMat(ls []string) [][]*Cell {
	mat := make([][]*Cell, len(ls))
	for i, line := range ls {
		mat[i] = make([]*Cell, len(line))
		for j, b := range line {
			mat[i][j] = &Cell{
				Sym:  b,
				Dirs: make(map[Dir]struct{}),
			}
		}
	}

	return mat
}

func numEnergized(mat [][]*Cell) int {
	result := 0

	for _, row := range mat {
		for _, cell := range row {
			if cell.Energized() {
				result++
			}
		}
	}

	return result
}
