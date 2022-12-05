package main

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	parseInstructionRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
)

func SolveCurrentDay(input string) string {
	inputParts := strings.Split(input, "\n\n")
	startingStacksInput := inputParts[0]
	rearrangementProcedureInput := strings.Split(inputParts[1], "\n")

	stacks := ParseStacksInput(startingStacksInput)
	rearrangementProcedure := ParseRearrangementProcedure(rearrangementProcedureInput)

	for _, instruction := range rearrangementProcedure {
		//fmt.Printf("Instruction: %+v\n", instruction)
		//fmt.Printf("Stacks: %v\n", stacks)
		instruction.Apply(stacks)
	}

	//fmt.Printf("Final Stacks: %v\n", stacks)

	solution := ""
	for _, stack := range stacks {
		solution += stack.Top()
	}

	return solution
}

type CrateStack struct {
	stack []string
}

func (c *CrateStack) Push(crate string) {
	c.stack = append(c.stack, crate)
}

func (c *CrateStack) Pop() string {
	top := c.Top()
	c.stack = c.stack[:len(c.stack)-1]

	return top
}

func (c CrateStack) Top() string {
	if len(c.stack) == 0 {
		return ""
	}
	return c.stack[len(c.stack)-1]
}

func ParseStacksInput(startingStacksInput string) []CrateStack {
	lines := strings.Split(startingStacksInput, "\n")

	stackCount := 0
	var stacks []CrateStack

	last := len(lines) - 1
	for i := range lines {
		// we iterate from bottom to top
		line := lines[last-i]

		if i == 0 {
			// prepare stacks
			// example line:
			//   " 1   2   3   4 "
			// 3x + x - 1 = length of line
			// x = (length of line + 1)/4
			stackCount = (len(line) + 1) / 4

			stacks = make([]CrateStack, stackCount)

			continue
		}

		// parse regular stack info here
		for i := 0; i < stackCount; i++ {
			// "[Z] [M] [P]"
			//   1   5   9
			cratePosition := i*4 + 1

			crate := string(line[cratePosition])
			if crate != " " {
				stacks[i].Push(crate)
			}
		}
	}

	return stacks
}

type RearrangementInstruction struct {
	Amount int
	From   int
	To     int
}

func (r RearrangementInstruction) Apply(stacks []CrateStack) {
	for i := 0; i < r.Amount; i++ {
		crate := stacks[r.From].Pop()
		stacks[r.To].Push(crate)
	}
}

func ParseRearrangementProcedure(rearrangementProcedureInput []string) []RearrangementInstruction {
	rearrangementProcedure := make([]RearrangementInstruction, 0, len(rearrangementProcedureInput))

	for _, input := range rearrangementProcedureInput {
		rearrangementProcedure = append(rearrangementProcedure, ParseRearrangementInstruction(input))
	}

	return rearrangementProcedure
}

func ParseRearrangementInstruction(rearrangementInstructionInput string) RearrangementInstruction {
	capturedInts := parseInstructionRegex.FindStringSubmatch(rearrangementInstructionInput)

	amount, err := strconv.Atoi(capturedInts[1])
	if err != nil {
		panic(err)
	}

	from, err := strconv.Atoi(capturedInts[2])
	if err != nil {
		panic(err)
	}

	to, err := strconv.Atoi(capturedInts[3])
	if err != nil {
		panic(err)
	}

	return RearrangementInstruction{
		amount,
		from - 1,
		to - 1,
	}
}
