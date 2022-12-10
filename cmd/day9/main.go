package main

import (
	"aoc22/day9"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	moves, err := day9.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	rope := day9.NewRope()
	for _, move := range moves {
		rope.Move(move)
	}

	result := rope.CountTailPositions()
	log.Printf("Number of tail positions: %d", result)

	lr, err := day9.NewLongRope(day9.Point{}, 8)
	if err != nil {
		log.Fatal(err)
	}

	for _, move := range moves {
		lr.Move(move)
	}

	result = lr.CountTailPositions()
	log.Printf("Part2: Number of tail positions: %d", result)

}
