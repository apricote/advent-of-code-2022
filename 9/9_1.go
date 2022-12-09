package main

import (
	"strconv"
	"strings"
)

type Knots []*[2]int

func SolveCurrentDay(input string) int {
	motions := ParseInput(input)

	knots := Knots{
		&[2]int{0, 0},
		&[2]int{0, 0},
	}

	return SimulateDaRope(knots, motions)
}

func SimulateDaRope(knots Knots, motions []Motion) int {
	visitedLocations := map[[2]int]bool{}
	visitedLocations[*knots[len(knots)-1]] = true

	for _, motion := range motions {
		// do motion.Length times
		for range make([]struct{}, motion.Length) {
			for i, knot := range knots {
				if i == 0 {
					// Move Head
					x, y := motion.Direction.Vec()
					knot[0] += x
					knot[1] += y

					continue
				}

				// Following knots
				prevKnot := knots[i-1]

				// Move Tail
				htVec := [2]int{prevKnot[0] - knot[0], prevKnot[1] - knot[1]}
				switch htVec {
				// Straight Lines
				case [2]int{2, 0}: // 2 Right
					knot[0] += 1
				case [2]int{-2, 0}: // 2 Left
					knot[0] -= 1
				case [2]int{0, 2}: // 2 Up
					knot[1] += 1
				case [2]int{0, -2}: // 2 Down
					knot[1] -= 1

					// Diagonally
				case [2]int{1, 2}, [2]int{2, 1}, [2]int{2, 2}: // Up Right, Right Up
					knot[0] += 1
					knot[1] += 1
				case [2]int{1, -2}, [2]int{2, -1}, [2]int{2, -2}: // Right Down, Down Right
					knot[0] += 1
					knot[1] += -1
				case [2]int{-1, -2}, [2]int{-2, -1}, [2]int{-2, -2}: // Down Left, Left Down
					knot[0] += -1
					knot[1] += -1
				case [2]int{-1, 2}, [2]int{-2, 1}, [2]int{-2, 2}: // Left Up, Up Left
					knot[0] += -1
					knot[1] += 1
				}
			}

			// Keep track of unique locations
			visitedLocations[*knots[len(knots)-1]] = true
		}
	}

	return len(visitedLocations)
}

type Direction string

const (
	DirectionR = "R"
	DirectionL = "L"
	DirectionU = "U"
	DirectionD = "D"
)

func (d Direction) Vec() (int, int) {
	switch d {
	case DirectionR:
		return 1, 0
	case DirectionL:
		return -1, 0
	case DirectionU:
		return 0, 1
	case DirectionD:
		return 0, -1
	}

	return 0, 0
}

type Motion struct {
	Direction Direction
	Length    int
}

func ParseInput(input string) []Motion {
	motions := []Motion{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")

		direction := Direction(parts[0])

		length, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		motions = append(motions, Motion{direction, length})
	}

	return motions
}
