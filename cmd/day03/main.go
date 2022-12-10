package main

import (
	"aoc22/day03"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	lines := strings.Split(input, "\n")

	rucksacks, err := day03.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	totalScore := 0
	for _, r := range rucksacks {
		duplicate, err := r.GetDuplicate()
		if err != nil {
			log.Fatal(err)
		}

		score, err := day03.GetPriorityScore(*duplicate)
		if err != nil {
			log.Fatal(err)
		}

		totalScore += score
	}

	log.Printf("Total Priority Score: %d", totalScore)

	rucksackGroups, err := day03.ParseInput2(lines)
	totalScore2 := 0
	for _, rg := range rucksackGroups {
		duplicate, err := rg.GetDuplicate()
		if err != nil {
			log.Fatal(err)
		}
		score, err := day03.GetPriorityScore(*duplicate)
		if err != nil {
			log.Fatal(err)
		}
		totalScore2 += score
	}

	log.Printf("Part 2 Total Priority Score: %d", totalScore2)

}
