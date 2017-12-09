package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	inp := bufio.NewScanner(os.Stdin)
	inp.Scan()

	x, _ := strconv.Atoi(inp.Text())

	// Access port
	if x == 1 {
		fmt.Println(0)
		return
	}

	m := getBottomRightSquare(x)
	n := int((math.Sqrt(float64(m)) - 1) / 2)

	// Steps from mid square to access port
	steps1 := n

	/**
	 * Get the edge that x resides in
	 * Calculate that edge's mid square
	 * Total steps = (steps from x to mid square) + (steps from mid square to access port)
	 */

	dist := m - x

	mid := 0
	if dist < (2 * n) {
		mid = m - n
	} else if dist < 2*(2*n) {
		mid = m - 3*n
	} else if dist < 3*(2*n) {
		mid = m - 5*n
	} else if dist < 4*(2*n) {
		mid = m - 7*n
	}

	// steps from x to mid square
	steps2 := int(math.Abs(float64(mid - x)))

	fmt.Println(steps1 + steps2)
}

func getBottomRightSquare(x int) int {
	// get next odd number that is a perfect square
	sqrt := int(math.Sqrt(float64(x)))

	if sqrt%2 == 0 {
		sqrt++
		return sqrt * sqrt
	}

	// number was itself a perfect square
	if sqrt*sqrt == x {
		return x
	}

	// the next odd number
	sqrt += 2
	return sqrt * sqrt
}
