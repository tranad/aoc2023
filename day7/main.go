package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	// "slices"
)

type CamelCardsHand struct {
	hand string
	bid  int
}

// Get type of hand for part 1
func (ch CamelCardsHand) Type1() (string, int) {
	handHist := make(map[string]int)
	for _, card := range ch.hand {
		handHist[string(card)] += 1
	}
	numUniqueCards := len(handHist)
	cardCountProduct := 1
	for _, v := range handHist {
		cardCountProduct *= v
	}
	switch {
	case numUniqueCards == 1:
		return "5Kind", 6
	case numUniqueCards == 2 && cardCountProduct == 4*1:
		return "4Kind", 5
	case numUniqueCards == 2 && cardCountProduct == 3*2:
		return "FullHouse", 4
	case numUniqueCards == 3 && cardCountProduct == 3*1*1:
		return "3Kind", 3
	case numUniqueCards == 3 && cardCountProduct == 2*2*1:
		return "2Pair", 2
	case numUniqueCards == 4 && cardCountProduct == 2*1*1*1:
		return "1Pair", 1
	case numUniqueCards == 5:
		return "HighCard", 0
	default:
		panic("Hand not matching any type")
	}
}

// Get type of hand for part 2
func (ch CamelCardsHand) Type2() (string, int) {
	handHist := make(map[string]int)
	for _, card := range ch.hand {
		handHist[string(card)] += 1
	}
	numUniqueCards := len(handHist)
	cardCountProduct := 1
	for _, v := range handHist {
		cardCountProduct *= v
	}
	// suppose J is an entry
	_, J_in_hist := handHist["J"]
	if J_in_hist {
		switch {
		case numUniqueCards == 1:
			return "5Kind", 6
		case numUniqueCards == 2 && cardCountProduct == 4*1:
			return "5Kind", 6
		case numUniqueCards == 2 && cardCountProduct == 3*2:
			return "5Kind", 6
		case numUniqueCards == 3 && cardCountProduct == 3*1*1:
			return "4Kind", 5
		case numUniqueCards == 3 && cardCountProduct == 2*2*1:
			if handHist["J"] == 2 {
				return "4Kind", 5
			} else {
				return "FullHouse", 4
			}
		case numUniqueCards == 4 && cardCountProduct == 2*1*1*1:
			return "3Kind", 3
		case numUniqueCards == 5:
			return "1Pair", 1
		default:
			panic("Hand not matching any type")
		}
	}
	return ch.Type1()
}

// Part 1 sorter
type Part1 []CamelCardsHand

func (cs Part1) Len() int {
	return len(cs)
}
func (cs Part1) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
func (cs Part1) Less(i, j int) bool {
	// You can customize the comparison logic based on your needs.
	// For example, compare based on IntegerValue.
	_, iweight := cs[i].Type1()
	_, jweight := cs[j].Type1()
	switch {
	case iweight < jweight:
		return true
	case iweight > jweight:
		return false
	case iweight == jweight:
		camelCardWeight := map[string]int{
			"A": 14,
			"K": 13,
			"Q": 12,
			"T": 10,
			"9": 9,
			"8": 8,
			"7": 7,
			"6": 6,
			"5": 5,
			"4": 4,
			"3": 3,
			"2": 2,
			"J": 11,
		}
		for k := 0; k < len(cs[0].hand); k++ {
			wik := camelCardWeight[string(cs[i].hand[k])]
			wjk := camelCardWeight[string(cs[j].hand[k])]
			if wik < wjk {
				return true
			} else if wik > wjk {
				return false
			} else {
				continue
			}
		}
		panic("Does not handle exact same hand")
	}
	panic("Does not handle exact same hand")
}

// Part 2 sorter
type Part2 []CamelCardsHand

func (cs Part2) Len() int {
	return len(cs)
}
func (cs Part2) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
func (cs Part2) Less(i, j int) bool {
	// You can customize the comparison logic based on your needs.
	// For example, compare based on IntegerValue.
	_, iweight := cs[i].Type2()
	_, jweight := cs[j].Type2()
	switch {
	case iweight < jweight:
		return true
	case iweight > jweight:
		return false
	case iweight == jweight:
		camelCardWeight := map[string]int{
			"A": 14,
			"K": 13,
			"Q": 12,
			"T": 10,
			"9": 9,
			"8": 8,
			"7": 7,
			"6": 6,
			"5": 5,
			"4": 4,
			"3": 3,
			"2": 2,
			"J": 1,
		}
		for k := 0; k < len(cs[0].hand); k++ {
			wik := camelCardWeight[string(cs[i].hand[k])]
			wjk := camelCardWeight[string(cs[j].hand[k])]
			if wik < wjk {
				return true
			} else if wik > wjk {
				return false
			} else {
				continue
			}
		}
		panic("Does not handle exact same hand")
	}
	panic("Does not handle exact same hand")
}

func parseLine(line string) CamelCardsHand {
	split := strings.Split(line, " ")
	// hand := []rune(split[0])
	hand := split[0]
	bid, _ := strconv.Atoi(split[1])
	cch := CamelCardsHand{hand, bid}
	return cch
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

	// MAIN
	var camelHands []CamelCardsHand
	for _, line := range lines {
		ch := parseLine(line)
		// fmt.Println(ch)
		// t,w := ch.Type()
		// fmt.Printf("\t{%v}{%v}\n", t, w)
		camelHands = append(camelHands, ch)
	}

	part1 := 0
	sort.Sort(Part1(camelHands))
	for i, ch := range camelHands {
		// t,w := ch.Type1()
		// fmt.Println(i, ch, t, w)
		part1 += (i + 1) * ch.bid
	}
	fmt.Printf("Part 1: %v\n", part1)

	part2 := 0
	sort.Sort(Part2(camelHands))
	for i, ch := range camelHands {
		// t,w := ch.Type()
		// fmt.Println(i, ch, t, w)
		part2 += (i + 1) * ch.bid
	}
	fmt.Printf("Part 2: %v\n", part2)
	return
}

////////////////////////////////
// Previous utility functions //
////////////////////////////////

// How are these not a standard functions iswtfg
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
func findMin(arr []int) int {
	if len(arr) == 0 {
		panic("Empty array")
	}

	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
func findArgMin(arr []int) int {
	if len(arr) == 0 {
		panic("Empty array")
	}

	min := arr[0]
	pos := 0
	for i, v := range arr {
		if v < min {
			min = v
			pos = i
		}
	}
	return pos
}

// BySecondEntry is a custom type to sort tuples by the second entry
type BySecondEntry [][]int

// Len returns the length of the array
func (a BySecondEntry) Len() int {
	return len(a)
}

// Swap swaps the elements with indexes i and j
func (a BySecondEntry) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less returns true if the tuple with index i should sort before the tuple with index j
func (a BySecondEntry) Less(i, j int) bool {
	return a[i][1] < a[j][1]
}

func splitIntoParagraphs(lines []string) (paragraphs [][]string) {
	var tempParagraph []string
	for i, line := range lines {
		if len(line) != 0 {
			tempParagraph = append(tempParagraph, line)
		}
		if i == len(lines)-1 {
			paragraphs = append(paragraphs, tempParagraph)
			break
		}
		if lines[i+1] == "" {
			paragraphs = append(paragraphs, tempParagraph)
			tempParagraph = nil
		}
	}
	return
}

func convertStringOfNumbers(numberString string) (numbers []int) {
	ns := strings.Split(numberString, " ")
	for _, n := range ns {
		num, err := strconv.Atoi(n)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return
}
