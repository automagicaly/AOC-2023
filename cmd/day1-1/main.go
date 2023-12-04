package main

import (
	"bufio"
	"fmt"
    "os"
    "regexp"
)

func main() {
    regex_1,_ := regexp.Compile(`one`)
    regex_2,_ := regexp.Compile(`two`)
    regex_3,_ := regexp.Compile(`three`)
    regex_4,_ := regexp.Compile(`four`)
    regex_5,_ := regexp.Compile(`five`)
    regex_6,_ := regexp.Compile(`six`)
    regex_7,_ := regexp.Compile(`seven`)
    regex_8,_ := regexp.Compile(`eight`)
    regex_9,_ := regexp.Compile(`nine`)

    reader := bufio.NewScanner(os.Stdin)
    calibration := 0
    for reader.Scan() {
        fd := -1
        ld := -1
        line := regex_1.ReplaceAllString(reader.Text(), "o1e")
        line = regex_2.ReplaceAllString(line, "t2o")
        line = regex_3.ReplaceAllString(line, "t3e")
        line = regex_4.ReplaceAllString(line, "4")
        line = regex_5.ReplaceAllString(line, "5e")
        line = regex_6.ReplaceAllString(line, "6")
        line = regex_7.ReplaceAllString(line, "7n")
        line = regex_8.ReplaceAllString(line, "e8t")
        line = regex_9.ReplaceAllString(line, "n9e")
        for _, r := range line {
            if r >= '0' && r <= '9' {
                if fd == -1 {
                    fd = int(r - '0')
                }
                ld = int(r - '0')
            }
        }
        calibrationValue := (fd*10 + ld)
        calibration += calibrationValue
        fmt.Printf("%d\n", calibrationValue)
    }
    fmt.Printf("%d\n", calibration)
}
