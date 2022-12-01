package main

import "sort"

func SolveCurrentDayWithTwist(input string) int {
	caloriesPerElf := getCaloriesPerElf(input)
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesPerElf)))

	return caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2]
}
