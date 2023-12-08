package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	// "sort"
	"strconv"
	"strings"
	"time"
	// "slices"
)

func parseLine(line string) (string, []int) {
	split := strings.Split(line, ":")
	label := split[0]
	numbers := convertStringOfNumbers(split[1])
	return label, numbers
}

func countWaysToBeatRecord(time int, dist int) int {
	time_f64 := float64(time)
	dist_f64 := float64(dist)
	sqrt_discriminant := math.Sqrt(time_f64*time_f64 - 4*dist_f64)
	rt1 := (time_f64 + sqrt_discriminant) / 2
	rt2 := (time_f64 - sqrt_discriminant) / 2
	return int(math.Floor(rt1) - math.Ceil(rt2) + 1)
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

	// Part 1
	_, times := parseLine(lines[0])
	_, dists := parseLine(lines[1])

	var counts []int
	for i := 0; i < len(times); i++ {
		count := countWaysToBeatRecord(times[i], dists[i])
		counts = append(counts, count)
	}
	multCounts := 1
	for _, count := range counts {
		multCounts *= count
	}
	fmt.Println(multCounts)

	// Part 2
	time_str := ""
	for _, t := range times {
		time_str += strconv.Itoa(t)
	}
	time, _ := strconv.Atoi(time_str)

	dist_str := ""
	for _, d := range dists {
		dist_str += strconv.Itoa(d)
	}
	dist, _ := strconv.Atoi(dist_str)
	fmt.Println(time, dist)
	fmt.Println(countWaysToBeatRecord(time, dist))

	return
}

// Previous utilitiy functions

// How are these not a standard functions iswtfg
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
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
