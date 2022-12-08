package main

import (
	"strconv"
	"strings"
)

func SolveCurrentDay(input string) int {
	trees := ParseInput(input)

	return CountVisibleTrees(trees)
}

func ParseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	trees := make([][]int, len(lines))

	for y, line := range lines {
		for _, heightStr := range line {
			height, err := strconv.Atoi(string(heightStr))
			if err != nil {
				panic(err)
			}

			trees[y] = append(trees[y], height)
		}
	}

	return trees
}

func CountVisibleTrees(trees [][]int) int {
	visibleTrees := map[[2]int]bool{}

	// Iterate all visibility lines
	// Top-Bottom
	// Bottom-Top
	// Left-Right
	// Right-Left
	CountVisibleTreesInDirection(trees, visibleTrees, DirectionLR)
	CountVisibleTreesInDirection(trees, visibleTrees, DirectionTB)
	CountVisibleTreesInDirection(trees, visibleTrees, DirectionBT)
	CountVisibleTreesInDirection(trees, visibleTrees, DirectionRL)

	return len(visibleTrees)
}

type Direction int

const (
	DirectionTB Direction = iota
	DirectionBT
	DirectionLR
	DirectionRL
)

func CountVisibleTreesInDirection(trees [][]int, visibleTrees map[[2]int]bool, direction Direction) {
	OUTER_LENGTH := 0
	INNER_LENGTH := 0

	switch direction {
	case DirectionTB, DirectionBT:
		OUTER_LENGTH = len(trees[0])
		INNER_LENGTH = len(trees)
	case DirectionLR, DirectionRL:
		OUTER_LENGTH = len(trees)
		INNER_LENGTH = len(trees[0])
	}

	LOS_START := 0
	LOS_MODIFIER := 0

	switch direction {
	case DirectionTB, DirectionLR:
		LOS_START = 0
		LOS_MODIFIER = 1
	case DirectionBT, DirectionRL:
		LOS_START = INNER_LENGTH - 1
		LOS_MODIFIER = -1
	}

	switch direction {
	case DirectionTB, DirectionBT:
		for x := 0; x < OUTER_LENGTH; x++ {
			for y := 0; y < INNER_LENGTH; y++ {
				largestInLOS := -1
				for yOther := LOS_START; (LOS_MODIFIER == 1 && yOther < y) || (LOS_MODIFIER == -1 && yOther > y); yOther += LOS_MODIFIER {
					if trees[yOther][x] > largestInLOS {
						largestInLOS = trees[yOther][x]
					}
				}

				if trees[y][x] > largestInLOS {
					visibleTrees[[2]int{y, x}] = true
				}
			}
		}
	case DirectionLR, DirectionRL:
		for y := 0; y < OUTER_LENGTH; y++ {
			for x := 0; x < INNER_LENGTH; x++ {
				largestInLOS := -1
				for xOther := LOS_START; (LOS_MODIFIER == 1 && xOther < x) || (LOS_MODIFIER == -1 && xOther > x); xOther += LOS_MODIFIER {
					if trees[y][xOther] > largestInLOS {
						largestInLOS = trees[y][xOther]
					}
				}

				if trees[y][x] > largestInLOS {
					visibleTrees[[2]int{y, x}] = true
				}
			}
		}
	}
}
