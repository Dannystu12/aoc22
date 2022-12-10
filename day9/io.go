package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(input []string) ([]move, error) {

	moves := make([]move, len(input))

	for i, line := range input {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid line %d: %s", i+1, line)
		}

		var dir direction
		switch strings.ToUpper(fields[0]) {
		case "U":
			dir = up
		case "D":
			dir = down
		case "L":
			dir = left
		case "R":
			dir = right
		default:
			return nil, fmt.Errorf("invalid line %d direction must be L|R|U|D : %s", i+1, line)
		}

		moveAmnt, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("invalid line %d move amount must be an integer: %s", i+1, line)
		}

		if moveAmnt < 0 {
			return nil, fmt.Errorf("invalid line %d move amount must be positive: %s", i+1, line)
		}

		moves[i] = move{dir, moveAmnt}
	}

	return moves, nil
}
