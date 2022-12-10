package main

import (
	"aoc22/day05"
	_ "embed"
	"log"
)

//go:embed input.txt
var input string

func main() {

	cargo, moves, err := day05.ParseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range moves {
		err = cargo.PerformMove(m, false)
		if err != nil {
			log.Fatal(err)
		}
	}

	message := cargo.GetMessage()
	log.Printf("Part 1 Message: %s\n", message)

	cargo, moves, err = day05.ParseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range moves {
		err = cargo.PerformMove(m, true)
		if err != nil {
			log.Fatal(err)
		}
	}

	message = cargo.GetMessage()
	log.Printf("PART 2 Message: %s\n", message)

}
