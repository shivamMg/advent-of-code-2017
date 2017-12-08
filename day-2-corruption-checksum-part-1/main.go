package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = ^MaxInt
)

func main() {
	inp := bufio.NewScanner(os.Stdin)

	sum := 0
	for inp.Scan() {
		min, max := MaxInt, MinInt

		for _, s := range strings.Split(inp.Text(), "\t") {
			no, _ := strconv.Atoi(s)
			if no < min {
				min = no
			}
			if no > max {
				max = no
			}
		}

		sum += max - min
	}

	fmt.Println(sum)
}
