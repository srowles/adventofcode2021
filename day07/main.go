package main

import (
	"fmt"
	"math"
	"time"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("7")
	start := time.Now().UTC()
	initial := adventofcode2021.MustIntCommaList(input)
	count := smallestFuel(initial)
	fmt.Println(count)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("7")
	start := time.Now().UTC()
	initial := adventofcode2021.MustIntCommaList(input)
	count := smallestFuel2(initial)
	fmt.Println(count)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func smallestFuel(input []int) int {
	maxPos := 0
	minPos := math.MaxInt64
	for _, p := range input {
		if p > maxPos {
			maxPos = p
		}
		if p < minPos {
			minPos = p
		}
	}

	minFuel := math.MaxInt64
	for p := minPos; p <= maxPos; p++ {
		fuel := 0
		for _, i := range input {
			f := adventofcode2021.Abs(i - p)
			fuel += f
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}

func smallestFuel2(input []int) int {
	maxPos := 0
	minPos := math.MaxInt64
	for _, p := range input {
		if p > maxPos {
			maxPos = p
		}
		if p < minPos {
			minPos = p
		}
	}

	minFuel := math.MaxInt64
	for p := minPos; p <= maxPos; p++ {
		fuel := 0
		for _, i := range input {
			f := adventofcode2021.Abs(i - p)
			fuel += sum(f)
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}

func sum(a int) int {
	total := 0
	cost := 1
	for i := 0; i < a; i++ {
		total += cost
		cost++
	}
	return total
}
