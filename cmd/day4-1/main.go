package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    grab_nums_regex, err := regexp.Compile(`\d+`)
    if err != nil {
        panic(err)
    }
    copies := make(map[int]int)
	for scanner.Scan() {
        winning_numbers := make(map[int]int)
        matches := 0
        line := scanner.Text()
        parts := strings.Split(line, ":")
        header := strings.Split(parts[0], " ")
        id, err := strconv.Atoi(header[len(header) - 1])
        if err != nil {
            panic(err)
        }
        copies[id]++
        number_sets := strings.Split(parts[1], "|")

        for _, i := range grab_nums_regex.FindAllString(number_sets[0], -1) {
            num, err := strconv.Atoi(i)
            if err != nil {
                panic(err)
            }
            winning_numbers[num] = 1
        }
        for _, i := range grab_nums_regex.FindAllString(number_sets[1], -1) {
            num, err := strconv.Atoi(i)
            if err != nil {
                panic(err)
            }
            matches += winning_numbers[num]
        }
        if matches > 0 {
            for ;matches > 0; matches-- {
                copies[id+matches] += copies[id]
            }
        }
	}
    total := 0
    for _, v := range copies {
        total += v
    }
    fmt.Println(total)
}
