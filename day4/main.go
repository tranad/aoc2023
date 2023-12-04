package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	// "slices"
)

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
type Card struct {
	ID             int
	WinningNumbers []int
	MyNumbers      []int
}

func containsElement(arr []int, target int) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}

func (c Card) Wins() int {
	count := 0
	for _, w := range c.WinningNumbers {
		if containsElement(c.MyNumbers, w) {
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return count
}

func (c Card) Points() int {
	wins := c.Wins()
	return int(math.Pow(2.0, float64(wins)-1.0))
}

func parseInput(line string) Card {
	card_id_str := strings.Split(strings.Split(line, ":")[0], " ")
	card_id, _ := strconv.Atoi(string(card_id_str[len(card_id_str)-1]))

	numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
	wn := strings.Split(numbers[0], " ")
	mn := strings.Split(numbers[1], " ")

	var winning_numbers []int
	var my_numbers []int
	for _, v := range wn {
		w, err := strconv.Atoi(string(v))
		if err == nil {
			winning_numbers = append(winning_numbers, w)
		}
	}
	for _, v := range mn {
		m, err := strconv.Atoi(string(v))
		if err == nil {
			my_numbers = append(my_numbers, m)
		}
	}
	return Card{card_id, winning_numbers, my_numbers}
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice to store the lines
	var lines []string

	// Iterate over each line and append it to the slice
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	ans1 := 0
	ans2 := 0

	copies := make(map[int]int)
	for k, _ := range lines {
		copies[k+1] = 1
	}
	for _, line := range lines {
		card := parseInput(line)
		points := card.Points()
		ans1 += points

		wins := card.Wins()
		for k := 1; k <= wins; k++ {
			copies[card.ID+k] += 1 * copies[card.ID]
		}
	}

	for _, v := range copies {
		ans2 += v
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
	return
}
