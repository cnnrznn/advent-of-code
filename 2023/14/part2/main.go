// --- Part Two ---
// The parabolic reflector dish deforms, but not in a way that focuses the beam. To do that, you'll need to move the rocks to the edges of the platform. Fortunately, a button on the side of the control panel labeled "spin cycle" attempts to do just that!
//
// Each cycle tilts the platform four times so that the rounded rocks roll north, then west, then south, then east. After each tilt, the rounded rocks roll as far as they can before the platform tilts in the next direction. After one cycle, the platform will have finished rolling the rounded rocks in those four directions in that order.
//
// Here's what happens in the example above after each of the first few cycles:
//
// After 1 cycle:
// .....#....
// ....#...O#
// ...OO##...
// .OO#......
// .....OOO#.
// .O#...O#.#
// ....O#....
// ......OOOO
// #...O###..
// #..OO#....
//
// After 2 cycles:
// .....#....
// ....#...O#
// .....##...
// ..O#......
// .....OOO#.
// .O#...O#.#
// ....O#...O
// .......OOO
// #..OO###..
// #.OOO#...O
//
// After 3 cycles:
// .....#....
// ....#...O#
// .....##...
// ..O#......
// .....OOO#.
// .O#...O#.#
// ....O#...O
// .......OOO
// #...O###.O
// #.OOO#...O
// This process should work if you leave it running long enough, but you're still worried about the north support beams. To make sure they'll survive for a while, you need to calculate the total load on the north support beams after 1000000000 cycles.
//
// In the above example, after 1000000000 cycles, the total load on the north support beams is 64.
//
// Run the spin cycle for 1000000000 cycles. Afterward, what is the total load on the north support beams?
package main

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/cnnrznn/advent-of-code/util"
)

func main() {
	lines, _ := util.ReadLines("../input.txt")
	mat := makeByteMap(lines)
	snapshots := make(map[string]int)
	found := false

	for i := 0; i < 1000000000; i++ {
		for j := 0; j < 4; j++ {
			tilt(mat)
			mat = rotate(mat)
		}

		if !found {
			h := snapshot(mat)
			if j, ok := snapshots[h]; ok {
				cycleLen := i - j
				cycles := (1000000000 - i) / cycleLen
				i = i + (cycles * cycleLen)
				found = true
			} else {
				snapshots[h] = i
			}
		}
	}

	log.Println(load(mat))
}

func load(mat [][]byte) int {
	rows := len(mat)
	result := 0

	for i, row := range mat {
		for _, b := range row {
			switch b {
			case 'O':
				result += rows - i
			}
		}
	}

	return result
}

func tilt(mat [][]byte) {
	rollPos := make(map[int]int)

	for i := range mat {
		for j := range mat[i] {
			switch mat[i][j] {
			case 'O':
				mat[i][j] = '.'
				mat[rollPos[j]][j] = 'O'
				rollPos[j]++
			case '#':
				rollPos[j] = i + 1
			}
		}
	}
}

func rotate(mat [][]byte) [][]byte {
	nRows := len(mat)
	newMat := duplicate(mat)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			newMat[j][nRows-i-1] = mat[i][j]
		}
	}

	return newMat
}

func duplicate(mat [][]byte) [][]byte {
	result := make([][]byte, len(mat[0]))
	for i := range result {
		result[i] = make([]byte, len(mat))
	}
	return result
}

func makeByteMap(lines []string) [][]byte {
	mat := make([][]byte, len(lines))
	for i, row := range lines {
		mat[i] = make([]byte, len(row))
		for j := range row {
			mat[i][j] = row[j]
		}
	}
	return mat
}

func snapshot(mat [][]byte) string {
	h := sha256.New()
	for _, row := range mat {
		if _, err := h.Write(row); err != nil {
			log.Fatal(err)
		}
	}
	return string(h.Sum(nil))
}

func print(mat [][]byte) {
	for _, row := range mat {
		fmt.Println(string(row))
	}
	fmt.Println("===================")
}
