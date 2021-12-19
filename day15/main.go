package main

import (
	"fmt"
	"time"

	"github.com/smasher164/graph"
	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("15")
	start := time.Now().UTC()
	risk := lowestRisk(input)
	fmt.Println(risk)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("15")
	start := time.Now().UTC()
	risk := lowestRiskFiveTimes(input)
	fmt.Println(risk)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func lowestRisk(input string) int {
	grid := parseGrid(input)
	path := graph.Single(grid, 0)
	return path.Distance(0, grid.coordToVertex[adventofcode2021.Coord{X: grid.maxx, Y: grid.maxy}])
}

func lowestRiskFiveTimes(input string) int {
	grid := parseGrid(input)
	grid.multiply(5)
	path := graph.Single(grid, 0)
	return path.Distance(0, grid.coordToVertex[adventofcode2021.Coord{X: grid.maxx, Y: grid.maxy}])
}

func Print(grid map[adventofcode2021.Coord]int, maxx, maxy int) {
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			fmt.Printf("%02d ", grid[adventofcode2021.Coord{X: x, Y: y}])
		}
		fmt.Println()
	}
	fmt.Println()
}

var mods = []adventofcode2021.Coord{
	{X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0},
}

func upDownLeftRight(coord adventofcode2021.Coord, maxx, maxy int) []adventofcode2021.Coord {
	var result []adventofcode2021.Coord
	for _, m := range mods {
		next := adventofcode2021.Coord{X: coord.X + m.X, Y: coord.Y + m.Y}
		if next.X < 0 || next.Y < 0 || next.X > maxx || next.Y > maxy {
			continue
		}
		result = append(result, next)
	}
	return result
}

type Grid struct {
	grid          map[adventofcode2021.Coord]int
	vertexToCoord map[graph.Vertex]adventofcode2021.Coord
	coordToVertex map[adventofcode2021.Coord]graph.Vertex

	maxx int
	maxy int
	maxi int
}

func (g *Grid) Vertices() []graph.Vertex {
	var result []graph.Vertex
	for c := range g.vertexToCoord {
		result = append(result, graph.Vertex(c))
	}
	return result
}

func (g *Grid) Neighbors(v graph.Vertex) []graph.Vertex {
	c := g.vertexToCoord[v]
	neighbours := upDownLeftRight(c, g.maxx, g.maxy)
	var result []graph.Vertex
	for _, n := range neighbours {
		result = append(result, g.coordToVertex[n])
	}
	return result
}

func (g *Grid) Weight(u, v graph.Vertex) int {
	c := g.vertexToCoord[v]
	return g.grid[c]
}

func parseGrid(input string) *Grid {
	g := &Grid{
		grid:          make(map[adventofcode2021.Coord]int),
		vertexToCoord: make(map[graph.Vertex]adventofcode2021.Coord),
		coordToVertex: make(map[adventofcode2021.Coord]graph.Vertex),
	}
	var maxx, maxy, i int
	for y, line := range adventofcode2021.MustStringList(input, "\n") {
		if y > maxy {
			maxy = y
		}
		for x, c := range line {
			if x > maxx {
				maxx = x
			}
			g.grid[adventofcode2021.Coord{X: x, Y: y}] = adventofcode2021.MustInt(string(c))
			g.vertexToCoord[graph.Vertex(i)] = adventofcode2021.Coord{X: x, Y: y}
			g.coordToVertex[adventofcode2021.Coord{X: x, Y: y}] = graph.Vertex(i)
			i++
		}
	}
	g.maxx = maxx
	g.maxy = maxy
	g.maxi = i
	return g
}

func (g *Grid) multiply(by int) {
	i := g.maxi
	inc := 0
	for gy := 0; gy < by; gy++ {
		curInc := inc
		for gx := 0; gx < by; gx++ {
			if gy == 0 && gx == 0 {
				curInc++
				continue
			}
			for y := 0; y <= g.maxy; y++ {
				for x := 0; x <= g.maxx; x++ {
					xfactor := (g.maxx + 1) * gx
					yfactor := (g.maxy + 1) * gy
					c := adventofcode2021.Coord{X: x, Y: y}
					nc := adventofcode2021.Coord{X: x + xfactor, Y: y + yfactor}
					val := g.grid[c]
					val += curInc
					if val > 9 {
						val = val - 9
					}
					g.grid[nc] = val
					g.vertexToCoord[graph.Vertex(i)] = nc
					g.coordToVertex[nc] = graph.Vertex(i)
					i++
				}
			}
			curInc++
		}
		inc++
	}
	g.maxx = ((g.maxx + 1) * by) - 1
	g.maxy = ((g.maxy + 1) * by) - 1
	g.maxi = i
}
