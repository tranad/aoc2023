package main

import (  
	"bufio"  
	"fmt"  
	"os"  
	"strconv"  
	// "unicode"  
)  
func main() {

    readFile, err := os.Open("input.txt")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    var total_calibration int
    for fileScanner.Scan() {
        line := fileScanner.Text()
        total_calibration += solution2(line)
    }

    fmt.Println(total_calibration)
    readFile.Close()
}

func solution1(input string) int {
        fmt.Println(input)
        var calibration int
        first_digit := -1
        last_digit := -1
        for pos, _ := range  input {
            // fmt.Printf("character %c starts at byte position %d\n", char, pos)
            char_left := input[pos]
            char_right := input[len(input) - pos - 1]
            val_left, err_left := strconv.Atoi(string(char_left))
            val_right, err_right := strconv.Atoi(string(char_right))
            if (err_left == nil) && (first_digit == -1) {
                first_digit = val_left
            }
            if (err_right == nil) && (last_digit == -1) {
                last_digit = val_right
            }
            if (first_digit >= 0) && (last_digit >=0) {
                calibration += 10*first_digit + last_digit
                fmt.Println(calibration)
                break
            }
        }
        return calibration
}


func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func check_get_digit(s string, i int, rev bool) (int,int) {
    sol2map := map[string]int{
        "one": 1,
        "two": 2,
        "six": 6,
        "four": 4,
        "five": 5,
        "nine": 9,
        "three": 3,
        "seven": 7,
        "eight": 8,
    }
    sol2map_rev := map[string]int{
        "eno": 1,
        "owt": 2,
        "xis": 6,
        "ruof": 4,
        "evif": 5,
        "enin": 9,
        "eerht": 3,
        "neves": 7,
        "thgie": 8,
    }
    var sol2 = make(map[bool]map[string]int)
    sol2[true] = sol2map_rev
    sol2[false] = sol2map

    char := s[i]
    val, err := strconv.Atoi(string(char))
    if err == nil {
        return val, i
    }

    if i+3 <= len(s) {
        val3,inmap3 := sol2[rev][s[i:i+3]]
        if inmap3 == true {
            return val3, i+3
        }
        if inmap3 == false && i+4 <= len(s) {
            val4,inmap4 := sol2[rev][s[i:i+4]]
            if inmap4 == true {
                return val4, i+4
            }
            if inmap4 == false && i+5 <= len(s) {
                val5,inmap5 := sol2[rev][s[i:i+5]]
                if inmap5 == true {
                    return val5, i+5
                }
            }
        }
    }
    return -1,-1
}


func solution2(input string) int {
    tupni := Reverse(input)

    fmt.Println(input)
    var calibration int
    first_digit := -1
    last_digit := -1
    for pos, _ := range  input {
        if first_digit == -1 {
            first_digit, _ = check_get_digit(input, pos, false)
        }
        if last_digit == -1 {
            last_digit, _ = check_get_digit(tupni, pos, true)
        }

        if (first_digit >= 0) && (last_digit >=0) {
            calibration += 10*first_digit + last_digit
            fmt.Println("\t", calibration)
            break
        }
    }
    return calibration
}

