package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleFold(t *testing.T) {
	tests := map[string]struct {
		input    string
		instr    int
		expected int
	}{
		"onefold": {
			input: `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`,
			instr:    0,
			expected: 17,
		},
		"allfold": {
			input: `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`,
			instr:    -1,
			expected: 16,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			grid := fold(test.input, test.instr)
			answer := 0
			for _, ok := range grid {
				if ok {
					answer++
				}
			}
			assert.Equal(t, test.expected, answer)
		})
	}
}
