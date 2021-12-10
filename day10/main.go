package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("10")
	start := time.Now().UTC()
	answer := findCorruptionScore(input)
	fmt.Println(answer)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("10")
	start := time.Now().UTC()
	answer := fixLines(input)
	fmt.Println(answer)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func fixLines(input string) int {
	var scores []int
	for _, line := range adventofcode2021.MustStringList(input, "\n") {
		corruption := findCorruptionScore(line)
		if corruption != 0 {
			continue
		}
		scores = append(scores, score(fixLine(line)))
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func score(input string) int {
	score := 0
	for _, r := range input {
		score *= 5
		score += value2[r]
	}
	return score
}

func fixLine(line string) string {
	s := adventofcode2021.Stack{}
	for _, r := range line {
		if _, ok := openToClose[r]; ok {
			s.Push(r)
		} else {
			_ = s.Pop()
		}
	}
	var fixes []string
	for !s.IsEmpty() {
		o := s.Pop()
		fixes = append(fixes, string(openToClose[o]))
	}

	return strings.Join(fixes, "")
}

func findCorruptionScore(input string) int {
	score := 0
	for _, line := range adventofcode2021.MustStringList(input, "\n") {
		_, e := checkLine(line)
		if v, ok := value[e]; ok {
			score += v
		}
	}
	return score
}

func checkLine(line string) (rune, rune) {
	s := adventofcode2021.Stack{}
	for _, r := range line {
		if _, ok := openToClose[r]; ok {
			s.Push(r)
		} else {
			previous := s.Pop()
			if previous != closeToOpen[r] {
				return openToClose[previous], r
			}
		}
	}

	return '_', '_'
}

var closeToOpen = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
	'>': '<',
}

var openToClose = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
	'<': '>',
}

var value = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var value2 = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}
