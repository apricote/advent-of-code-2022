package main

import (
	"strconv"

	"github.com/apricote/advent-of-code-2022/util"
)

func SolveCurrentDay(input string) int {
	caloriesPerElf := getCaloriesPerElf(input)

	maxCalories := 0
	for _, calories := range caloriesPerElf {
		if calories > maxCalories {
			maxCalories = calories
		}
	}

	return maxCalories
}

func getCaloriesPerElf(input string) []int {
	var caloriesPerElf []int

	var currentElf int

	for _, line := range util.SplitLines(input) {
		if line == "" {
			currentElf += 1
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if len(caloriesPerElf) < currentElf+1 {
			caloriesPerElf = append(caloriesPerElf, 0)

		}
		caloriesPerElf[currentElf] += calories
	}

	return caloriesPerElf
}
