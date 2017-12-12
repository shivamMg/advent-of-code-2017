package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type R []rune

func (r R) Len() int {
	return len(r)
}

func (r R) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r R) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func main() {
	inp := bufio.NewScanner(os.Stdin)
	count := 0

	for inp.Scan() {
		passphrase := inp.Text()
		m := make(map[string]bool)
		valid := true
		for _, word := range strings.Split(passphrase, " ") {
			r := R(word)
			sort.Sort(r)

			if _, ok := m[string(r)]; ok {
				// invalid
				valid = false
				break
			}

			m[string(r)] = true
		}

		if valid {
			count++
		}
	}

	fmt.Println(count)
}
