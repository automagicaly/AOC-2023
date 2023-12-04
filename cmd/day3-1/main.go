package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var deltas = [8][2]int{
    {-1, -1},
    {-1, 0},
    {-1, 1},
    {0, -1},
    {0, 1},
    {1, -1},
    {1, 0},
    {1, 1},
}

func calculate_ratio(x, y int, lines []string) int {
    fmt.Println("===============================")
    visited := make(map[[2]int]bool)
    numbers := make([]int, 0)
    number_str := ""
    for _, d := range deltas {
        fmt.Println(d)
        x_ := x + d[0]
        y_ := y + d[1]
        r := lines[y_][x_]
        if y_ < 0 || y_ >= len(lines) || x_ < 0 || x_ >= len(lines[y_]) || visited[[2]int{x_, y_}] || r < '0' || r > '9'{
            continue
        }
        visited[[2]int{x_, y_}] = true
        for ;x_ >=0; x_-- {
            visited[[2]int{x_, y_}] = true
            r = lines[y_][x_]
            if r < '0' || r > '9' {
                break
            }
            number_str = string(r) + number_str
        }
        x_ = x + d[0] + 1
        for ;x_ < len(lines[y_]); x_++ {
            visited[[2]int{x_, y_}] = true
            r = lines[y_][x_]
            if r < '0' || r > '9' {
                break
            }
            number_str += string(r)
        }
        if len(number_str) > 0 {
            fmt.Printf("found id: %s\n", number_str)
            num, err := strconv.Atoi(number_str)
            if err != nil {
                panic(err)
            }
            numbers = append(numbers, num)
        }
        number_str = ""
    }

    ratio := 0
    if len(numbers) == 2 {
        ratio = numbers[0] * numbers[1]
    }
    fmt.Printf("x: %d\ny: %d\nnumbers[]: [", x, y)
    for _, i := range numbers {
        fmt.Printf(" %d,", i)
    }
    fmt.Println(" ]")
    fmt.Printf("ratio: %d\n", ratio)

    return ratio
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    var lines []string
	for scanner.Scan() {
        new_line := scanner.Text()
        fmt.Println(new_line)
        lines = append(lines, new_line)
	}
    total_sum := 0
    for y, line := range lines{
        for x, r := range line {
            if r == '*' {
                total_sum += calculate_ratio(x, y, lines)
            }
        }
    }
    fmt.Println(total_sum)
}
