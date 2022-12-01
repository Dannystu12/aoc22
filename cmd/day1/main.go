package main

import (
	"aoc22/day1"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	inventories, err := day1.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	maxCalories := day1.GetMaxCalories(inventories)
	if maxCalories == nil {
		log.Println("No max calories found")
	} else {
		log.Println("Max Calories: ", *maxCalories)
	}

}
