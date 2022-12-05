package main

import (
	"aoc22/day5"
	_ "embed"
	"log"
)

//go:embed input.txt
var input string

func main() {

	cargo, moves, err := day5.ParseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range moves {
		err = cargo.PerformMove(m)
		if err != nil {
			log.Fatal(err)
		}
	}

	message := cargo.GetMessage()
	log.Printf("Message: %s\n", message)

}
