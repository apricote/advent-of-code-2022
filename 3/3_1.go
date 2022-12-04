package main

import (
	"math/bits"
	"strings"
)

func SolveCurrentDay(input string) int {
	rucksacks := strings.Split(input, "\n")

	score := 0
	for _, rucksack := range rucksacks {
		score += GetRucksackScore(rucksack)
	}

	return score
}

func GetRucksackCompartments(rucksack string) (uint64, uint64) {
	compartmentLenght := len(rucksack) / 2

	var typesInCompartmentOne uint64
	var typesInCompartmentTwo uint64

	for i := 0; i < len(rucksack); i++ {
		content := rucksack[i]
		var typeBit byte

		if content >= 97 {
			// a in ASCII is 97 but in our priority list its 1
			typeBit = content - 97 + 1
		} else {
			// A in ASCII is 65 but in our priority list its 27
			typeBit = content - 65 + 27
		}

		if i < compartmentLenght {
			typesInCompartmentOne |= 1 << typeBit
		} else {
			typesInCompartmentTwo |= 1 << typeBit
		}
	}

	return typesInCompartmentOne, typesInCompartmentTwo
}

func GetRucksackScore(rucksack string) int {
	typesInCompartmentOne, typesInCompartmentTwo := GetRucksackCompartments(rucksack)

	commonTypes := typesInCompartmentOne & typesInCompartmentTwo

	return GetPriority(commonTypes)
}

func GetPriority(typeContents uint64) int {
	return bits.TrailingZeros64(typeContents)
}
