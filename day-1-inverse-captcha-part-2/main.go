package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inp := bufio.NewScanner(os.Stdin)
	inp.Scan()

	a := []rune(inp.Text())
	l := len(a)

	sum := 0
	for i, r := range a {
		if r == a[i+l/2] {
			// Ignoring errors
			digit, _ := strconv.Atoi(string(r))

			sum += digit
		}

		if i == l/2-1 {
			break
		}
	}

	fmt.Println(sum * 2)
}
