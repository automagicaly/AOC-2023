package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	power_sum := 0
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
		power := expected_min_count["red"] * expected_min_count["green"] * expected_min_count["blue"]
		fmt.Printf("Game power: %d\n", power)
		power_sum += power
	}
	fmt.Printf("sum: %d\n", power_sum)
}
