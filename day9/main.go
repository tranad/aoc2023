package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Sequence struct {
	sequence []int
}

func allZero(arr []int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

func (seq Sequence) Next() (last int) {
	var nseq []int
	for i := 0; i < len(seq.sequence)-1; i++ {
		diff := seq.sequence[i+1] - seq.sequence[i]
		nseq = append(nseq, diff)
	}
	defer func() {
		last += seq.sequence[len(seq.sequence)-1]
	}()
	if allZero(nseq) {
		return 0
	} else {
		return Sequence{nseq}.Next()
	}
}

func (seq Sequence) Prev() (first int) {
	var nseq []int
	for i := 0; i < len(seq.sequence)-1; i++ {
		diff := seq.sequence[i+1] - seq.sequence[i]
		nseq = append(nseq, diff)
	}
	defer func() {
		first = seq.sequence[0] - first
	}()
	if allZero(nseq) {
		return 0
	} else {
		return Sequence{nseq}.Prev()
	}
}

func parseLine(line string) Sequence {
	split := strings.Split(line, " ")
	var sequence []int
	for _, s := range split {
		num, err := strconv.Atoi(s)
		if err == nil {
			sequence = append(sequence, num)
		}
	}
	return Sequence{sequence}
}

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
	defer timer("Day 9")()
	ans1 := 0
	ans2 := 0
	for _, line := range lines {
		seq := parseLine(line)
		ans1 += seq.Next()
		ans2 += seq.Prev()
	}
	fmt.Printf("Part 1: %v\n", ans1)
	fmt.Printf("Part 2: %v\n", ans2)
	return
}

// How are these not a standard functions iswtfg
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
