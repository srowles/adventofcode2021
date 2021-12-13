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
	input := adventofcode2021.MustInputFromWebsite("12")
	start := time.Now().UTC()
	answer := countPaths(input, 1)
	fmt.Println(answer)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("12")
	start := time.Now().UTC()
	answer := countPaths(input, 2)
	fmt.Println(answer)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

type graph struct {
	lookup map[string]*node
}

type node struct {
	name        string
	connections map[string]*node
}

func (n *node) String() string {
	var connections []string
	for c := range n.connections {
		connections = append(connections, c)
	}
	sort.Strings(connections)

	return fmt.Sprintf("%s: %s", n.name, strings.Join(connections, ","))
}

func countPaths(input string, maxVisits int) int {
	g := createGraph(input)

	paths := g.FindPaths(g.lookup["start"], []string{}, maxVisits)
	// fmt.Println(paths)
	return len(paths)
}

func alreadyVisited(path []string, name string, maxVisits int) int {
	visits := 0
	for _, p := range path {
		if p == name {
			visits++
		}
	}
	return visits
}

func (g *graph) FindPaths(startNode *node, path []string, maxVisits int) [][]string {
	if startNode.name[0] > 96 {
		count := alreadyVisited(path, startNode.name, maxVisits)
		if startNode.name == "start" && count > 0 {
			return nil
		} else if count >= maxVisits {
			return nil
		}

		if count == 1 {
			maxVisits = 1
		}
	}
	var paths [][]string
	currentPath := append(path, startNode.name)

	if startNode.name == "end" {
		return [][]string{currentPath}
	}

	for _, node := range startNode.connections {
		subPath := g.FindPaths(node, currentPath, maxVisits)
		if subPath != nil {
			paths = append(paths, subPath...)
		}
	}

	return paths
}

func createGraph(input string) *graph {
	g := &graph{
		lookup: make(map[string]*node),
	}
	for _, line := range adventofcode2021.MustStringList(input, "\n") {
		parts := strings.Split(line, "-")
		s := parts[0]
		e := parts[1]
		n := getOrCreateNode(g, s)
		c := getOrCreateNode(g, e)
		n.connections[e] = c
		c.connections[s] = n
	}

	return g
}

func getOrCreateNode(g *graph, name string) *node {
	n, ok := g.lookup[name]
	if !ok {
		n = &node{
			name:        name,
			connections: make(map[string]*node),
		}
		g.lookup[name] = n
		return n
	}

	return n
}
