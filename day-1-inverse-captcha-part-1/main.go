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
	a = append(a, a[0])

	sum := 0
	for i, r := range a {
		if r == a[i+1] {
			// Ignoring errors
			digit, _ := strconv.Atoi(string(r))

			sum += digit
		}

		if i == len(a)-2 {
			break
		}
	}

	fmt.Println(sum)
}
