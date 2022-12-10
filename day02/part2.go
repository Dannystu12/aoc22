package day02

import (
	"fmt"
	"strings"
)

func ParseInput2(input []string) (Games, error) {
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

		outcome, err := stringToResult(splt[1])
		if err != nil {
			return nil, fmt.Errorf("line %d is invalid: %w", i+1, err)
		}

		playerMove := getPlayerMove(opponentMove, outcome)

		game := Game{
			MyMove:        playerMove,
			OpponentsMove: opponentMove,
		}

		games = append(games, game)
	}

	return games, nil
}

func stringToResult(s string) (result, error) {
	s_ := strings.TrimSpace(strings.ToLower(s))
	if len(s_) != 1 {
		return INVALID_RESULT, fmt.Errorf("outcome must be one of [X,Y,Z], given %s", s)
	}

	switch s_ {
	case "x":
		return LOSS, nil
	case "y":
		return DRAW, nil
	case "z":
		return WIN, nil
	default:
		return INVALID_RESULT, fmt.Errorf("outcome must be one of [X,Y,Z], given %s", s)
	}
}

func getPlayerMove(opponentMove move, outcome result) move {
	outcome, ok := outcome.Validate()
	if !ok {
		return INVALID_MOVE
	}

	opponentMove, ok = opponentMove.Validate()
	if !ok {
		return INVALID_MOVE
	}

	switch outcome {
	case WIN:
		return opponentMove.BeatenBy()
	case DRAW:
		return opponentMove
	case LOSS:
		return opponentMove.Beats()
	}
	return INVALID_MOVE
}
