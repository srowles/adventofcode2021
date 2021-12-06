package main

import (
	"testing"

	"github.com/srowles/adventofcode2021"
	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	tests := map[string]struct {
		startGrid    grid
		start, end   adventofcode2021.Coord
		expectedGrid grid
	}{
		"l1": {
			startGrid: grid{points: make(map[adventofcode2021.Coord]int)},
			start:     adventofcode2021.Coord{X: 2, Y: 2},
			end:       adventofcode2021.Coord{X: 2, Y: 1},
			expectedGrid: grid{
				points: map[adventofcode2021.Coord]int{
					{X: 2, Y: 2}: 1,
					{X: 2, Y: 1}: 1,
				},
			},
		},
		"l2": {
			startGrid: grid{points: make(map[adventofcode2021.Coord]int)},
			start:     adventofcode2021.Coord{X: 1, Y: 1},
			end:       adventofcode2021.Coord{X: 1, Y: 3},
			expectedGrid: grid{
				points: map[adventofcode2021.Coord]int{
					{X: 1, Y: 1}: 1,
					{X: 1, Y: 2}: 1,
					{X: 1, Y: 3}: 1,
				},
			},
		},
		"diag1": {
			startGrid: grid{points: make(map[adventofcode2021.Coord]int)},
			start:     adventofcode2021.Coord{X: 1, Y: 1},
			end:       adventofcode2021.Coord{X: 3, Y: 3},
			expectedGrid: grid{
				points: map[adventofcode2021.Coord]int{
					{X: 1, Y: 1}: 1,
					{X: 2, Y: 2}: 1,
					{X: 3, Y: 3}: 1,
				},
			},
		},
		"diag2": {
			startGrid: grid{points: make(map[adventofcode2021.Coord]int)},
			start:     adventofcode2021.Coord{X: 5, Y: 7},
			end:       adventofcode2021.Coord{X: 7, Y: 5},
			expectedGrid: grid{
				points: map[adventofcode2021.Coord]int{
					{X: 5, Y: 7}: 1,
					{X: 6, Y: 6}: 1,
					{X: 7, Y: 5}: 1,
				},
			},
		},
		"diag3": {
			startGrid: grid{points: make(map[adventofcode2021.Coord]int)},
			start:     adventofcode2021.Coord{X: 5, Y: 5},
			end:       adventofcode2021.Coord{X: 8, Y: 2},
			expectedGrid: grid{
				points: map[adventofcode2021.Coord]int{
					{X: 5, Y: 5}: 1,
					{X: 6, Y: 4}: 1,
					{X: 7, Y: 3}: 1,
					{X: 8, Y: 2}: 1,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			addLine(test.startGrid, test.start, test.end, true)
			assert.Equal(t, test.expectedGrid, test.startGrid)
		})
	}
}

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"card1": {
			input: `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
			expected: 5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := countOverlaps(test.input, false)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"card1": {
			input: `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
			expected: 12,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := countOverlaps(test.input, true)
			assert.Equal(t, test.expected, actual)
		})
	}
}
