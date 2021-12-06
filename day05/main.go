package main

import (
	"fmt"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("5")
	fmt.Println(countOverlaps(input, false))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("5")
	fmt.Println(countOverlaps(input, true))
}

type grid struct {
	points map[adventofcode2021.Coord]int
}

func (g grid) String() string {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			coord := adventofcode2021.Coord{X: x, Y: y}
			count, ok := g.points[coord]
			if !ok {
				fmt.Print(".")
			} else {
				fmt.Print(count)
			}
		}
		fmt.Println()
	}
	return ""
}

func countOverlaps(input string, includeDiagonals bool) int {
	g := grid{
		points: make(map[adventofcode2021.Coord]int),
	}
	lines := adventofcode2021.MustStringList(input, "\n")
	linesAdded := 0
	for _, line := range lines {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			panic(err)
		}
		addLine(g, adventofcode2021.Coord{X: x1, Y: y1}, adventofcode2021.Coord{X: x2, Y: y2}, includeDiagonals)
		linesAdded++
	}

	count := 0
	for _, c := range g.points {
		if c > 1 {
			count++
		}
	}

	return count
}

func addLine(g grid, start, end adventofcode2021.Coord, includeDiagonals bool) {
	if start.X == end.X {
		addVerticalLine(g, start, end)
	}
	if start.Y == end.Y {
		addHorizontalLine(g, start, end)
	}
	if includeDiagonals && abs(start.X-end.X) == abs(start.Y-end.Y) {
		addDiagonalLine(g, start, end)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func addHorizontalLine(g grid, start, end adventofcode2021.Coord) {
	if start.X > end.X {
		start, end = end, start
	}
	for x := start.X; x <= end.X; x++ {
		g.points[adventofcode2021.Coord{X: x, Y: start.Y}] = g.points[adventofcode2021.Coord{X: x, Y: start.Y}] + 1
	}

}

func addVerticalLine(g grid, start, end adventofcode2021.Coord) {
	if start.Y > end.Y {
		start, end = end, start
	}
	for y := start.Y; y <= end.Y; y++ {
		g.points[adventofcode2021.Coord{X: start.X, Y: y}] = g.points[adventofcode2021.Coord{X: start.X, Y: y}] + 1
	}
}

func addDiagonalLine(g grid, start, end adventofcode2021.Coord) {
	if start.X > end.X {
		start, end = end, start
	}

	incy := 1
	if start.Y > end.Y {
		incy = -1
	}

	for a := 0; a <= end.X-start.X; a++ {
		g.points[adventofcode2021.Coord{X: start.X + a, Y: start.Y + (a * incy)}] = g.points[adventofcode2021.Coord{X: start.X + a, Y: start.Y + (a * incy)}] + 1
	}
}
