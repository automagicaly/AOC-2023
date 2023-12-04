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
    total := 0
	for scanner.Scan() {
        winning_numbers := make(map[int]int)
        matches := 0
        line := scanner.Text()
        line = strings.Split(line, ":")[1]
        number_sets := strings.Split(line, "|")

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
            total += 1 << (matches - 1)
        }
	}
    fmt.Println(total)
}
