package main

import (
	"container/list"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("14")
	start := time.Now().UTC()
	chain := insert(input, 10)
	_, most := chain.Most()
	_, least := chain.Least()
	fmt.Println(most - least)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("14")
	start := time.Now().UTC()
	chain := noChainInsert(input, 40)
	_, most := chain.Most()
	_, least := chain.Least()
	fmt.Println(most - least)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

type chain struct {
	list   *list.List
	counts map[string]int
	rules  map[string]string
	pairs  map[string]int
}

func (c chain) Debug() {
	fmt.Println("chain:", c.list)
	fmt.Println("counts:", c.counts)
}

func (c chain) String() string {
	var result strings.Builder
	for e := c.list.Front(); e != nil; e = e.Next() {
		result.WriteString(e.Value.(string))
	}
	return result.String()
}

func (c chain) Least() (string, int) {
	smallestCount := math.MaxInt64
	letter := "unknown"
	for l, c := range c.counts {
		if c < smallestCount {
			smallestCount = c
			letter = l
		}
	}
	return letter, smallestCount
}

func (c chain) Most() (string, int) {
	largestCount := 0
	letter := "unknown"
	for l, c := range c.counts {
		if c > largestCount {
			largestCount = c
			letter = l
		}
	}
	return letter, largestCount
}

func insert(input string, steps int) chain {
	c := chain{
		list:   list.New(),
		counts: make(map[string]int),
		rules:  make(map[string]string),
	}

	for _, line := range adventofcode2021.MustStringList(input, "\n") {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		if len(parts) == 1 {
			for _, element := range line {
				c.list.PushBack(string(element))
				c.counts[string(element)] = c.counts[string(element)] + 1
			}
			continue
		}
		c.rules[parts[0]] = parts[1]
	}

	for s := 0; s < steps; s++ {
		for e := c.list.Front(); e != nil && e.Next() != nil; e = e.Next() {
			input := e.Value.(string) + e.Next().Value.(string)
			if replace, ok := c.rules[input]; ok {
				c.counts[replace] = c.counts[replace] + 1
				c.list.InsertAfter(replace, e)
				e = e.Next()
			}
		}
	}

	return c
}

func noChainInsert(input string, steps int) chain {
	c := chain{
		counts: make(map[string]int),
		rules:  make(map[string]string),
		pairs:  make(map[string]int),
	}

	for _, line := range adventofcode2021.MustStringList(input, "\n") {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		if len(parts) == 1 {
			for i := 0; i < len(line)-1; i++ {
				c.pairs[string(line[i])+string(line[i+1])] = c.pairs[string(line[i])+string(line[i+1])] + 1
				c.counts[string(line[i])] = c.counts[string(line[i])] + 1
				c.counts[string(line[i+1])] = c.counts[string(line[i+1])] + 1
			}
			continue
		}
		c.rules[parts[0]] = parts[1]
	}

	for i := 0; i < steps; i++ {
		next := make(map[string]int)
		for pair, count := range c.pairs {
			new, ok := c.rules[pair]
			if !ok {
				next[pair] = count
				continue
			}
			next[string(pair[0])+new] = next[string(pair[0])+new] + count
			next[new+string(pair[1])] = next[new+string(pair[1])] + count
			c.counts[new] = c.counts[new] + count
		}
		c.pairs = next
	}

	return c
}
