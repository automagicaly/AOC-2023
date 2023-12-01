package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    reader := bufio.NewScanner(os.Stdin)
    calibration := 0
    for reader.Scan() {
        fd := -1
        ld := -1
        for _, r := range reader.Text() {
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
