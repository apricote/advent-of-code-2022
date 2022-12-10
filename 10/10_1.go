package main

import (
	"strconv"
	"strings"
)

func SolveCurrentDay(input string) int {
	commands := ParseInput(input)

	// signal strengths
	signalStrengthsSum := 0

	SimulateCPU(commands, 220, func(cycle, X int) {
		// 20, 60, 100, 140, 180, 220
		if (cycle-20)%40 == 0 {
			// log.Printf("CPU: Cycle=%3d X=%3d", cycle, X)
			signalStrengthsSum += cycle * X
		}
	})

	return signalStrengthsSum
}

type Instruction string

const (
	InstructionNoop Instruction = "noop"
	InstructionAddx Instruction = "addx"
)

type Command struct {
	Instruction Instruction
	Value       int
}

func ParseInput(input string) []Command {
	lines := strings.Split(input, "\n")

	commands := make([]Command, 0, len(lines))

	for _, line := range lines {
		if line == "noop" {
			commands = append(commands, Command{
				InstructionNoop, 0,
			})
			continue
		}

		// addx
		V, err := strconv.Atoi(strings.TrimPrefix(line, "addx "))
		if err != nil {
			panic(err)
		}

		commands = append(commands, Command{
			InstructionAddx, V,
		})
	}

	return commands
}

func SimulateCPU(commands []Command, simulateCycles int, beforeCommand func(cycle, X int)) {
	// register
	X := 1

	// Command Buffer
	commandPointer := 0
	commandInProgress := false

	for cycle := 1; cycle <= simulateCycles; cycle++ {
		beforeCommand(cycle, X)

		command := commands[commandPointer]

		switch command.Instruction {
		case InstructionAddx:
			if commandInProgress {
				commandInProgress = false
				commandPointer++

				X += command.Value
			} else {
				commandInProgress = true
			}
		case InstructionNoop:
			commandPointer++
		}
	}
}
