package main

import (
	"fmt"
	"strconv"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("3")
	fmt.Println(powerConsumption(input))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("3")
	fmt.Println(lifeSupport(input))
}

type counter struct {
	zeros, ones int
}

func lifeSupport(input string) int {
	return oxygen(input) * co2(input)
}

func oxygen(input string) int {
	numbers := adventofcode2021.MustStringList(input)
	for i := 0; i < len(numbers[0]); i++ {
		counts := counter{}
		for _, n := range numbers {
			switch n[i] {
			case '0':
				counts = counter{
					zeros: counts.zeros + 1,
					ones:  counts.ones,
				}
			case '1':
				counts = counter{
					zeros: counts.zeros,
					ones:  counts.ones + 1,
				}
			}
		}
		var findChar byte
		if counts.ones < counts.zeros {
			findChar = '0'
		} else {
			findChar = '1'
		}

		newNumbers := make([]string, 0, len(numbers))
		for _, n := range numbers {
			if n[i] == findChar {
				newNumbers = append(newNumbers, n)
			}
		}
		numbers = newNumbers
		if len(numbers) == 1 {
			result, err := strconv.ParseInt(numbers[0], 2, 64)
			if err != nil {
				panic(err)
			}
			return int(result)
		}
	}

	return 0
}

func co2(input string) int {
	numbers := adventofcode2021.MustStringList(input)
	for i := 0; i < len(numbers[0]); i++ {
		counts := counter{}
		for _, n := range numbers {
			switch n[i] {
			case '0':
				counts = counter{
					zeros: counts.zeros + 1,
					ones:  counts.ones,
				}
			case '1':
				counts = counter{
					zeros: counts.zeros,
					ones:  counts.ones + 1,
				}
			}
		}
		var findChar byte
		if counts.ones > counts.zeros {
			findChar = '0'
		} else if counts.ones == counts.zeros {
			findChar = '0'
		} else {
			findChar = '1'
		}

		newNumbers := make([]string, 0, len(numbers))
		for _, n := range numbers {
			if n[i] == findChar {
				newNumbers = append(newNumbers, n)
			}
		}
		numbers = newNumbers
		if len(numbers) == 1 {
			result, err := strconv.ParseInt(numbers[0], 2, 64)
			if err != nil {
				panic(err)
			}
			return int(result)
		}
	}

	return 0
}

func powerConsumption(input string) int {
	numbers := adventofcode2021.MustStringList(input)
	counts := make([]counter, len(numbers[0]))
	for _, n := range numbers {
		for i, b := range n {
			switch b {
			case '0':
				counts[i] = counter{
					zeros: counts[i].zeros + 1,
					ones:  counts[i].ones,
				}
			case '1':
				counts[i] = counter{
					zeros: counts[i].zeros,
					ones:  counts[i].ones + 1,
				}
			}
		}
	}

	gamma := make([]int, len(numbers[0]))
	epsilon := make([]int, len(numbers[0]))
	for i, c := range counts {
		if c.ones > c.zeros {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}
	gammaRate := number(gamma)
	epsionRate := number(epsilon)
	return gammaRate * epsionRate
}

func number(numbers []int) int {
	var result int
	for i, n := range numbers {
		result |= n << (len(numbers) - i - 1)
	}
	return result
}
