package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type elfInventory []int

func parseInput(input []string) ([]elfInventory, error) {
	inventory := make([]elfInventory, 0)

	currentElf := make(elfInventory, 0)
	for i, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			if len(currentElf) > 0 {
				inventory = append(inventory, currentElf)
				currentElf = make(elfInventory, 0)
			}
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid line (not an integer): %d: %s: %w", i+1, line, err)
		}
		currentElf = append(currentElf, calories)
	}

	if len(currentElf) > 0 {
		inventory = append(inventory, currentElf)
	}
	return inventory, nil
}

func getMostCalories(inventory []elfInventory) *int {
	if inventory == nil || len(inventory) == 0 {
		return nil
	}

	var mostCalories *int
	for _, elfInventory := range inventory {
		if elfInventory == nil {
			continue
		}

		elfCalories := 0
		for _, item := range elfInventory {
			elfCalories += item

		}
		if mostCalories == nil || elfCalories > *mostCalories {
			mostCalories = &elfCalories
		}
	}

	return mostCalories
}
