package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("13")
	start := time.Now().UTC()
	grid := fold(input, 0)
	answer := 0
	for _, ok := range grid {
		if ok {
			answer++
		}
	}
	fmt.Println(answer)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("13")
	start := time.Now().UTC()
	grid := fold(input, -1)
	render(grid)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

type foldInstruction struct {
	axis  string
	value int
}

func fold(input string, instruction int) map[adventofcode2021.Coord]bool {
	grid := make(map[adventofcode2021.Coord]bool)
	var folds []foldInstruction
	for _, line := range adventofcode2021.MustStringList(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			grid[adventofcode2021.Coord{X: adventofcode2021.MustInt(parts[0]), Y: adventofcode2021.MustInt(parts[1])}] = true
			continue
		}
		parts = strings.Split(line, "=")
		axis := parts[0][len(parts[0])-1:]
		value := adventofcode2021.MustInt(parts[1])
		folds = append(folds, foldInstruction{axis: axis, value: value})
	}

	if instruction == -1 {
		for _, f := range folds {
			grid = foldAlong(grid, f)
		}
	} else {
		grid = foldAlong(grid, folds[instruction])
	}

	return grid
}

func foldAlong(grid map[adventofcode2021.Coord]bool, instruction foldInstruction) map[adventofcode2021.Coord]bool {
	maxx, maxy := max(grid)
	if instruction.axis == "y" {
		for y := instruction.value + 1; y <= maxy; y++ {
			for x := 0; x <= maxx; x++ {
				c := adventofcode2021.Coord{X: x, Y: y}
				if grid[c] {
					newy := instruction.value - (y - instruction.value)
					grid[adventofcode2021.Coord{X: x, Y: newy}] = true
				}

			}
		}
		// blank all entries at or below the fold
		for y := instruction.value; y <= maxy; y++ {
			for x := 0; x <= maxx; x++ {
				c := adventofcode2021.Coord{X: x, Y: y}
				delete(grid, c)
			}
		}
	}
	if instruction.axis == "x" {
		for y := 0; y <= maxy; y++ {
			for x := instruction.value; x <= maxx; x++ {
				c := adventofcode2021.Coord{X: x, Y: y}
				if grid[c] {
					newx := instruction.value - (x - instruction.value)
					grid[adventofcode2021.Coord{X: newx, Y: y}] = true
				}

			}
		}
		// blank all entries at or below the fold
		for y := 0; y <= maxy; y++ {
			for x := instruction.value; x <= maxx; x++ {
				c := adventofcode2021.Coord{X: x, Y: y}
				delete(grid, c)
			}
		}
	}

	return grid
}

func max(grid map[adventofcode2021.Coord]bool) (int, int) {
	var maxx, maxy int
	for c := range grid {
		if c.X > maxx {
			maxx = c.X
		}
		if c.Y > maxy {
			maxy = c.Y
		}
	}
	return maxx, maxy
}

func render(grid map[adventofcode2021.Coord]bool) {
	maxx, maxy := max(grid)
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			ok := grid[adventofcode2021.Coord{X: x, Y: y}]
			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
