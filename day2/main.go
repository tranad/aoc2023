package main

import (  
	"bufio"  
	"fmt"  
	"os"  
	"strconv"  
        "strings"
)  
func main() {
    readFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    var total_possible_game_ids int
    var total_game_power int
    for fileScanner.Scan() {
        line := fileScanner.Text()
        fmt.Println(line)
        game_id, game_possible := solution1(line)
        game_id, game_power := solution2(line)
        if game_possible {
            total_possible_game_ids += game_id
        }
        total_game_power += game_power
    }
    fmt.Println(total_possible_game_ids)
    fmt.Println(total_game_power)
    readFile.Close()
}

func solution2(game_result string) (int, int) {
    game_id_str := strings.Split(game_result, ":")[0][5:]
    game_id, _ := strconv.Atoi(string(game_id_str))
    fmt.Println("Game id ", game_id)

    max_seen_color_map := get_game_results(game_result)
    game_power := 1
    for _,v := range max_seen_color_map {
        game_power *= v
    }
    return game_id, game_power
}

func solution1(game_result string) (int, bool) {
    game_id_str := strings.Split(game_result, ":")[0][5:]
    game_id, _ := strconv.Atoi(string(game_id_str))
    fmt.Println("Game id ", game_id)

    max_seen_color_map := get_game_results(game_result)
    max_color_map := map[string]int {
        "red": 12,
        "green": 13,
        "blue": 14,
    }
    for k, _ := range max_color_map {
        if max_seen_color_map[k] > max_color_map[k] {
            return game_id, false
        }
    }
    return game_id, true
}

func get_game_results(game_result string) map[string]int {
    game_sets := strings.Split(strings.Split(game_result, ":")[1], ";")


    max_seen_color_map := map[string]int {
        "red": 0,
        "green": 0,
        "blue": 0,
    }
    for _,reveal := range game_sets {
        reveal_count := get_reveal_count(reveal)
        for k, v := range reveal_count {
            if max_seen_color_map[k] < v {
                max_seen_color_map[k] = v
            }
        }
        fmt.Println(reveal_count)
    }
    fmt.Println(max_seen_color_map)
    return max_seen_color_map
}

func get_reveal_count(reveal string) map[string]int {
    reveal_map := map[string]int {
        "red": 0,
        "blue": 0,
        "green": 0,
    }

    num_colors := strings.Split(reveal, ",")

    for _,r := range num_colors {
        nc := strings.Split(r, " ")
        num,_ := strconv.Atoi(nc[1])
        color := nc[2]
        reveal_map[color] += num
    }
    return reveal_map
}
