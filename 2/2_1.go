package main

import "strings"

type Move string

const (
	MoveRock     Move = "ROCK"
	MovePaper    Move = "PAPER"
	MoveScissors Move = "SCISSORS"
)

type Game struct {
	EnemyMove Move
	OurMove   Move
}

func SolveCurrentDay(input string) int {
	lines := strings.Split(input, "\n")

	totalScore := 0

	for _, line := range lines {
		totalScore += ParseLine(line).Score()
	}

	return totalScore
}

func (g Game) Score() int {
	score := 0

	switch g.OurMove {
	case MoveRock:
		score += 1

		switch g.EnemyMove {
		case MoveRock:
			score += 3
		case MovePaper:
			score += 0
		case MoveScissors:
			score += 6
		}
	case MovePaper:
		score += 2

		switch g.EnemyMove {
		case MoveRock:
			score += 6
		case MovePaper:
			score += 3
		case MoveScissors:
			score += 0
		}
	case MoveScissors:
		score += 3

		switch g.EnemyMove {
		case MoveRock:
			score += 0
		case MovePaper:
			score += 6
		case MoveScissors:
			score += 3
		}
	}

	return score
}

func ParseLine(line string) Game {
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
		ourMove = MoveRock
	case "Y":
		ourMove = MovePaper
	case "Z":
		ourMove = MoveScissors
	}

	return Game{
		EnemyMove: enemyMove,
		OurMove:   ourMove,
	}
}
