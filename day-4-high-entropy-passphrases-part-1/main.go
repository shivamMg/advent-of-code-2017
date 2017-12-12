package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inp := bufio.NewScanner(os.Stdin)
	count := 0

	for inp.Scan() {
		passphrase := inp.Text()
		m := make(map[string]bool)
		valid := true
		for _, word := range strings.Split(passphrase, " ") {
			if _, ok := m[word]; ok {
				// invalid
				valid = false
				break
			}
			m[word] = true
		}

		if valid {
			count++
		}
	}

	fmt.Println(count)
}
