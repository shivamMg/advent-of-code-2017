package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	list := make([]int, 0)
	inp := bufio.NewScanner(os.Stdin)

	for inp.Scan() {
		i, _ := strconv.Atoi(inp.Text())
		list = append(list, i)
	}

	i := 0
	steps := 0
	for i < len(list) {
		offset := list[i]

		if offset >= 3 {
			list[i]--
		} else {
			list[i]++
		}

		i = i + offset

		steps++
	}

	fmt.Println(steps)
}
