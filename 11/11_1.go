package main

import (
	"sort"
	"strconv"
	"strings"
)

func SolveCurrentDay(input string) int {
	monkeys := ParseInput(input)

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for _, worryLevel := range monkey.Items {
				monkey.Productivity++

				worryLevel = monkey.Operation.Apply(worryLevel)
				worryLevel = worryLevel / 3

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
		productivies = append(productivies, monkey.Productivity)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(productivies)))

	monkeyBusiness := productivies[0] * productivies[1]

	return monkeyBusiness
}

type Monkey struct {
	Items           []int
	Operation       Operation
	TestDivisibleBy int
	TestTrue        int
	TestFalse       int
	Productivity    int
}

type Operator string

const (
	OperatorAdd      = "+"
	OperatorMultiply = "*"
)

type Operation struct {
	Operator   Operator
	OperandInt int
	OperandOld bool
}

func (o Operation) Apply(old int) int {
	operand := o.OperandInt
	if o.OperandOld {
		operand = old
	}

	switch o.Operator {
	case OperatorAdd:
		return old + operand
	case OperatorMultiply:
		return old * operand
	default:
		panic("unexpected operator")
	}
}

func ParseInput(input string) []*Monkey {
	monkeys := []*Monkey{}

	for _, description := range strings.Split(input, "\n\n") {
		lines := strings.Split(description, "\n")

		monkey := &Monkey{
			Items: []int{},
		}

		// 1 - starting items
		for _, item := range strings.Split(strings.TrimPrefix(lines[1], "  Starting items: "), ", ") {
			monkey.Items = append(monkey.Items, ParseInt(item))
		}

		// 2 - operation
		operationStrings := strings.Split(strings.TrimPrefix(lines[2], "  Operation: new = old "), " ")
		//   0 - operator
		monkey.Operation.Operator = Operator(operationStrings[0])
		//   1 - operand
		if operationStrings[1] == "old" {
			monkey.Operation.OperandOld = true
		} else {
			monkey.Operation.OperandInt = ParseInt(operationStrings[1])
		}

		// 3 - testDivisibleBy
		monkey.TestDivisibleBy = ParseInt(strings.TrimPrefix(lines[3], "  Test: divisible by "))

		// 4 - testtrue
		monkey.TestTrue = ParseInt(strings.TrimPrefix(lines[4], "    If true: throw to monkey "))

		// 5 - testfalse
		monkey.TestFalse = ParseInt(strings.TrimPrefix(lines[5], "    If false: throw to monkey "))

		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func ParseInt(s string) int {
	item, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return item
}
