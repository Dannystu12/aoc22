package day2

import (
	"fmt"
	"strings"
)

type move string

const (
	Rock     move = "Rock"
	Scissors move = "Scissors"
	Paper    move = "Paper"
	INVALID  move = "Invalid"
)

type Game struct {
	MyMove        move
	OpponentsMove move
}

func ParseInput(input []string) ([]Game, error) {

	games := make([]Game, 0)

	if input == nil || len(input) == 0 {
		return games, nil
	}

	for i, line := range input {
		splt := strings.Fields(line)
		if len(splt) != 2 {
			return nil, fmt.Errorf("line %d is invalid, it doesn't have 2 entries: %s", i+1, line)
		}

		opponentMove, err := stringToOpponentMove(splt[0])
		if err != nil {
			return nil, fmt.Errorf("line %d is invalid: %w", i+1, err)
		}

		playerMove, err := stringToPlayerMove(splt[1])
		if err != nil {
			return nil, fmt.Errorf("line %d is invalid: %w", i+1, err)
		}

		game := Game{
			MyMove:        playerMove,
			OpponentsMove: opponentMove,
		}

		games = append(games, game)
	}

	return games, nil
}

func stringToPlayerMove(s string) (move, error) {
	s_ := strings.TrimSpace(strings.ToLower(s))
	if len(s_) != 1 {
		return INVALID, fmt.Errorf("player move must be one of [X,Y,Z], given %s", s)
	}

	switch s_ {
	case "x":
		return Rock, nil
	case "y":
		return Paper, nil
	case "z":
		return Scissors, nil
	default:
		return INVALID, fmt.Errorf("opponent move must be one of [X,Y,Z], given %s", s)
	}
}

func stringToOpponentMove(s string) (move, error) {
	s_ := strings.TrimSpace(strings.ToLower(s))
	if len(s_) != 1 {
		return INVALID, fmt.Errorf("opponent move must be one of [A,B,C], given %s", s)
	}

	switch s_ {
	case "a":
		return Rock, nil
	case "b":
		return Paper, nil
	case "c":
		return Scissors, nil
	default:
		return INVALID, fmt.Errorf("opponent move must be one of [A,B,C], given %s", s)
	}
}
