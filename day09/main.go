package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("9")
	start := time.Now().UTC()
	answer := getRisk(input)
	fmt.Println(answer)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("9")
	start := time.Now().UTC()
	answer := getBasins(input)
	fmt.Println(answer)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func getRisk(input string) int {
	grid, maxx, maxy := getGrid(input)

	risk := 0
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			c := adventofcode2021.Coord{X: x, Y: y}
			v := getVal(grid, c)
			if v < getVal(grid, adventofcode2021.Coord{X: x - 1, Y: y}) &&
				v < getVal(grid, adventofcode2021.Coord{X: x + 1, Y: y}) &&
				v < getVal(grid, adventofcode2021.Coord{X: x, Y: y - 1}) &&
				v < getVal(grid, adventofcode2021.Coord{X: x, Y: y + 1}) {
				risk += v + 1
			}
		}
	}
	return risk
}

func getBasins(input string) int {
	grid, maxx, maxy := getGrid(input)

	var sizes []int
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			c := adventofcode2021.Coord{X: x, Y: y}
			v := getVal(grid, c)
			if v < getVal(grid, adventofcode2021.Coord{X: x - 1, Y: y}) &&
				v < getVal(grid, adventofcode2021.Coord{X: x + 1, Y: y}) &&
				v < getVal(grid, adventofcode2021.Coord{X: x, Y: y - 1}) &&
				v < getVal(grid, adventofcode2021.Coord{X: x, Y: y + 1}) {
				// low spot, try and "grow" this spot to get full size
				size := findSize(grid, c)
				sizes = append(sizes, size)
			}
		}
	}

	sort.Ints(sizes)
	l := len(sizes) - 1
	return sizes[l] * sizes[l-1] * sizes[l-2]
}

// expand from a point until we hit 9's
func findSize(grid map[adventofcode2021.Coord]int, c adventofcode2021.Coord) int {
	var points []adventofcode2021.Coord
	points = append(points, c)
	idx := 0
	for {
		p := points[idx]
		c := adventofcode2021.Coord{X: p.X - 1, Y: p.Y}
		if getVal(grid, c) < 9 {
			points = appendIfNotAlreadyAdded(points, c)
		}
		c = adventofcode2021.Coord{X: p.X, Y: p.Y - 1}
		if getVal(grid, c) < 9 {
			points = appendIfNotAlreadyAdded(points, c)
		}
		c = adventofcode2021.Coord{X: p.X + 1, Y: p.Y}
		if getVal(grid, c) < 9 {
			points = appendIfNotAlreadyAdded(points, c)
		}
		c = adventofcode2021.Coord{X: p.X, Y: p.Y + 1}
		if getVal(grid, c) < 9 {
			points = appendIfNotAlreadyAdded(points, c)
		}
		if len(points) == idx+1 {
			break
		}
		idx++
	}
	// de-dupe
	deduped := make(map[adventofcode2021.Coord]bool)
	for _, p := range points {
		deduped[p] = true
	}
	return len(deduped)
}

func appendIfNotAlreadyAdded(points []adventofcode2021.Coord, point adventofcode2021.Coord) []adventofcode2021.Coord {
	for _, p := range points {
		if p == point {
			return points
		}
	}
	points = append(points, point)
	return points
}

func getGrid(input string) (map[adventofcode2021.Coord]int, int, int) {
	input = strings.TrimSpace(input)
	grid := make(map[adventofcode2021.Coord]int)
	var maxx, maxy int
	for y, row := range strings.Split(input, "\n") {
		for x, value := range row {
			c := adventofcode2021.Coord{X: x, Y: y}
			grid[c] = int(value - '0')
			maxx = x
		}
		maxy = y
	}
	return grid, maxx, maxy
}

func getVal(grid map[adventofcode2021.Coord]int, c adventofcode2021.Coord) int {
	r, ok := grid[c]
	if !ok {
		return 9
	}
	return r
}
