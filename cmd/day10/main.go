package main

import (
	"aoc22/day10"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	lines := strings.Split(input, "\n")

	commands, err := day10.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	cpu := day10.NewCPU()
	for _, cmd := range commands {
		cpu.ProcessCommand(cmd)
	}

	cycles := []int{20, 60, 100, 140, 180, 220}
	total := 0
	for _, c := range cycles {
		signalStrength, ok := cpu.GetSignalStrength(c)
		if !ok {
			log.Fatalf("could not get signal strength for cycle: %d\n", c)
		}
		total += signalStrength
	}

	log.Printf("Total signal strength: %d\n", total)

	crt := day10.NewCRT()

	result := crt.Draw(cpu)
	log.Println("Par2 result:")
	log.Print("\n" + strings.Replace(result, ".", " ", -1))

}
