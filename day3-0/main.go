package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    var lines []string
	for scanner.Scan() {
        new_line := scanner.Text()
        fmt.Println(new_line)
        lines = append(lines, new_line)
	}
    deltas := [][]int{
        {-1, -1},
        {-1, 0},
        {-1, 1},
        {0, -1},
        {0, 1},
        {1, -1},
        {1, 0},
        {1, 1},
    }
    id_sum := 0
    for y, line := range lines{
        number_str := ""
        valid := false
        for x, r := range line {
            if r >= '0' && r <= '9' {
                number_str += string(r)
                for _, d := range deltas {
                    x_ := x + d[0]
                    y_ := y + d[1]
                    if y_ < 0 || y_ >= len(lines) || x_ < 0 || x_ >= len(lines[y_]) {
                        continue
                    }
                    s := lines[y_][x_]
                    if s == '.' {
                        continue
                    }
                    if s < '0' || s > '9' {
                        valid = true
                    }
                }
                fmt.Printf("number_str: %s\n", number_str)
                fmt.Printf("valid: %t\n", valid)
            } else {
                if !valid {
                    number_str = ""
                    continue 
                }
                id, err := strconv.Atoi(number_str)
                if err != nil {
                    panic(err)
                }
                fmt.Printf("id: %d\n", id)
                id_sum += id
                number_str = ""
                valid = false
            }
        }
        if valid {
            id, err := strconv.Atoi(number_str)
            if err != nil {
                panic(err)
            }
            fmt.Printf("id: %d\n", id)
            id_sum += id
        }
        fmt.Printf("sum: %d\n", id_sum)
    }
    fmt.Printf("sum: %d\n", id_sum)
}
