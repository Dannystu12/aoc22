package day08

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(lines []string) (treeGrid, error) {

	result := make(treeGrid, 0)
	for i, line := range lines {
		line = strings.TrimSpace(line)
		row := make([]treeHeight, 0)
		for j, char := range line {
			parsedChar, err := strconv.ParseInt(string(char), 10, 32)
			if err != nil {
				return nil, fmt.Errorf("error on line %d col %d, invalid int: %w", i, j, err)
			}
			parsedHeight := int(parsedChar)
			th := treeHeight(parsedHeight)
			if err := th.isValid(); err != nil {
				return nil, fmt.Errorf("error on line %d col %d, invalid tree height: %w", i, j, err)
			}
			row = append(row, th)
		}
		result = append(result, row)
	}

	if err := result.isValid(); err != nil {
		return nil, fmt.Errorf("invalid tree grid: %w", err)
	}

	return result, nil
}
