package main

import "strings"

func SolveCurrentDayWithTwist(input string) int {
	lines := strings.Split(input, "\n")

	totalScore := 0

	for _, line := range lines {
		totalScore += ParseLinePart2(line).Score()
	}

	return totalScore
}

func ParseLinePart2(line string) Game {
	moves := strings.Split(line, " ")

	var enemyMove Move
	var ourMove Move

	switch moves[0] {
	case "A":
		enemyMove = MoveRock
	case "B":
		enemyMove = MovePaper
	case "C":
		enemyMove = MoveScissors
	}

	switch moves[1] {
	case "X":
		switch enemyMove {
		case MoveRock:
			ourMove = MoveScissors
		case MovePaper:
			ourMove = MoveRock
		case MoveScissors:
			ourMove = MovePaper
		}
	case "Y":
		switch enemyMove {
		case MoveRock:
			ourMove = MoveRock
		case MovePaper:
			ourMove = MovePaper
		case MoveScissors:
			ourMove = MoveScissors
		}
	case "Z":
		switch enemyMove {
		case MoveRock:
			ourMove = MovePaper
		case MovePaper:
			ourMove = MoveScissors
		case MoveScissors:
			ourMove = MoveRock
		}
	}

	return Game{
		EnemyMove: enemyMove,
		OurMove:   ourMove,
	}
}
