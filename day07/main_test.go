package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestFuel(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"ex1": {
			input:    []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expected: 37,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := smallestFuel(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSmallestFuel2(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"ex1": {
			input:    []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expected: 168,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := smallestFuel2(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
