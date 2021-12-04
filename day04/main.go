package main

import (
	"fmt"
	"strings"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("4")
	fmt.Println(winningBoardSum(input))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("4")
	fmt.Println(lastWinningBoardSum(input))
}

func winningBoardSum(input string) int {
	numbers, cards := parse(input)
	var winner card
	var lastCalled int
outer:
	for _, n := range numbers {
		for _, c := range cards {
			c.markNumber(n)
			if c.hasWon() {
				winner = *c
				lastCalled = n
				break outer
			}
		}
	}

	return winner.unmarkedSum() * lastCalled
}

func lastWinningBoardSum(input string) int {
	numbers, cards := parse(input)
	var winner card
	var lastCalled int

	for _, n := range numbers {
		for _, c := range cards {
			c.markNumber(n)
			if c.hasWon() {
				winner = *c
				lastCalled = n
			}
		}
		// remove cards that have won
		var newCards []*card
		for _, c := range cards {
			if !c.hasWon() {
				newCards = append(newCards, c)
			}
		}
		cards = newCards
	}

	return winner.unmarkedSum() * lastCalled
}

func parse(input string) ([]int, []*card) {
	var cards []*card
	var numbers []int
	var lines []string
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		line := strings.TrimSpace(line)
		if len(numbers) == 0 {
			numbers = adventofcode2021.MustIntCommaList(line)
			continue
		}
		if line == "" && len(lines) > 0 {
			cards = append(cards, parseCard(strings.Join(lines, "\n")))
			lines = make([]string, 0)
			continue
		}
		lines = append(lines, line)
	}

	cards = append(cards, parseCard(strings.Join(lines, "\n")))

	return numbers, cards
}

type card struct {
	rows   [][]int
	marked map[adventofcode2021.Coord]bool
}

// 22 13 17 11  0
//  8  2 23  4 24
// 21  9 14 16  7
//  6 10  3 18  5
//  1 12 20 15 19
func parseCard(input string) *card {
	card := &card{
		marked: make(map[adventofcode2021.Coord]bool),
	}
	for _, r := range strings.Split(strings.TrimSpace(input), "\n") {
		row := make([]int, 5)
		_, err := fmt.Sscanf(r, "%d %d %d %d %d", &row[0], &row[1], &row[2], &row[3], &row[4])
		if err != nil {
			panic(err.Error())
		}
		card.rows = append(card.rows, row)
	}
	return card
}

func (c *card) markNumber(number int) {
	for y, row := range c.rows {
		for x, num := range row {
			if num == number {
				c.marked[adventofcode2021.Coord{X: x, Y: y}] = true
			}
		}
	}
}

// full row or column == win
func (c *card) hasWon() bool {
	for _, win := range wins {
		count := 0
		for coord := range win {
			if c.marked[coord] {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}

	return false
}

func (c *card) unmarkedSum() int {
	marked := make(map[int]bool)
	for coord := range c.marked {
		marked[c.rows[coord.Y][coord.X]] = true
	}

	result := 0
	for _, row := range c.rows {
		for _, num := range row {
			if !marked[num] {
				result += num
			}
		}
	}
	return result
}

var wins []map[adventofcode2021.Coord]bool

// make winning row/column maps
func init() {
	// rows
	for y := 0; y < 5; y++ {
		row := make(map[adventofcode2021.Coord]bool)
		col := make(map[adventofcode2021.Coord]bool)
		for x := 0; x < 5; x++ {
			row[adventofcode2021.Coord{X: x, Y: y}] = true
			col[adventofcode2021.Coord{X: y, Y: x}] = true
		}
		wins = append(wins, row)
		wins = append(wins, col)
	}
}
