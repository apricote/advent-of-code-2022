package main

import "strings"

func SolveCurrentDayWithTwist(input string) string {
	inputParts := strings.Split(input, "\n\n")
	startingStacksInput := inputParts[0]
	rearrangementProcedureInput := strings.Split(inputParts[1], "\n")

	stacks := ParseStacksInput(startingStacksInput)
	rearrangementProcedure := ParseRearrangementProcedure(rearrangementProcedureInput)

	for _, instruction := range rearrangementProcedure {
		//fmt.Printf("Instruction: %+v\n", instruction)
		//fmt.Printf("Stacks: %v\n", stacks)
		instruction.Apply9001(stacks)
	}

	//fmt.Printf("Final Stacks: %v\n", stacks)

	solution := ""
	for _, stack := range stacks {
		solution += stack.Top()
	}

	return solution
}

func (r RearrangementInstruction) Apply9001(stacks []CrateStack) {
	tmpStack := CrateStack{}
	for i := 0; i < r.Amount; i++ {
		crate := stacks[r.From].Pop()
		tmpStack.Push(crate)
	}

	for i := 0; i < r.Amount; i++ {
		crate := tmpStack.Pop()
		stacks[r.To].Push(crate)
	}
}
