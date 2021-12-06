package main

import (
	"fmt"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("6")
	initial := adventofcode2021.MustIntCommaList(input)
	count := spawnTrack(initial, 80)
	fmt.Println(count)
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("6")
	initial := adventofcode2021.MustIntCommaList(input)
	count := spawnTrack(initial, 256)
	fmt.Println(count)
}

func spawnTrack(fish []int, days int) int {
	counter := make([]int, 9)
	for _, f := range fish {
		counter[f] = counter[f] + 1
	}

	for d := 0; d < days; d++ {
		newCounter := make([]int, 9)
		for i := 8; i > 0; i-- {
			newCounter[i-1] = counter[i]
		}
		newCounter[8] = counter[0]
		newCounter[6] += counter[0]
		counter = newCounter
	}

	count := 0
	for _, c := range counter {
		count += c
	}
	return count
}
