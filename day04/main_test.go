package main

import (
	"testing"

	"github.com/srowles/adventofcode2021"
	"github.com/stretchr/testify/assert"
)

func TestParseCard(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected *card
	}{
		"card1": {
			input: `22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19`,
			expected: &card{
				rows: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				marked: make(map[adventofcode2021.Coord]bool),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := parseCard(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestNumberMarkCard(t *testing.T) {
	tests := map[string]struct {
		input    *card
		numbers  []int
		expected map[adventofcode2021.Coord]bool
	}{
		"card1": {
			input: &card{
				marked: make(map[adventofcode2021.Coord]bool),
				rows: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
			},
			numbers: []int{22, 2, 5, 19},
			expected: map[adventofcode2021.Coord]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 1}: true,
				{X: 4, Y: 3}: true,
				{X: 4, Y: 4}: true,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for _, n := range test.numbers {
				test.input.markNumber(n)
			}
			assert.Equal(t, test.expected, test.input.marked)
		})
	}
}

func TestHasWon(t *testing.T) {
	tests := map[string]struct {
		input       *card
		numbers     []int
		expectedWin bool
	}{
		"noWin": {
			input: &card{
				marked: make(map[adventofcode2021.Coord]bool),
				rows: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
			},
			numbers:     []int{22, 2, 5, 19},
			expectedWin: false,
		},
		"rowWin": {
			input: &card{
				marked: make(map[adventofcode2021.Coord]bool),
				rows: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
			},
			numbers:     []int{21, 9, 14, 16, 7},
			expectedWin: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for _, n := range test.numbers {
				test.input.markNumber(n)
			}
			assert.Equal(t, test.expectedWin, test.input.hasWon())
		})
	}
}

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input           string
		expectedNumbers []int
		expectedCards   []*card
	}{
		"ex1": {
			input: `
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`,
			expectedNumbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			expectedCards: []*card{
				{
					marked: make(map[adventofcode2021.Coord]bool),
					rows: [][]int{
						{22, 13, 17, 11, 0},
						{8, 2, 23, 4, 24},
						{21, 9, 14, 16, 7},
						{6, 10, 3, 18, 5},
						{1, 12, 20, 15, 19},
					},
				},
				{
					marked: make(map[adventofcode2021.Coord]bool),
					rows: [][]int{
						{3, 15, 0, 2, 22},
						{9, 18, 13, 17, 5},
						{19, 8, 7, 25, 23},
						{20, 11, 10, 24, 4},
						{14, 21, 16, 12, 6},
					},
				},
				{
					marked: make(map[adventofcode2021.Coord]bool),
					rows: [][]int{
						{14, 21, 17, 24, 4},
						{10, 16, 15, 9, 19},
						{18, 8, 23, 26, 20},
						{22, 11, 13, 6, 5},
						{2, 0, 12, 3, 7},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			numbers, cards := parse(test.input)
			assert.Equal(t, test.expectedNumbers, numbers)
			assert.Equal(t, test.expectedCards, cards)
		})
	}
}

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"ex1": {
			input: `
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`,
			expected: 4512,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := winningBoardSum(test.input)
			assert.Equal(t, test.expected, answer)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"ex1": {
			input: `
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`,
			expected: 1924,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := lastWinningBoardSum(test.input)
			assert.Equal(t, test.expected, answer)
		})
	}
}
