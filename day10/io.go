package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(lines []string) ([]command, error) {

	commands := make([]command, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)

		if len(fields) == 0 {
			return nil, fmt.Errorf("error on line %d no data to parse: %s", i+1, line)
		}

		switch fields[0] {
		case "noop":
			commands[i] = noopCommand{}
		case "addx":
			if len(fields) != 2 {
				return nil, fmt.Errorf("error on line %d addx requires an integer argument", i+1)
			}

			arg, err := strconv.Atoi(fields[1])
			if err != nil {
				return nil, fmt.Errorf("error on line %d addx requires a valid integer argument given %s: %w", i+1, fields[1], err)
			}

			commands[i] = addXCommand{
				value: arg,
			}
		default:
			return nil, fmt.Errorf("error on line %d unrecognised command: %s", i+1, fields[0])
		}

	}

	return commands, nil
}
