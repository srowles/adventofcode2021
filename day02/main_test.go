package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedAnswer int
	}{
		"ex1": {
			input: `forward 5
down 5
forward 8
up 3
down 8
forward 2`,
			expectedAnswer: 150,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := commands(test.input)
			assert.Equal(t, test.expectedAnswer, answer)
		})
	}
}

func TestCommandsWithAim(t *testing.T) {
	tests := map[string]struct {
		input          string
		expectedAnswer int
	}{
		"ex1": {
			input: `forward 5
down 5
forward 8
up 3
down 8
forward 2`,
			expectedAnswer: 900,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := commandsWithAim(test.input)
			assert.Equal(t, test.expectedAnswer, answer)
		})
	}
}
