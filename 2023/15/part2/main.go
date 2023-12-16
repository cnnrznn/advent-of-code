// --- Part Two ---
// You convince the reindeer to bring you the page; the page confirms that your HASH algorithm is working.
//
// The book goes on to describe a series of 256 boxes numbered 0 through 255. The boxes are arranged in a line starting from the point where light enters the facility. The boxes have holes that allow light to pass from one box to the next all the way down the line.
//
//	+-----+  +-----+         +-----+
//
// Light | Box |  | Box |   ...   | Box |
// ----------------------------------------->
//
//	|  0  |  |  1  |   ...   | 255 |
//	+-----+  +-----+         +-----+
//
// Inside each box, there are several lens slots that will keep a lens correctly positioned to focus light passing through the box. The side of each box has a panel that opens to allow you to insert or remove lenses as necessary.
//
// Along the wall running parallel to the boxes is a large library containing lenses organized by focal length ranging from 1 through 9. The reindeer also brings you a small handheld label printer.
//
// The book goes on to explain how to perform each step in the initialization sequence, a process it calls the Holiday ASCII String Helper Manual Arrangement Procedure, or HASHMAP for short.
//
// Each step begins with a sequence of letters that indicate the label of the lens on which the step operates. The result of running the HASH algorithm on the label indicates the correct box for that step.
//
// The label will be immediately followed by a character that indicates the operation to perform: either an equals sign (=) or a dash (-).
//
// If the operation character is a dash (-), go to the relevant box and remove the lens with the given label if it is present in the box. Then, move any remaining lenses as far forward in the box as they can go without changing their order, filling any space made by removing the indicated lens. (If no lens in that box has the given label, nothing happens.)
//
// If the operation character is an equals sign (=), it will be followed by a number indicating the focal length of the lens that needs to go into the relevant box; be sure to use the label maker to mark the lens with the label given in the beginning of the step so you can find it later. There are two possible situations:
//
// If there is already a lens in the box with the same label, replace the old lens with the new lens: remove the old lens and put the new lens in its place, not moving any other lenses in the box.
// If there is not already a lens in the box with the same label, add the lens to the box immediately behind any lenses already in the box. Don't move any of the other lenses when you do this. If there aren't any lenses in the box, the new lens goes all the way to the front of the box.
// Here is the contents of every box after each step in the example initialization sequence above:
//
// After "rn=1":
// Box 0: [rn 1]
//
// After "cm-":
// Box 0: [rn 1]
//
// After "qp=3":
// Box 0: [rn 1]
// Box 1: [qp 3]
//
// After "cm=2":
// Box 0: [rn 1] [cm 2]
// Box 1: [qp 3]
//
// After "qp-":
// Box 0: [rn 1] [cm 2]
//
// After "pc=4":
// Box 0: [rn 1] [cm 2]
// Box 3: [pc 4]
//
// After "ot=9":
// Box 0: [rn 1] [cm 2]
// Box 3: [pc 4] [ot 9]
//
// After "ab=5":
// Box 0: [rn 1] [cm 2]
// Box 3: [pc 4] [ot 9] [ab 5]
//
// After "pc-":
// Box 0: [rn 1] [cm 2]
// Box 3: [ot 9] [ab 5]
//
// After "pc=6":
// Box 0: [rn 1] [cm 2]
// Box 3: [ot 9] [ab 5] [pc 6]
//
// After "ot=7":
// Box 0: [rn 1] [cm 2]
// Box 3: [ot 7] [ab 5] [pc 6]
// All 256 boxes are always present; only the boxes that contain any lenses are shown here. Within each box, lenses are listed from front to back; each lens is shown as its label and focal length in square brackets.
//
// To confirm that all of the lenses are installed correctly, add up the focusing power of all of the lenses. The focusing power of a single lens is the result of multiplying together:
//
// One plus the box number of the lens in question.
// The slot number of the lens within the box: 1 for the first lens, 2 for the second lens, and so on.
// The focal length of the lens.
// At the end of the above example, the focusing power of each lens is as follows:
//
// rn: 1 (box 0) * 1 (first slot) * 1 (focal length) = 1
// cm: 1 (box 0) * 2 (second slot) * 2 (focal length) = 4
// ot: 4 (box 3) * 1 (first slot) * 7 (focal length) = 28
// ab: 4 (box 3) * 2 (second slot) * 5 (focal length) = 40
// pc: 4 (box 3) * 3 (third slot) * 6 (focal length) = 72
// So, the above example ends up with a total focusing power of 145.
//
// With the help of an over-enthusiastic reindeer in a hard hat, follow the initialization sequence. What is the focusing power of the resulting lens configuration?
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cnnrznn/advent-of-code/util"
)

type Box struct {
	Arr     []Lens
	Indices map[string]int
}

type Lens struct {
	Label string
	Power int
}

func main() {
	lines, _ := util.ReadLines("../input.txt")
	boxes := map[int]*Box{}

	steps := []string{}

	for _, line := range lines {
		steps = append(steps, strings.Split(line, ",")...)
	}

	for _, step := range steps {
		do(step, boxes)
	}

	fmt.Println(score(boxes))
}

func do(step string, boxes map[int]*Box) {
	label := ""
	focal := -1
	cmd := -1

	switch {
	case step[len(step)-1] == '-':
		cmd = 0
		label = step[:len(step)-1]
	default:
		cmd = 1
		parts := strings.Split(step, "=")
		label = parts[0]
		focal, _ = strconv.Atoi(parts[1])
	}

	boxN := hash(label)
	if _, ok := boxes[boxN]; !ok {
		boxes[boxN] = &Box{
			Indices: make(map[string]int),
		}
	}

	if cmd > 0 {
		addLens(label, focal, boxes[boxN])
	} else {
		removeLens(label, boxes[boxN])
	}

}

func addLens(label string, focal int, box *Box) {
	lens := Lens{
		Label: label,
		Power: focal,
	}

	if i, ok := box.Indices[label]; ok {
		box.Arr[i] = lens
	} else {
		box.Indices[label] = len(box.Arr)
		box.Arr = append(box.Arr, lens)
	}

	//log.Println("add", label, box.Arr, box.Indices)
}

func removeLens(label string, box *Box) {
	if i, ok := box.Indices[label]; ok {
		//log.Println("remove", i, label, box.Indices, box.Arr)
		delete(box.Indices, label)
		for j := i; j < len(box.Arr)-1; j++ {
			box.Arr[j] = box.Arr[j+1]
			box.Indices[box.Arr[j].Label]--
		}
		if len(box.Arr) > 0 {
			box.Arr = box.Arr[:len(box.Arr)-1]
		}
	}
}

func hash(s string) int {
	result := 0

	for _, r := range s {
		result += int(r)
		result = result * 17
		result = result % 256
	}

	return result
}

func score(boxes map[int]*Box) int {
	result := 0

	for boxN, box := range boxes {
		for slot, lens := range box.Arr {
			result += (boxN + 1) * (slot + 1) * lens.Power
		}
	}

	return result
}
