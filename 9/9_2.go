package main

func SolveCurrentDayWithTwist(input string) int {
	motions := ParseInput(input)

	knots := Knots{
		&[2]int{0, 0},
		&[2]int{0, 0},
		&[2]int{0, 0},
		&[2]int{0, 0},
		&[2]int{0, 0},
		&[2]int{0, 0},
		&[2]int{0, 0},
		&[2]int{0, 0},
		&[2]int{0, 0},
		&[2]int{0, 0},
	}

	return SimulateDaRope(knots, motions)
}
