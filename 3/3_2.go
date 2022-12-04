package main

import "strings"

func SolveCurrentDayWithTwist(input string) int {
	rucksacks := strings.Split(input, "\n")

	score := 0
	for i := 0; i < len(rucksacks); i += 3 {
		score += GetBadgePriority([]string{rucksacks[i+0], rucksacks[i+1], rucksacks[i+2]})
	}

	return score
}

func GetBadgePriority(rucksacks []string) int {
	var commonTypes uint64
	commonTypes = ^commonTypes

	for _, rucksack := range rucksacks {
		c1, c2 := GetRucksackCompartments(rucksack)
		totalTypes := c1 | c2

		commonTypes &= totalTypes
	}

	return GetPriority(commonTypes)
}
