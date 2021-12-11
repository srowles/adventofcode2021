package main

import (
	"fmt"
	"time"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("11")
	start := time.Now().UTC()
	answer := countFlashes(input, 100)
	fmt.Println(answer)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("11")
	start := time.Now().UTC()
	answer := findAllStep(input)
	fmt.Println(answer)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func countFlashes(input string, steps int) int {
	lines := adventofcode2021.MustStringList(input, "\n")
	grid := make(map[adventofcode2021.Coord]int)
	maxy := len(lines)
	maxx := 0
	for y, line := range lines {
		maxx = len(line)
		for x, r := range line {
			c := adventofcode2021.Coord{X: x, Y: y}
			energy := int(r) - int('0')
			grid[c] = int(energy)
		}
	}

	flashCount := 0
	for i := 0; i < steps; i++ {
		flashes := stepLine(grid, maxx, maxy)

		for c := range flashes {
			grid[c] = 0
			flashCount++
		}

		// fmt.Println()
		// print(grid, maxx, maxy)
	}

	return flashCount
}

func findAllStep(input string) int {
	lines := adventofcode2021.MustStringList(input, "\n")
	grid := make(map[adventofcode2021.Coord]int)
	maxy := len(lines)
	maxx := 0
	for y, line := range lines {
		maxx = len(line)
		for x, r := range line {
			c := adventofcode2021.Coord{X: x, Y: y}
			energy := int(r) - int('0')
			grid[c] = int(energy)
		}
	}

	step := 1
	for {
		flashes := stepLine(grid, maxx, maxy)

		for c := range flashes {
			grid[c] = 0
		}

		if len(flashes) == 100 {
			return step
		}
		step++
	}
}

func stepLine(grid map[adventofcode2021.Coord]int, maxx, maxy int) map[adventofcode2021.Coord]bool {
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			c := adventofcode2021.Coord{X: x, Y: y}
			grid[c] = grid[c] + 1
		}
	}

	flashes := make(map[adventofcode2021.Coord]bool)
	for {
		flash := false
		for y := 0; y < maxy; y++ {
			for x := 0; x < maxx; x++ {
				c := adventofcode2021.Coord{X: x, Y: y}
				if grid[c] > 9 {
					if !flashes[c] {
						flash = true
						flashes[c] = true
						for addc := range adjacent {
							grid[adventofcode2021.Coord{X: c.X + addc.X, Y: c.Y + addc.Y}] = grid[adventofcode2021.Coord{X: c.X + addc.X, Y: c.Y + addc.Y}] + 1
						}
					}
				}
			}
		}
		if !flash {
			break
		}
	}

	return flashes
}

var adjacent = map[adventofcode2021.Coord]bool{
	{X: -1, Y: -1}: true,
	{X: -1, Y: 0}:  true,
	{X: -1, Y: 1}:  true,
	{X: 0, Y: -1}:  true,
	{X: 0, Y: 1}:   true,
	{X: 1, Y: -1}:  true,
	{X: 1, Y: 0}:   true,
	{X: 1, Y: 1}:   true,
}

func print(grid map[adventofcode2021.Coord]int, maxx, maxy int) {
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			fmt.Print(grid[adventofcode2021.Coord{X: x, Y: y}])
		}
		fmt.Println()
	}
}
