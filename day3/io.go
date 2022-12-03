package day3

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func ParseInput(input []string) ([]rucksack, error) {
	if input == nil || len(input) == 0 {
		return []rucksack{}, nil
	}

	result := make([]rucksack, 0, len(input))
	for i, line := range input {
		line = strings.TrimSpace(line)

		lineLength := len(line)

		if lineLength == 0 {
			continue
		}

		if utf8.RuneCountInString(line) != lineLength {
			return nil, fmt.Errorf("error on line %d: contains utf8 characters: %s", i+1, line)
		}

		if lineLength%2 != 0 {
			return nil, fmt.Errorf("error on line %d: must contain an even number of characters: %s", i+1, line)
		}

		firstHalf := line[:lineLength/2]
		secondHalf := line[lineLength/2:]

		compartmentA := make(compartment, len(firstHalf))
		compartmentB := make(compartment, len(secondHalf))

		for _, c := range firstHalf {
			compartmentA[byte(c)] = true
		}

		numDuplicates := 0
		for _, c := range secondHalf {
			if compartmentA[byte(c)] && !compartmentB[byte(c)] {
				numDuplicates++
			}
			compartmentB[byte(c)] = true
			if numDuplicates > 1 {
				break
			}
		}

		if numDuplicates != 1 {
			return nil, fmt.Errorf("error on line %d: must contain exactly one duplicate between compartments: %s", i+1, line)
		}

		//if len(compartmentA)+len(compartmentB) != lineLength {
		//	return nil, fmt.Errorf("error on line %d: compartment size mismatch: A:%d B: %d line: %s", i+1, len(compartmentA), len(compartmentB), line)
		//}

		result = append(result, rucksack{compartmentA, compartmentB})
	}

	return result, nil

}
