package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	inventories, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	maxCalories := getMostCalories(inventories)
	log.Println("Max Calories: ", maxCalories)

}
