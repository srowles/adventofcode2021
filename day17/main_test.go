package main

import (
	"testing"

	"github.com/srowles/adventofcode2021"
	"github.com/stretchr/testify/assert"
)

func TestParseTargetArea(t *testing.T) {
	topLeft, bottomRight := parseTarget("target area: x=20..30, y=-10..-5")
	assert.Equal(t, adventofcode2021.Coord{X: 20, Y: -5}, topLeft)
	assert.Equal(t, adventofcode2021.Coord{X: 30, Y: -10}, bottomRight)
}

func TestStep(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		startPosition adventofcode2021.Coord
		startVelocity adventofcode2021.Coord
		steps         int
		endPosition   adventofcode2021.Coord
		inTarget      bool
	}{
		"zero": {},
		"one": {
			startVelocity: adventofcode2021.Coord{X: 7, Y: 2},
			steps:         1,
			endPosition:   adventofcode2021.Coord{X: 7, Y: 2},
		},
		"two": {
			startVelocity: adventofcode2021.Coord{X: 7, Y: 2},
			steps:         2,
			endPosition:   adventofcode2021.Coord{X: 13, Y: 3},
		},
		"three": {
			startVelocity: adventofcode2021.Coord{X: 7, Y: 2},
			steps:         3,
			endPosition:   adventofcode2021.Coord{X: 18, Y: 3},
		},
		"four": {
			startVelocity: adventofcode2021.Coord{X: 7, Y: 2},
			steps:         4,
			endPosition:   adventofcode2021.Coord{X: 22, Y: 2},
		},
		"five": {
			startVelocity: adventofcode2021.Coord{X: 7, Y: 2},
			steps:         5,
			endPosition:   adventofcode2021.Coord{X: 25, Y: 0},
		},
		"six": {
			startVelocity: adventofcode2021.Coord{X: 7, Y: 2},
			steps:         6,
			endPosition:   adventofcode2021.Coord{X: 27, Y: -3},
		},
		"seven": {
			startVelocity: adventofcode2021.Coord{X: 7, Y: 2},
			steps:         7,
			endPosition:   adventofcode2021.Coord{X: 28, Y: -7},
			inTarget:      true,
		},
		"nine": {
			startVelocity: adventofcode2021.Coord{X: 6, Y: 3},
			steps:         9,
			endPosition:   adventofcode2021.Coord{X: 21, Y: -9},
			inTarget:      true,
		},
		"nineZero1step": {
			startVelocity: adventofcode2021.Coord{X: 9, Y: 0},
			steps:         1,
			endPosition:   adventofcode2021.Coord{X: 9, Y: 0},
		},
		"nineZero2step": {
			startVelocity: adventofcode2021.Coord{X: 9, Y: 0},
			steps:         2,
			endPosition:   adventofcode2021.Coord{X: 17, Y: -1},
		},
		"nineZero4step": {
			startVelocity: adventofcode2021.Coord{X: 9, Y: 0},
			steps:         4,
			endPosition:   adventofcode2021.Coord{X: 30, Y: -6},
			inTarget:      true,
		},
	}

	// Testing
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			grid := &trench{
				targetTopLeft:     adventofcode2021.Coord{X: 20, Y: -5},
				targetBottomRight: adventofcode2021.Coord{X: 30, Y: -10},
				probePosition:     test.startPosition,
				probeVelocity:     test.startVelocity,
			}
			for k := 0; k < test.steps; k++ {
				grid.Step()
			}
			assert.Equal(t, test.endPosition, grid.probePosition)
			assert.Equal(t, test.inTarget, grid.InTarget())
		})
	}
}
