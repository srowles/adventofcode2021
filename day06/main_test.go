package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulateFishTrack(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"pasrt1": {
			input:    []int{3, 4, 3, 1, 2},
			expected: 5934,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := spawnTrack(test.input, 80)
			assert.Equal(t, test.expected, result)
		})
	}
}
