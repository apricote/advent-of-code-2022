package main

func SolveCurrentDayWithTwist(input string) int {
	heightMap, _, end := ParseInput(input)

	startingLocations := [][2]int{}

	for y, row := range heightMap {
		for x, height := range row {
			if height == 1 {
				startingLocations = append(startingLocations, [2]int{x, y})
			}
		}
	}

	return Dijkstra(heightMap, startingLocations, end)
}
