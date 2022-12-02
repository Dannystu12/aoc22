package main

import (
	"aoc22/day2"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	lines := strings.Split(input, "\n")

	games, err := day2.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	score := games.GetTotalScore()
	fmt.Printf("The total score in part 1 is: %d\n", score)

	games, err = day2.ParseInput2(lines)
	if err != nil {
		log.Fatal(err)
	}

	score = games.GetTotalScore()
	fmt.Printf("The total score in part 2 is: %d\n", score)

}