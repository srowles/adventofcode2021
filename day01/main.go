package main

import (
	"fmt"
	"math"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("1")
	fmt.Println(increase(input))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("1")
	fmt.Println(increaseSlidingWindow(input))
}

func increase(input string) int {
	readings := adventofcode2021.MustIntList(input)
	previous := math.MaxInt64
	increaseCount := 0
	for _, v := range readings {
		if v > previous {
			increaseCount++
		}
		previous = v
	}

	return increaseCount
}

func increaseSlidingWindow(input string) int {
	readings := adventofcode2021.MustIntList(input)
	previous := math.MaxInt64
	increaseCount := 0
	for i := 0; i < len(readings); i++ {
		if i+3 > len(readings) {
			break
		}
		s := sum(readings[i : i+3])
		if s > previous {
			increaseCount++
		}
		previous = s
	}

	return increaseCount
}

func sum(vals []int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}
