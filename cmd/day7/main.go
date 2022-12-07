package main

import (
	"aoc22/day7"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	fs, err := day7.ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	total := uint(0)
	fs.Traverse(func(key day7.FsEntryKey, entry day7.FsEntry) {
		if entry.IsDir() {
			sz := entry.GetSize()
			if sz <= 100000 {
				total += sz
			}
		}
	})

	log.Printf("Total size: %d", total)

}
