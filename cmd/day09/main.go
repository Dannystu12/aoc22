package main

import (
	"aoc22/day09"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	moves, err := day09.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	rope := day09.NewRope()
	for _, move := range moves {
		rope.Move(move)
	}

	result := rope.CountTailPositions()
	log.Printf("Number of tail positions: %d", result)

	lr, err := day09.NewLongRope(day09.Point{}, 8)
	if err != nil {
		log.Fatal(err)
	}

	for _, move := range moves {
		lr.Move(move)
	}

	result = lr.CountTailPositions()
	log.Printf("Part2: Number of tail positions: %d", result)

}
