package main

import (
	"aoc22/day08"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	treeGrid, err := day08.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	visibleTrees := treeGrid.GetVisibleTrees()
	log.Printf("Number of visible trees: %d\n", len(visibleTrees))

	scenicScores := treeGrid.GetScenicScores()
	if scenicScores == nil {
		log.Fatal("No scenic scores found")
	}

	maxScore := 0
	for _, row := range scenicScores {
		for _, score := range row {
			if score > maxScore {
				maxScore = score
			}
		}
	}

	log.Printf("Max scenic score: %d\n", maxScore)

}
