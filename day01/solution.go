package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type elfInventory []int

func ParseInput(input []string) ([]elfInventory, error) {
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

func GetMaxCalories(inventory []elfInventory) *int {
	_, max := getMaxCalories(inventory)
	return max
}

func GetMaxCaloriesTopN(inventory []elfInventory, topN uint) *int {

	if inventory == nil || len(inventory) == 0 || topN == 0 {
		return nil
	}

	invCpy := make([]elfInventory, len(inventory))
	copy(invCpy, inventory)

	var total *int
	var i uint
	for i = 0; i < topN; i++ {
		idx, max := getMaxCalories(invCpy)
		if max == nil {
			break
		}

		if total == nil {
			total = max
		} else {
			*total += *max
		}

		invCpy = append(invCpy[:*idx], invCpy[*idx+1:]...)
	}

	return total
}

func getMaxCalories(inventory []elfInventory) (*int, *int) {
	if inventory == nil || len(inventory) == 0 {
		return nil, nil
	}

	var mostCalories *int
	var idx *int
	for i, elfInventory := range inventory {
		if elfInventory == nil {
			continue
		}

		elfCalories := 0
		elfIdx := i
		for _, item := range elfInventory {
			elfCalories += item
		}
		if mostCalories == nil || elfCalories > *mostCalories {
			mostCalories = &elfCalories
			idx = &elfIdx
		}
	}

	return idx, mostCalories
}
