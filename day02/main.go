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
	input := adventofcode2021.MustInputFromWebsite("2")
	fmt.Println(commands(input))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("2")
	fmt.Println(commandsWithAim(input))
}

func commands(input string) int {
	instructions := adventofcode2021.MustStringList(input, "\n")
	depth := 0
	horizontal := 0
	for _, i := range instructions {
		var cmd string
		var dist int
		_, err := fmt.Sscanf(i, "%s %d", &cmd, &dist)
		if err != nil {
			panic(err)
		}
		switch cmd {
		case "forward":
			horizontal += dist
		case "down":
			depth += dist
		case "up":
			depth -= dist
		default:
			panic("unknown: " + cmd)
		}
	}

	return depth * horizontal
}

func commandsWithAim(input string) int {
	instructions := adventofcode2021.MustStringList(input, "\n")
	depth := 0
	horizontal := 0
	aim := 0
	for _, i := range instructions {
		var cmd string
		var dist int
		_, err := fmt.Sscanf(i, "%s %d", &cmd, &dist)
		if err != nil {
			panic(err)
		}
		switch cmd {
		case "forward":
			horizontal += dist
			depth = depth + (aim * dist)
		case "down":
			aim += dist
		case "up":
			aim -= dist
		default:
			panic("unknown: " + cmd)
		}
	}

	return depth * horizontal
}
