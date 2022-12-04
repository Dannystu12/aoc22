package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(input []string) ([]sectionPair, error) {

	result := make([]sectionPair, len(input))
	for i, line := range input {
		line = strings.TrimSpace(line)
		splt := strings.Split(line, ",")
		if len(splt) != 2 {
			return nil, fmt.Errorf("line %d: invalid format line must contain 2 ranges", i+1)
		}

		rngA, rngB := splt[0], splt[1]

		srA, err := strToRange(rngA)
		if err != nil {
			return nil, fmt.Errorf("line %d: invalid format: %w", i+1, err)
		}

		srB, err := strToRange(rngB)
		if err != nil {
			return nil, fmt.Errorf("line %d: invalid format: %w", i+1, err)
		}

		result[i] = sectionPair{srA, srB}
	}

	return result, nil
}

func strToRange(s string) (sectionRange, error) {
	parts := strings.Split(strings.TrimSpace(s), "-")
	if len(parts) != 2 {
		return sectionRange{}, fmt.Errorf("invalid format line must contain 2 numbers seperated by '-'")
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return sectionRange{}, fmt.Errorf("%v is not a valid integer", parts[0])
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return sectionRange{}, fmt.Errorf("%v is not a valid integer", parts[1])
	}

	if start > end {
		return sectionRange{}, fmt.Errorf("%d is greater than %d", start, end)
	}

	return sectionRange{Min: start, Max: end}, nil

}
