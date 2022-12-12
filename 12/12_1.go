package main

import (
	"container/heap"
	"math"
	"strings"
)

func SolveCurrentDay(input string) int {
	heightMap, start, end := ParseInput(input)

	return Dijkstra(heightMap, [][2]int{start}, end)
}

func ParseInput(input string) ([][]int, [2]int, [2]int) {
	heightmap := [][]int{}
	start := [2]int{}
	end := [2]int{}

	for y, line := range strings.Split(input, "\n") {
		heightline := []int{}

		for x, char := range line {
			switch char {
			case 'S':
				// Start
				start = [2]int{x, y}
				heightline = append(heightline, 1)
			case 'E':
				// End
				end = [2]int{x, y}
				heightline = append(heightline, 26)
			default:
				// ascii a == 97
				// heightmap a == 1
				height := int(char - 96)
				heightline = append(heightline, height)
			}
		}

		heightmap = append(heightmap, heightline)
	}

	return heightmap, start, end
}

func Dijkstra(heightMap [][]int, start [][2]int, end [2]int) int {
	// create vertex set Q
	Q := make(PriorityQueue, 0)

	dist := map[[2]int]int{}

	maxX := len(heightMap[0]) - 1
	maxY := len(heightMap) - 1

	// for each vertex v in Graph:
	for y, row := range heightMap {
		for x := range row {
			v := [2]int{x, y}
			// dist[v] ← INFINITY
			dist[v] = math.MaxInt

			// add v to Q
			heap.Push(&Q, &Item{
				value:    v,
				priority: dist[v],
				index:    -1,
			})
		}
	}

	for _, v := range start {
		dist[v] = 0
		if vItem := Q.Find(v); vItem != nil {
			Q.update(vItem, v, dist[v])
		}
	}

	// while Q is not empty:
	for Q.Len() > 0 {
		// u ← vertex in Q with min dist[u]
		// remove u from Q
		u := heap.Pop(&Q).(*Item)

		// Exit early if the current node is equal to destination node
		if u.value == end {
			continue
		}

		// for each neighbor v of u still in Q:
		for _, v := range GetPossibleNeighbors(heightMap, u.value, maxX, maxY) {
			// alt ← dist[u] + length(u, v)
			alt := dist[u.value] + 1

			// if alt < dist[v]:
			if alt < dist[v] {
				// dist[v] ← alt
				// log.Printf("New shortest path to location %3d;%3d: %4d (previous: %4d; coming from %3d;%3d)", v[0], v[1], alt, dist[v], u.value[0], u.value[1])
				dist[v] = alt

				if vItem := Q.Find(v); vItem != nil {
					Q.update(vItem, v, dist[v])
				}
			}
		}
	}

	return dist[end]
}

func GetPossibleNeighbors(heightMap [][]int, coords [2]int, maxX, maxY int) [][2]int {
	neighborOffsets := [][2]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	currentHeight := heightMap[coords[1]][coords[0]]

	neighbors := [][2]int{}

	for _, offset := range neighborOffsets {
		newX := coords[0] + offset[0]
		newY := coords[1] + offset[1]

		if newX < 0 || newX > maxX || newY < 0 || newY > maxY {
			continue
		}

		neighborHeight := heightMap[newY][newX]
		heightDifference := neighborHeight - currentHeight

		if heightDifference <= 1 {
			neighbors = append(neighbors, [2]int{newX, newY})
		}
	}

	return neighbors
}
