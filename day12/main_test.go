package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPaths(t *testing.T) {
	tests := map[string]struct {
		input         string
		expectedPaths int
		maxVisits     int
	}{
		"ex1": {
			input: `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
			expectedPaths: 10,
			maxVisits:     1,
		},
		"ex2": {
			input: `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
			expectedPaths: 36,
			maxVisits:     2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			paths := countPaths(test.input, test.maxVisits)
			assert.Equal(t, test.expectedPaths, paths)
		})
	}
}
