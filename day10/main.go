package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// MAIN
	defer timer("Day 10")()
	pipeMap := make(map[Coordinate]Pipe)
	startPipe := Pipe{Coordinate{-1, -1}, "."}
	for row, line := range lines {
		for col, shape := range line {
			coord := Coordinate{row, col}
			pipe := Pipe{coord, string(shape)}
			pipeMap[coord] = pipe
			if pipe.shape == "S" {
				startPipe = pipe
			}
		}
	}

	// Part 1
	// Get the loop as a slice of Pipes{coord, shape}
	// ans1 = len(loop)/2
	var history []Pipe
	loop := findLoop(startPipe, &pipeMap, &history)
	ans1 := len(loop) / 2

	// Part 2
	// Scan row by row across the input.
	// Start with parity = 0
	// As we go left to right, if we encounter a coord in the loop
	// parity = parity+1 % 2
	// for every coord not in loop at parity =1, add to area
	ans2 := countArea(loop, &pipeMap, lines)

	fmt.Printf("Part 1: %v\n", ans1)
	fmt.Printf("Part 2: %v\n", ans2)
	// for row, line := range lines {
	// fmt.Printf("row %v: ", row)
	// for col := 0; col < len(line); col++{
	// thisCoord := Coordinate{row, col}
	// _, thisInLoop := loopMap[thisCoord]
	// if thisInLoop {
	// fmt.Printf("%v", "X")
	// } else {
	// fmt.Printf("%v", string(line[col]))
	// }
	// }
	// fmt.Println()
	// }
	return
}

func countArea(loop []Pipe, pipeMap *map[Coordinate]Pipe, lines []string) (area int) {
	loopMap := make(map[Coordinate]string)
	// For quickly checking if a coord is in the loop
	for _, pipe := range loop {
		loopMap[pipe.coord] = pipe.shape
	}
	for row, line := range lines {
		inside := false
		// fmt.Println(line)
		for col, _ := range line {
			_, partOfLoop := loopMap[Coordinate{row, col}]
			if partOfLoop {
				fmt.Print("X")
			} else {
				fmt.Print(string(line[col]))
			}
		}
		// fmt.Println()
		for col, _ := range line {
			shape, partOfLoop := loopMap[Coordinate{row, col}]
			// If not part of loop and inside, area=+
			// If part of loop and shape is |,L,J flip parity
			// | obvioulsy a wall, nothing to the right of it
			// J will have nothing to the right
			// 7 will have nothing to the right
			// L will have -,J,7 to the right
			// F will have -,J,7 to the right

			// Only change on
			// | (entering a wall)
			// L (entering a wall) (positive area) L------Jxxxx
			// J (exiting a wall) (positive area)  L------7xxxx
			// Do not change on
			// F (entering a wall) (negative area) F------Jxxx
			// 7 (exiting a wall) (negative area)  F------7yyy
			// will get the above on next line
			// turns out S = J for me
			if shape == "S" {
				sPipe := findS(loop[0], loop[1], loop[len(loop)-1])
				shape = sPipe.shape
			}
			if !partOfLoop {
				if inside {
					area++
				}
			} else if shape == "|" || shape == "J" || shape == "L" {
				inside = !inside
			}
		}
		fmt.Println()
	}
	return area
}

func findS(start Pipe, left Pipe, right Pipe) Pipe {
	shapes := [6]string{"-", "|", "J", "7", "F", "L"}
	for _, shape := range shapes {
		tmpPipe := Pipe{start.coord, shape}
		r := tmpPipe.Joins(right)
		l := tmpPipe.Joins(left)
		if l && r {
			return tmpPipe
		}
	}
	return start
}

func findLoop(start Pipe, pipeMap *map[Coordinate]Pipe, history *[]Pipe) []Pipe {
	*history = append(*history, start)
	nbhrs := start.coord.Neighbors()
	var nextPipes []Pipe
	for _, coord := range nbhrs {
		// check it's within bounds
		dirPipe, valid := (*pipeMap)[coord]
		if valid == true {
			// check it's a valid joining
			if start.Joins(dirPipe) && !dirPipe.In(*history) {
				nextPipes = append(nextPipes, dirPipe)
			}
		}
	}
	for _, nextPipe := range nextPipes {
		if nextPipe.shape == "S" {
			// fmt.Printf("DONE %v", history)
			// panic("Done")
			return *history
		} else {
			return findLoop(nextPipe, pipeMap, history)
		}
	}
	return *history
}

type Coordinate struct {
	row int
	col int
}

func (c Coordinate) Neighbors() map[string]Coordinate {
	north := Coordinate{c.row - 1, c.col}
	south := Coordinate{c.row + 1, c.col}
	east := Coordinate{c.row, c.col - 1}
	west := Coordinate{c.row, c.col + 1}
	nbhr := make(map[string]Coordinate)
	nbhr["west"] = west
	nbhr["east"] = east
	nbhr["north"] = north
	nbhr["south"] = south
	return nbhr
}

type Pipe struct {
	coord Coordinate
	shape string
}

func (target Pipe) In(arr []Pipe) bool {
	for _, element := range arr {
		if element.coord.row == target.coord.row && element.coord.col == target.coord.col && element.shape == target.shape {
			return true
		}
	}
	return false
}

func (p Pipe) Joins(q Pipe) (valid bool) {
	delta := Coordinate{q.coord.row - p.coord.row, q.coord.col - p.coord.col}
	shapeMatchMap := make(map[string][]string)
	switch {
	// north
	case delta.row == -1 && delta.col == 0:
		shapeMatchMap["|"] = []string{"|", "7", "F"}
		shapeMatchMap["J"] = []string{"|", "7", "F"}
		shapeMatchMap["L"] = []string{"|", "7", "F"}
		shapeMatchMap["S"] = []string{"|", "7", "F"}
	// south
	case delta.row == 1 && delta.col == 0:
		shapeMatchMap["|"] = []string{"|", "J", "L"}
		shapeMatchMap["7"] = []string{"|", "J", "L"}
		shapeMatchMap["F"] = []string{"|", "J", "L"}
		shapeMatchMap["S"] = []string{"|", "J", "L"}
	// west
	case delta.row == 0 && delta.col == -1:
		shapeMatchMap["-"] = []string{"-", "F", "L"}
		shapeMatchMap["J"] = []string{"-", "F", "L"}
		shapeMatchMap["7"] = []string{"-", "F", "L"}
		shapeMatchMap["S"] = []string{"-", "F", "L"}
	// east
	case delta.row == 0 && delta.col == 1:
		shapeMatchMap["-"] = []string{"-", "J", "7"}
		shapeMatchMap["L"] = []string{"-", "J", "7"}
		shapeMatchMap["F"] = []string{"-", "J", "7"}
		shapeMatchMap["S"] = []string{"-", "J", "7"}
	}
	return isIn(q.shape, shapeMatchMap[p.shape])
}

// How are these not a standard functions iswtfg
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n%s took %v\n", name, time.Since(start))
	}
}

func isIn(target string, arr []string) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}
