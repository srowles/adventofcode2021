package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCorruptLine(t *testing.T) {
	tests := map[string]struct {
		input                string
		expectedExpectedRune rune
		expectedFoundRune    rune
	}{
		"1": {
			input:                `{([(<{}[<>[]}>{[]{[(<()>`,
			expectedExpectedRune: ']',
			expectedFoundRune:    '}',
		},
		"2": {
			input:                `[[<[([]))<([[{}[[()]]]`,
			expectedExpectedRune: ']',
			expectedFoundRune:    ')',
		},
		"3": {
			input:                `[{[{({}]{}}([{[{{{}}([]`,
			expectedExpectedRune: ')',
			expectedFoundRune:    ']',
		},
		"4": {
			input:                `[<(<(<(<{}))><([]([]()`,
			expectedExpectedRune: '>',
			expectedFoundRune:    ')',
		},
		"5": {
			input:                `<{([([[(<>()){}]>(<<{{`,
			expectedExpectedRune: ']',
			expectedFoundRune:    '>',
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			expected, found := checkLine(test.input)
			assert.Equal(t, test.expectedExpectedRune, expected, string(expected))
			assert.Equal(t, test.expectedFoundRune, found, string(found))
		})
	}
}

func TestFindCorruptionScore(t *testing.T) {
	tests := map[string]struct {
		input         string
		expectedScore int
	}{
		"ex1": {
			input: `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`,
			expectedScore: 26397,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			score := findCorruptionScore(test.input)
			assert.Equal(t, test.expectedScore, score)
		})
	}
}

func TestFixLinesScore(t *testing.T) {
	tests := map[string]struct {
		input         string
		expectedScore int
	}{
		"ex1": {
			input: `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`,
			expectedScore: 288957,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			score := fixLines(test.input)
			assert.Equal(t, test.expectedScore, score)
		})
	}
}

func TestFixLine(t *testing.T) {
	tests := map[string]struct {
		input  string
		result string
		score  int
	}{
		"ex1": {
			input:  "[({(<(())[]>[[{[]{<()<>>",
			result: "}}]])})]",
			score:  288957,
		},
		"ex2": {
			input:  "[(()[<>])]({[<{<<[]>>(",
			result: ")}>]})",
			score:  5566,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			answer := fixLine(test.input)
			assert.Equal(t, test.result, answer)
			score := score(answer)
			assert.Equal(t, test.score, score)
		})
	}
}
