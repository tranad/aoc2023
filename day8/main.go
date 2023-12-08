package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"
	// "sort"
	// "strconv"
	// "slices"
)

type Node struct {
	id    string
	left  string
	right string
}

func part1(startingNodeID string, nodeMap map[string]Node, instructions string) (numSteps int) {
	currentNode := startingNodeID
	for {
		for _, inst := range instructions {
			LR := string(inst)
			switch {
			case LR == "L":
				currentNode = nodeMap[currentNode].left
			case LR == "R":
				currentNode = nodeMap[currentNode].right
			default:
				panic("No such node")
			}
			numSteps++
			if currentNode == "ZZZ" {
				break
			}

		}
		if currentNode == "ZZZ" {
			break
		}
	}
	return numSteps
}

type Result struct {
	id          string
	zLocs       []int
	zLocIDs     []string
	cycleLen    int
	cycleOffset int
}

func part2(ch chan Result, wg *sync.WaitGroup, startingNodeID string, nodeMap map[string]Node, instructions string) {
	defer wg.Done()
	currentNode := startingNodeID
	var zLocs []int
	var zLocIDs []string
	numSteps := 0
	nodeHistory := make(map[string][]int)
	nodeHistory[startingNodeID] = []int{0}

	notComplete := true
	for notComplete {
		for _, inst := range instructions {
			LR := string(inst)
			switch {
			case LR == "L":
				currentNode = nodeMap[currentNode].left
			case LR == "R":
				currentNode = nodeMap[currentNode].right
			default:
				panic("No such node")
			}
			numSteps++

			// Record Z hit
			if string(currentNode[2]) == "Z" {
				zLocs = append(zLocs, numSteps)
				zLocIDs = append(zLocIDs, currentNode)
			}

			// Record nodeHistory
			_, visited := nodeHistory[currentNode]
			if visited && len(zLocs) == 2 && numSteps >= len(instructions) {
				cycleLen := numSteps
				cycleOffset := nodeHistory[currentNode][0]
				r := Result{startingNodeID, zLocs, zLocIDs, cycleLen, cycleOffset}
				notComplete = false
				ch <- r
				break
			} else {
				nodeHistory[currentNode] = append(nodeHistory[currentNode], numSteps)
			}
		}
	}
}

func parseLine(line string) Node {
	split := strings.Split(line, " = ")
	id := split[0]
	lr := strings.Split(split[1], ", ")
	left := lr[0][1:]
	right := lr[1][:len(lr[1])-1]
	return Node{id, left, right}
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
	nodeMap := make(map[string]Node)
	for i, line := range lines {
		if i >= 2 {
			node := parseLine(line)
			// fmt.Println(node)
			// nodes = append(nodes, node)
			nodeMap[node.id] = node
		}
	}
	// MAIN
	instructions := lines[0]

	// Part 1
	ans1 := part1("AAA", nodeMap, instructions)
	fmt.Printf("Part 1: %v\n", ans1)

	// Part 2
	var startingNodeIDs []string
	for k, _ := range nodeMap {
		if string(k[2]) == "A" {
			startingNodeIDs = append(startingNodeIDs, k)
		}
	}

	resultChan := make(chan Result)
	var wg sync.WaitGroup

	// Number of workers
	numWorkers := len(startingNodeIDs)
	// trueCount := 0
	wg.Add(numWorkers)

	// Start workers
	fmt.Println(startingNodeIDs)
	for _, startNodeID := range startingNodeIDs {
		go part2(resultChan, &wg, startNodeID, nodeMap, instructions)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Main goroutine reads from the result channel
	var cycles []int64
	for result := range resultChan {
		fmt.Printf("Worker %v\n", result)
		cycles = append(cycles, int64(result.cycleLen-result.cycleOffset))
	}
	ans2 := calculateLCM(cycles)
	fmt.Printf("Part 2: %v\n", ans2)

	return
}

func gcd(a, b *big.Int) *big.Int {
	for b.Sign() != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

func lcm(a, b *big.Int) *big.Int {
	g := gcd(a, b)
	if g.Sign() == 0 {
		return new(big.Int)
	}
	return new(big.Int).Abs(a.Mul(a, b)).Div(a, g)
}

func calculateLCM(numbers []int64) *big.Int {
	if len(numbers) == 0 {
		return new(big.Int)
	}

	result := big.NewInt(numbers[0])
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, big.NewInt(numbers[i]))
	}

	return result
}

// How are these not a standard functions iswtfg
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
