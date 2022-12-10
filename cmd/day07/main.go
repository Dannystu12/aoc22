package main

import (
	"aoc22/day07"
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	capacity := uint(70000000)

	fs, err := day07.ParseInput(lines, capacity)
	if err != nil {
		log.Fatal(err)
	}

	total := uint(0)
	fs.Traverse(func(key day07.FsEntryKey, entry day07.FsEntry) {
		if entry.IsDir() {
			sz := entry.GetSize()
			if sz <= 100000 {
				total += sz
			}
		}
	})

	log.Printf("Total size: %d", total)

	d, err := fs.RecommendDirectoryForDeletion(30000000)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Size of directory to delete: %d", d.GetSize())

}
