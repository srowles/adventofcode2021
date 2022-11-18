package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("17")
	start := time.Now().UTC()
	fmt.Println(input)
	t := Create(input)

	highestHight := 0
	winCount := 0
	for x := 0; x <= t.targetBottomRight.X; x++ {
		for y := -2000; y < 2000; y++ {
			t.probeVelocity = adventofcode2021.Coord{X: x, Y: y}
			t.probePosition = adventofcode2021.Coord{}
			t.maxY = 0
			for !t.InTarget() && !t.BeyondTarget() {
				t.Step()
			}
			if t.InTarget() {
				winCount++
				if t.maxY > highestHight {
					highestHight = t.maxY
				}
				continue
			}
		}
	}

	fmt.Println(highestHight, winCount, "took:", time.Now().UTC().Sub(start))
}

func Create(input string) *trench {
	t := &trench{}
	tl, br := parseTarget(input)
	t.targetTopLeft = tl
	t.targetBottomRight = br

	return t
}

type trench struct {
	targetTopLeft     adventofcode2021.Coord
	targetBottomRight adventofcode2021.Coord
	probePosition     adventofcode2021.Coord
	probeVelocity     adventofcode2021.Coord
	maxY              int
}

func (t *trench) Step() {
	t.probePosition = adventofcode2021.Coord{X: t.probePosition.X + t.probeVelocity.X, Y: t.probePosition.Y + t.probeVelocity.Y}
	if t.probePosition.Y > t.maxY {
		t.maxY = t.probePosition.Y
	}
	newVelX := t.probeVelocity.X
	switch {
	case t.probeVelocity.X > 0:
		newVelX -= 1
	case t.probeVelocity.X < 0:
		newVelX += 1
	}
	t.probeVelocity = adventofcode2021.Coord{X: newVelX, Y: t.probeVelocity.Y - 1}
}

func (t *trench) InTarget() bool {
	var inX, inY bool
	if t.probePosition.X >= t.targetTopLeft.X && t.probePosition.X <= t.targetBottomRight.X {
		inX = true
	}

	if t.probePosition.Y <= t.targetTopLeft.Y && t.probePosition.Y >= t.targetBottomRight.Y {
		inY = true
	}
	return inX && inY
}

func (t *trench) BeyondTarget() bool {
	var beyondX, beyondY bool
	if t.probePosition.X > t.targetBottomRight.X {
		beyondX = true
	}

	if t.probePosition.Y < t.targetBottomRight.Y {
		beyondY = true
	}
	return beyondX || beyondY
}

func parseTarget(input string) (topLeft, bottomRight adventofcode2021.Coord) {
	_, values, _ := strings.Cut(strings.TrimSpace(input), ": ")
	x, y, _ := strings.Cut(values, ", ")
	_, xx, _ := strings.Cut(x, "=")
	_, yy, _ := strings.Cut(y, "=")

	x1, x2, _ := strings.Cut(xx, "..")
	y1, y2, _ := strings.Cut(yy, "..")

	return adventofcode2021.Coord{
			X: adventofcode2021.MustInt(x1),
			Y: adventofcode2021.MustInt(y2),
		},
		adventofcode2021.Coord{
			X: adventofcode2021.MustInt(x2),
			Y: adventofcode2021.MustInt(y1),
		}
}
