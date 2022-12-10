package main

import (
	"aoc22/day01"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	inventories, err := day01.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	maxCalories := day01.GetMaxCalories(inventories)
	if maxCalories == nil {
		log.Println("No max calories found")
	} else {
		log.Println("Max Calories: ", *maxCalories)
	}

	top3MaxCalories := day01.GetMaxCaloriesTopN(inventories, 3)
	if top3MaxCalories == nil {
		log.Println("No top 3 max calories found")
	} else {
		log.Println("Top 3 Max Calories: ", *top3MaxCalories)
	}

}
