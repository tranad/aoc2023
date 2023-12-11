package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
	"math"
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
	defer timer("Day 11")()
	fmt.Println(solution2(lines, 1))
	fmt.Println(solution2(lines, 1000000-1))
	return
}

func solution2(lines []string, dilation int) int {
	var galaxies []Coordinate
	for row, line := range lines {
		for col, _ := range line {
			coord := Coordinate{row, col}
			if string(lines[row][col]) == "#" {
				galaxies = append(galaxies, coord)
			}
		}
	}

	var filledRows []int
	var filledCols []int
	for _, galaxy := range galaxies {
		filledRows = append(filledRows, galaxy.row)
		filledCols = append(filledCols, galaxy.col)
	}
	var emptyRows []int
	var emptyCols []int
	for r := 0; r < len(lines); r++ {
		if !intIsIn(r, filledRows) {
			emptyRows = append(emptyRows, r)
		}
	}
	for c := 0; c < len(lines[0]); c++ {
		if !intIsIn(c, filledCols) {
			emptyCols = append(emptyCols, c)
		}
	}

	distSum := 0
	for i, _ := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			gi := galaxies[i]
			gj := galaxies[j]

			extraRows := 0
			brow, arow := gi.row, gj.row
			if gi.row < gj.row {
				arow, brow = gi.row, gj.row
			}
			for _, er := range emptyRows {
				if arow < er && er < brow {
					extraRows++
				}
			}

			extraCols := 0
			bcol, acol := gi.col, gj.col
			if gi.col < gj.col {
				acol, bcol = gi.col, gj.col
			}
			for _, er := range emptyCols {
				if acol < er && er < bcol {
					extraCols++
				}
			}

			x := float64(brow - arow + extraRows*dilation)
			y := float64(bcol - acol + extraCols*dilation)
			distSum += int(math.Abs(x) + math.Abs(y))
		}
	}
	return distSum
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

func (a Coordinate) DistanceTo(b Coordinate) int {
	// Manhattan distance
	return int(math.Abs(float64((b.row - a.row))) + math.Abs(float64((b.col - a.col))))
}

// How are these not a standard functions iswtfg
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n%s took %v\n", name, time.Since(start))
	}
}

func stringIsIn(target string, arr []string) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}

func intIsIn(target int, arr []int) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}
