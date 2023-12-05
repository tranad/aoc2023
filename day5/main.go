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

type FarmMap struct {
	source                string
	destination           string
	sourceDestinationCode [][]int
}

func (fm FarmMap) Convert(s int) (d int, dn string) {
	dn = fm.destination
	d = s
	for _, v := range fm.sourceDestinationCode {
		sourceStart := v[0]
		destinationStart := v[1]
		rangeLength := v[2]
		if sourceStart <= s && s < sourceStart+rangeLength {
			d = destinationStart + (s - sourceStart)
			return
		}
	}
	return
}

func (fm FarmMap) ReverseConvert(d int) (s int, sn string) {
	sn = fm.source
	s = d
	for _, v := range fm.sourceDestinationCode {
		sourceStart := v[0]
		destinationStart := v[1]
		rangeLength := v[2]
		if destinationStart <= s && s < destinationStart+rangeLength {
			s = sourceStart + (d - destinationStart)
			return
		}
	}
	return
}

func traverseFarmMaps(input int, source string, destination string, fms []FarmMap) (output int) {
	output = input
	for _, fm := range fms {
		output, _ = fm.Convert(output)
	}

	// Check source destination in case want to stop somwhere intermediate
	// fmt.Println(output)
	// for {
	//     for _,fm := range fms {
	//         // fmt.Println(fm)
	//         if fm.source == source {
	//             output, source = fm.Convert(output)
	//         }
	//         if source == destination {
	//             break
	//         }
	//         // fmt.Println(output)
	//     }
	//     if source == destination {
	//         break
	//     }
	// }
	// // fmt.Println()
	return
}

func reverseTraverseFarmMaps(input int, source string, destination string, fms []FarmMap) (output int) {
	output = input
	for i := len(fms) - 1; i >= 0; i-- {
		fm := fms[i]
		output, _ = fm.ReverseConvert(output)
	}
	return
}

func getSeedList(paragraph []string) []int {
	line := paragraph[0]
	sl := strings.Split(line, ": ")
	list_half := sl[1]
	seeds := convertStringOfNumbers(list_half)
	return seeds
}
func getFarmMap(paragraph []string) FarmMap {
	header := paragraph[0]
	sd := strings.Split(strings.Split(header, " map:")[0], "-")
	source := sd[0]
	destination := sd[2]

	var sdcode [][]int
	for i, line := range paragraph {
		if i == 0 {
			continue
		}
		dsr := convertStringOfNumbers(line)
		sdcode = append(sdcode, []int{dsr[1], dsr[0], dsr[2]})
	}
	return FarmMap{source, destination, sdcode}
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

	paragraphs := splitIntoParagraphs(lines)
	var farmMaps []FarmMap
	for _, paragraph := range paragraphs[1:] {
		farmMaps = append(farmMaps, getFarmMap(paragraph))
	}

	// Part 1
	seeds := getSeedList(paragraphs[0])
	locations := make([]int, 0, len(seeds))
	for _, seed := range seeds {
		loc := traverseFarmMaps(seed, "seed", "location", farmMaps)
		locations = append(locations, loc)
	}
	fmt.Printf("Part 1: %d \n", findMin(locations))

	// Part 2: try backwards
	defer timer("part2")()
	var seeds2 [][]int
	for i := 0; i < len(seeds); i += 2 {
		seeds2 = append(seeds2, []int{seeds[i], seeds[i+1]})
	}

	sortedLocCode := farmMaps[len(farmMaps)-1].sourceDestinationCode
	sort.Sort(BySecondEntry(sortedLocCode))
	for _, code := range sortedLocCode {
		loc_start := code[1]
		loc_range := code[2]
		for k := 0; k < loc_range; k++ {
			loc := loc_start + k
			seed := reverseTraverseFarmMaps(loc, "location", "seed", farmMaps)

			// Check if valid seed
			for _, sr := range seeds2 {
				if sr[0] <= seed && seed < sr[0]+sr[1] {
					panicMessage := fmt.Sprintf("Part 2: %d", loc)
					panic(panicMessage) // too lazy type break x3
				}
			}
		}
	}
	return
}

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
