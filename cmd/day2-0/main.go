package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	id_sum := 0
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ":")
		draws := strings.Split(data[1], ";")
		expected_min_count := make(map[string]int)
		for _, draw := range draws {
			color_counts := strings.Split(draw, ",")
			for _, color_count := range color_counts {
				tmp := strings.Split(strings.TrimSpace(color_count), " ")
				count, err := strconv.Atoi(tmp[0])
				if err != nil {
					panic(err)
				}
				color := tmp[1]
				expected_min_count[color] = max(expected_min_count[color], count)
			}
		}
		if expected_min_count["red"] > MAX_RED || expected_min_count["green"] > MAX_GREEN || expected_min_count["blue"] > MAX_BLUE {
			continue
		}
		game_id, err := strconv.Atoi(strings.Split(data[0], " ")[1])
		if err != nil {
			panic(err)
		}
        fmt.Printf("Valid game: %d\n", game_id)
		id_sum += game_id
	}
    fmt.Printf("sum: %d\n", id_sum)
}
