package main

import (
	"aoc22/day6"
	_ "embed"
	"log"
)

//go:embed input.txt
var input string

func main() {

	marker, ok := day6.Signal(input).GetMarker()
	if !ok {
		log.Fatal("Failed to get marker")
	}

	log.Printf("Marker is %d", marker)

}