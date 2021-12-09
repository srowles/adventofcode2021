package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowPoints(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"ex1": {
			input: `2199943210
3987894921
9856789892
8767896789
9899965678`,
			expected: 15,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getRisk(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestBasins(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"ex1": {
			input: `2199943210
3987894921
9856789892
8767896789
9899965678`,
			expected: 1134,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getBasins(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
