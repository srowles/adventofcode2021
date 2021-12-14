package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInserts(t *testing.T) {
	tests := map[string]struct {
		input                   string
		steps                   int
		expected                string
		mostCommon, leastCommon string
	}{
		"twice": {
			input: `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
			steps:       2,
			expected:    "NBCCNBBBCBHCB",
			mostCommon:  "B",
			leastCommon: "H",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			chain := insert(test.input, test.steps)
			assert.Equal(t, test.expected, chain.String())
			least, _ := chain.Least()
			most, _ := chain.Most()
			assert.Equal(t, test.leastCommon, least)
			assert.Equal(t, test.mostCommon, most)
		})
	}
}

func TestInsertsNoOutputCheck(t *testing.T) {
	tests := map[string]struct {
		input                             string
		steps                             int
		mostCommon, leastCommon           string
		mostCommonCount, leastCommonCount int
	}{
		"two": {
			input: `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
			steps:            2,
			mostCommon:       "B",
			mostCommonCount:  6,
			leastCommon:      "H",
			leastCommonCount: 1,
		},
		"three": {
			input: `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
			steps:            3,
			mostCommon:       "B",
			mostCommonCount:  11,
			leastCommon:      "H",
			leastCommonCount: 4,
		},
		"ten": {
			input: `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
			steps:            10,
			mostCommon:       "B",
			mostCommonCount:  1749,
			leastCommon:      "H",
			leastCommonCount: 161,
		},
		"40": {
			input: `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
			steps:            40,
			mostCommon:       "B",
			mostCommonCount:  2192039569602,
			leastCommon:      "H",
			leastCommonCount: 3849876073,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			chain := noChainInsert(test.input, test.steps)
			least, lastCount := chain.Least()
			most, mostCount := chain.Most()
			assert.Equal(t, test.leastCommon, least)
			assert.Equal(t, test.mostCommon, most)
			assert.Equal(t, test.leastCommonCount, lastCount)
			assert.Equal(t, test.mostCommonCount, mostCount)
		})
	}
}
