package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountIncrease(t *testing.T) {
	tests := map[string]struct {
		input                 string
		expectedIncreaseCount int
	}{
		"ex1": {
			input: `199
200
208
210
200
207
240
269
260
263`,
			expectedIncreaseCount: 7,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			increase := increase(test.input)
			assert.Equal(t, test.expectedIncreaseCount, increase)
		})
	}
}

func TestCountIncreaseSliding(t *testing.T) {
	tests := map[string]struct {
		input                 string
		expectedIncreaseCount int
	}{
		"ex1": {
			input: `199
200
208
210
200
207
240
269
260
263`,
			expectedIncreaseCount: 5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			increase := increaseSlidingWindow(test.input)
			assert.Equal(t, test.expectedIncreaseCount, increase)
		})
	}
}
