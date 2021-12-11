package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountFlashes(t *testing.T) {
	tests := map[string]struct {
		input           string
		expectedFlashes int
		steps           int
	}{
		"1": {
			input: `11111
19991
19191
19991
11111`,
			expectedFlashes: 9,
			steps:           2,
		},
		"2": {
			input: `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
			expectedFlashes: 204,
			steps:           10,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			flashes := countFlashes(test.input, test.steps)
			assert.Equal(t, test.expectedFlashes, flashes)
		})
	}
}

func TestAllFlashes(t *testing.T) {
	tests := map[string]struct {
		input           string
		expectedFlashes int
	}{
		"ex": {
			input: `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
			expectedFlashes: 195,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			flashes := findAllStep(test.input)
			assert.Equal(t, test.expectedFlashes, flashes)
		})
	}
}
