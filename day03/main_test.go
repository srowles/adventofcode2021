package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumbers(t *testing.T) {
	tests := map[string]struct {
		input          []int
		expectedAnswer int
	}{
		"22": {
			input:          []int{1, 0, 1, 1, 0},
			expectedAnswer: 22,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := number(test.input)
			assert.Equal(t, test.expectedAnswer, answer)
		})
	}
}

func TestPowerConsumption(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedAnswer int
	}{
		"ex1": {
			input: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
			expectedAnswer: 198,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := powerConsumption(test.input)
			assert.Equal(t, test.expectedAnswer, answer)
		})
	}
}

func TestOxygen(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedAnswer int
	}{
		"ex1": {
			input: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
			expectedAnswer: 23,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := oxygen(test.input)
			assert.Equal(t, test.expectedAnswer, answer)
		})
	}
}

func TestCO2(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedAnswer int
	}{
		"ex1": {
			input: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
			expectedAnswer: 10,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := co2(test.input)
			assert.Equal(t, test.expectedAnswer, answer)
		})
	}
}

func TestLifeSupport(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedAnswer int
	}{
		"ex1": {
			input: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
			expectedAnswer: 230,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := lifeSupport(test.input)
			assert.Equal(t, test.expectedAnswer, answer)
		})
	}
}
