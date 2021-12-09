package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("8")
	start := time.Now().UTC()
	count := count1478(input)
	fmt.Println(count)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("8")
	start := time.Now().UTC()
	count := fullDecode(input)
	fmt.Println(count)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func count1478(input string) int {
	count := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, " | ")
		for _, digit := range strings.Split(strings.TrimSpace(parts[1]), " ") {
			d := getDigit(digit)
			switch d {
			case 1, 4, 7, 8:
				count++
			}
		}
	}
	return count
}

func fullDecode(input string) int {
	total := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, " | ")
		signals := strings.Split(strings.TrimSpace(parts[0]), " ")
		// obvious segments first
		var one, seven, four, eight string
		var zeroSixNine []string
		for _, s := range signals {
			switch len(s) {
			case 2:
				one = s
			case 3:
				seven = s
			case 4:
				four = s
			case 5:
				// no needed
			case 6:
				zeroSixNine = append(zeroSixNine, s)
			case 7:
				eight = s
			}
		}

		// now work out segment from the letters we have
		abfg := intersect(zeroSixNine...)
		cf := one
		bcdf := four
		acf := seven
		abcdefg := eight
		a := sub(acf, cf)
		bfg := sub(abfg, a)
		g := sub(bfg, bcdf)
		c := sub(cf, bfg)
		f := sub(cf, c)
		b := sub(sub(bfg, f), g)
		d := sub(sub(bcdf, bfg), c)
		aeg := sub(abcdefg, bcdf)
		eg := sub(aeg, a)
		e := sub(eg, g)

		var number string
		digits := parts[1]
		for _, digit := range strings.Split(digits, " ") {
			segments := make(map[rune]bool)
			for _, r := range digit {
				switch r {
				case rune(a[0]):
					segments['a'] = true
				case rune(b[0]):
					segments['b'] = true
				case rune(c[0]):
					segments['c'] = true
				case rune(d[0]):
					segments['d'] = true
				case rune(e[0]):
					segments['e'] = true
				case rune(f[0]):
					segments['f'] = true
				case rune(g[0]):
					segments['g'] = true
				}
			}
			n := fromSegments(segments)
			number = number + n
		}
		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err.Error())
		}
		total += n
	}
	return total
}

func fromSegments(segments map[rune]bool) string {
	switch {
	case eq(segments, zero):
		return "0"
	case eq(segments, one):
		return "1"
	case eq(segments, two):
		return "2"
	case eq(segments, three):
		return "3"
	case eq(segments, four):
		return "4"
	case eq(segments, five):
		return "5"
	case eq(segments, six):
		return "6"
	case eq(segments, seven):
		return "7"
	case eq(segments, eight):
		return "8"
	case eq(segments, nine):
		return "9"
	}
	panic(fmt.Sprintf("segments not mapped: %v", segments))
}

func eq(a, b map[rune]bool) bool {
	count := 0
	for ar := range a {
		if b[ar] {
			count++
		} else {
			return false
		}
	}
	return count == len(b)
}

var zero, one, two, three, four, five, six, seven, eight, nine map[rune]bool

func init() {
	zero = stringToMap("abcefg")
	one = stringToMap("cf")
	two = stringToMap("acdeg")
	three = stringToMap("acdfg")
	four = stringToMap("bcdf")
	five = stringToMap("abdfg")
	six = stringToMap("abdefg")
	seven = stringToMap("acf")
	eight = stringToMap("abcdefg")
	nine = stringToMap("abcdfg")
}

func stringToMap(letters string) map[rune]bool {
	result := make(map[rune]bool)
	for _, r := range letters {
		result[r] = true
	}
	return result
}

//   0:      1:      2:      3:      4:
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....
//
//   5:      6:      7:      8:      9:
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

// sub subtracts b from a
func sub(a, b string) string {
	var result string
	for _, ac := range a {
		found := false
		for _, bc := range b {
			if ac == bc {
				found = true
				break
			}
		}
		if !found {
			result = result + string(ac)
		}
	}
	return result
}

// intersect returns all that are only in all supplied values
func intersect(wires ...string) string {
	var result string
	requiredCount := len(wires)
	for i := int('a'); i <= int('g'); i++ {
		count := 0
		for _, w := range wires {
			count += strings.Count(w, string(rune(i)))
		}
		if count == requiredCount {
			result = result + string(rune(i))
		}
	}
	return result
}

func getDigit(d string) int {
	switch len(d) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}
	return -1
}
