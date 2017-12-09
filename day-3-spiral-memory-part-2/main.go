package main

import (
	"fmt"
)

const Size = 400

type Coord struct {
	x, y int
}

func main() {
	x := 0
	fmt.Scan(&x)

	// Access port
	if x == 1 {
		fmt.Println(2)
		return
	}

	grid := NewGrid()
	cur := Coord{Size / 2, Size / 2}
	grid[cur.x][cur.y] = 1

	v := 0
	i := 1
	for i < Size/2 {
		n := 2*i + 1
		var found bool

		cur.Right()
		found, v = run(&grid, cur.Up, &cur, n, x)
		if found {
			break
		}

		cur.Left()
		found, v = run(&grid, cur.Left, &cur, n, x)
		if found {
			break
		}

		cur.Down()
		found, v = run(&grid, cur.Down, &cur, n, x)
		if found {
			break
		}

		cur.Right()
		found, v = run(&grid, cur.Right, &cur, n, x)
		if found {
			break
		}

		i++
	}

	fmt.Println(v)
}

func run(grid *[][]int, movement func(), cur *Coord, n int, x int) (bool, int) {
	s := sum(*grid, (*cur).Adj())
	if s > x {
		return true, s
	}
	(*grid)[(*cur).x][(*cur).y] = s

	for j := 0; j < n-2; j++ {
		movement()

		s := sum(*grid, (*cur).Adj())
		if s > x {
			return true, s
		}
		(*grid)[(*cur).x][(*cur).y] = s
	}

	return false, 0
}

func NewGrid() [][]int {
	g := make([][]int, Size)
	for i := 0; i < Size; i++ {
		g[i] = make([]int, Size)
	}

	return g
}

func (c *Coord) Up() {
	c.x--
}

func (c *Coord) Down() {
	c.x++
}

func (c *Coord) Right() {
	c.y++
}

func (c *Coord) Left() {
	c.y--
}

func (c *Coord) Adj() []Coord {
	return []Coord{
		Coord{c.x, c.y + 1},
		Coord{c.x - 1, c.y + 1},
		Coord{c.x - 1, c.y},
		Coord{c.x - 1, c.y - 1},
		Coord{c.x, c.y - 1},
		Coord{c.x + 1, c.y - 1},
		Coord{c.x + 1, c.y},
		Coord{c.x + 1, c.y + 1},
	}
}

func sum(grid [][]int, coords []Coord) (sum int) {
	sum = 0
	for _, c := range coords {
		sum += grid[c.x][c.y]
	}
	return
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}
