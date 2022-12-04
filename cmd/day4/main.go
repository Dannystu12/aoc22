package main

import (
	"aoc22/day4"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	lines := strings.Split(input, "\n")

	sectionPairs, err := day4.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	fullyContainsCount := 0
	anyOverlapCount := 0
	for _, sp := range sectionPairs {
		if sp.FullyContains() {
			fullyContainsCount++
		}
		if sp.AnyOverlap() {
			anyOverlapCount++
		}
	}

	log.Printf("Ranges that fully contain each other: %d\n", fullyContainsCount)
	log.Printf("Ranges that overlap at all: %d\n", anyOverlapCount)

}
