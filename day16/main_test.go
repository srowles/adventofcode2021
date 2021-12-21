package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceToInt(t *testing.T) {
	tests := map[string]struct {
		input    []uint64
		expected uint64
	}{
		"6": {
			input:    []uint64{1, 1, 0},
			expected: 6,
		},
		"4": {
			input:    []uint64{1, 0, 0},
			expected: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := sliceToInt(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestDecodeBinary(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []uint64
	}{
		"decode1": {
			input:    "D2FE28",
			expected: []uint64{1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := binaryDecode(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		input              string
		expectedVersionSum int
	}{
		"simple": {
			input:              `D2FE28`,
			expectedVersionSum: 6,
		},
		"twoPackets": {
			input:              "38006F45291200",
			expectedVersionSum: 9,
		},
		"twoPacketsSubType1": {
			input:              "EE00D40C823060",
			expectedVersionSum: 14,
		},
		"ex1": {
			input:              `8A004A801A8002F478`,
			expectedVersionSum: 16,
		},
		"ex2": {
			input:              `620080001611562C8802118E34`,
			expectedVersionSum: 12,
		},
		"ex3": {
			input:              `C0015000016115A2E0802F182340`,
			expectedVersionSum: 23,
		},
		"ex4": {
			input:              `A0016C880162017C3686B18A3D4780`,
			expectedVersionSum: 31,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			risk := versionSum(test.input)
			assert.Equal(t, test.expectedVersionSum, risk)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		input       string
		expectedSum int
	}{
		"1": {
			input:       `C200B40A82`,
			expectedSum: 3,
		},
		"2": {
			input:       `04005AC33890`,
			expectedSum: 54,
		},
		"3": {
			input:       `880086C3E88112`,
			expectedSum: 7,
		},
		"4": {
			input:       `CE00C43D881120`,
			expectedSum: 9,
		},
		"5": {
			input:       `D8005AC2A8F0`,
			expectedSum: 1,
		},
		"6": {
			input:       `F600BC2D8F`,
			expectedSum: 0,
		},
		"7": {
			input:       `9C005AC2F8F0`,
			expectedSum: 0,
		},
		"8": {
			input:       `9C0141080250320F1802104A08`,
			expectedSum: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			risk := value(test.input)
			assert.Equal(t, test.expectedSum, risk)
		})
	}
}
