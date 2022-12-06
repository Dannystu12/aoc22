package main

import (
	"aoc22/day6"
	_ "embed"
	"log"
)

//go:embed input.txt
var input string

func main() {

	marker, ok := day6.Signal(input).GetStartOfPacketMarker()
	if !ok {
		log.Fatal("Failed to get marker")
	}

	log.Printf("Start of packet marker is %d", marker)

	marker, ok = day6.Signal(input).GetStartOfMessageMarker()
	if !ok {
		log.Fatal("Failed to get marker")
	}

	log.Printf("Start of message marker is %d", marker)

}
