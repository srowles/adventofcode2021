package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		input        string
		expectedRisk int
	}{
		"ex1": {
			input: `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`,
			expectedRisk: 40,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			risk := lowestRisk(test.input)
			assert.Equal(t, test.expectedRisk, risk)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		input        string
		expectedRisk int
	}{
		"ex1": {
			input: `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`,
			expectedRisk: 315,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			risk := lowestRiskFiveTimes(test.input)
			assert.Equal(t, test.expectedRisk, risk)
		})
	}
}
