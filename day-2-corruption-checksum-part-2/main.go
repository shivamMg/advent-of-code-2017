package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp := bufio.NewScanner(os.Stdin)

	sum := 0
	for inp.Scan() {
		quotient := 0

		var nums []int
		for _, s := range strings.Split(inp.Text(), "\t") {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}

	label:
		for i, num := range nums {
			j := i + 1
			for j < len(nums) {
				x, y := sort(num, nums[j])
				if y%x == 0 {
					quotient = y / x
					break label
				}
				j++
			}
		}

		sum += quotient
	}

	fmt.Println(sum)
}

func sort(x, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}
