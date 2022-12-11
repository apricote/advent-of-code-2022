package main

import (
	"sort"
)

func SolveCurrentDayWithTwist(input string) int {
	monkeys := ParseInput(input)

	godModeModulos := 1
	for _, monkey := range monkeys {
		godModeModulos *= monkey.TestDivisibleBy
	}

	for i := 0; i < 10_000; i++ {
		for _, monkey := range monkeys {
			for _, worryLevel := range monkey.Items {
				monkey.Productivity++

				worryLevel = monkey.Operation.Apply(worryLevel)
				worryLevel = worryLevel % godModeModulos

				destinationMonkey := -1

				if worryLevel%monkey.TestDivisibleBy == 0 {
					destinationMonkey = monkey.TestTrue
				} else {
					destinationMonkey = monkey.TestFalse
				}

				monkeys[destinationMonkey].Items = append(monkeys[destinationMonkey].Items, worryLevel)
			}

			monkey.Items = []int{}
		}
	}

	productivies := []int{}
	for _, monkey := range monkeys {
		// log.Printf("Monkey %d inspected items %d times.", i, monkey.Productivity)
		productivies = append(productivies, monkey.Productivity)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(productivies)))

	monkeyBusiness := productivies[0] * productivies[1]

	return monkeyBusiness
}
