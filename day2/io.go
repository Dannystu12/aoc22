package day2

import (
	"fmt"
	"strings"
)

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
		return INVALID_MOVE, fmt.Errorf("player move must be one of [X,Y,Z], given %s", s)
	}

	switch s_ {
	case "x":
		return ROCK, nil
	case "y":
		return PAPER, nil
	case "z":
		return SCISSORS, nil
	default:
		return INVALID_MOVE, fmt.Errorf("opponent move must be one of [X,Y,Z], given %s", s)
	}
}

func stringToOpponentMove(s string) (move, error) {
	s_ := strings.TrimSpace(strings.ToLower(s))
	if len(s_) != 1 {
		return INVALID_MOVE, fmt.Errorf("opponent move must be one of [A,B,C], given %s", s)
	}

	switch s_ {
	case "a":
		return ROCK, nil
	case "b":
		return PAPER, nil
	case "c":
		return SCISSORS, nil
	default:
		return INVALID_MOVE, fmt.Errorf("opponent move must be one of [A,B,C], given %s", s)
	}
}
