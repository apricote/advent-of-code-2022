package main

func SolveCurrentDayWithTwist(input string) string {
	commands := ParseInput(input)

	display := ""

	SimulateCPU(commands, 240, func(cycle, X int) {
		horizontalPosition := (cycle - 1) % 40
		if horizontalPosition == 0 {
			display += "\n"
		}

		spriteVisible := horizontalPosition == X || horizontalPosition == X-1 || horizontalPosition == X+1
		if spriteVisible {
			display += "#"
		} else {
			display += "."
		}

		// log.Printf("Cycle: %3d X: %3d horizontalPosition: %2d Visible: %t", cycle, X, horizontalPosition, spriteVisible)
	})

	return display
}
