package main

import (
	// "bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	// "slices"
)

type pair struct {
	left  int
	right int
}

func main() {
	// boiler plate load input
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(f), "\n")
	// ans1 := solution1(lines)
	// fmt.Println(ans1)

	ans2 := solution2(lines)
	fmt.Println(ans2)
}

func solution2(lines []string) int {
	// For each number, get area and location of symbols
	// store each * symbol and have an array of numbers
	// go through each * loc and get gear ratio

	// problem parameters
	gear_map := make(map[pair][]int)

	ans := 0
	for pos, line := range lines {
		number_bounds, _ := get_number_indices(line)
		for _, number_bound := range number_bounds {
			left := number_bound.left
			right := number_bound.right
			number, _ := strconv.Atoi(line[left:right])

			horizontal_start := int(math.Max(float64(left-1), 0.0))
			horizontal_end := int(math.Min(float64(right+1), float64(len(line))))
			vertical_start := int(math.Max(float64(pos-1), 0.0))
			vertical_end := int(math.Min(float64(pos+1+1), float64(len(lines)-1)))

			for row := vertical_start; row < vertical_end; row++ {
				for col := horizontal_start; col < horizontal_end; col++ {
					symbol := string(lines[row][col])
					if symbol == "*" {
						gear := pair{row, col}
						_, is_gear := gear_map[gear]
						if is_gear {
							gear_map[gear] = append(gear_map[gear], number)
						} else {
							gear_list := []int{number}
							gear_map[gear] = gear_list
						}
					}
				}
			}
		}
	}

	for _, v := range gear_map {
		if len(v) == 2 {
			ans += v[0] * v[1]
		}
	}
	return ans
}

func solution1(lines []string) int {
	// problem parameters
	symbol_map := map[string]string{
		"!":  "",
		"@":  "",
		"#":  "",
		"$":  "",
		"%":  "",
		"^":  "",
		"&":  "",
		"*":  "",
		"(":  "",
		")":  "",
		"-":  "",
		"_":  "",
		"+":  "",
		"=":  "",
		"\\": "",
		"/":  "",
		"'":  "",
		"?":  "",
		">":  "",
		"<":  "",
		"`":  "",
		"~":  "",
		"|":  "",
	}

	ans := 0
	for pos, line := range lines {

		number_bounds, _ := get_number_indices(line)
		for _, number_bound := range number_bounds {
			left := number_bound.left
			right := number_bound.right
			number, _ := strconv.Atoi(line[left:right])

			horizontal_start := int(math.Max(float64(left-1), 0.0))
			horizontal_end := int(math.Min(float64(right+1), float64(len(line))))
			vertical_start := int(math.Max(float64(pos-1), 0.0))
			vertical_end := int(math.Min(float64(pos+1+1), float64(len(lines)-1)))

			area := ""
			for row := vertical_start; row < vertical_end; row++ {
				area += lines[row][horizontal_start:horizontal_end]
			}
			for _, entry := range area {
				_, is_symbol := symbol_map[string(entry)]
				if is_symbol == true {
					fmt.Println(number)
					ans += number
				}
			}
		}
	}
	return ans
}

func get_number_indices(line string) ([]pair, []int) {
	var number_bounds []int

	prev_parity := 1
	cur_parity := 1

	for pos, dig := range line {
		_, err := strconv.Atoi(string(dig))
		if err == nil {
			cur_parity = 0
		} else {
			cur_parity = 1
		}
		if cur_parity != prev_parity {
			number_bounds = append(number_bounds, pos)
		}
		if pos == len(line)-1 && cur_parity == 0 {
			number_bounds = append(number_bounds, pos+1)
		}
		prev_parity = cur_parity
	}

	var paired_nb []pair
	for pos := 0; pos < len(number_bounds); pos += 2 {
		if pos == len(number_bounds)-1 {
			break
		}
		bnd := pair{number_bounds[pos], number_bounds[pos+1]}

		paired_nb = append(paired_nb, bnd)

	}
	var numbers []int
	for _, bnd := range paired_nb {
		num, _ := strconv.Atoi(line[bnd.left:bnd.right])
		numbers = append(numbers, num)
	}
	return paired_nb, numbers
}
