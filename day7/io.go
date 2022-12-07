package day7

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(lines []string, capacity uint) (*simpleFS, error) {
	simpleFS := newSimpleFS(capacity)

	isListingFiles := false

	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) <= 1 {
			return nil, fmt.Errorf("line %d is too short", i+1)
		}

		if fields[0] == "$" {
			isListingFiles = false
			cmd := fields[1]
			switch cmd {
			case "ls":
				if len(fields) != 2 {
					return nil, fmt.Errorf("error on line %d ls does not expect any arguments", i+1)
				}
				isListingFiles = true
				continue
			case "cd":
				if len(fields) != 3 {
					return nil, fmt.Errorf("error on line %d cd expects a single directory or '..'", i+1)
				}
				err := simpleFS.cd(fields[2])
				if err != nil {
					return nil, fmt.Errorf("error on line %d could not change directory: %w", i+1, err)
				}
				continue

			default:
				return nil, fmt.Errorf("unknown command %s", cmd)
			}
		}

		if !isListingFiles {
			return nil, fmt.Errorf("error on line %d invalid input (expecting ls output): %s", i+1, line)
		}

		if len(fields) != 2 {
			return nil, fmt.Errorf("error on line %d ls output must have 2 fields: %s", i+1, line)
		}

		if fields[0] == "dir" {
			dir, err := newDir(dirName(fields[1]))
			if err != nil {
				return nil, fmt.Errorf("error on line %d could not create directory: %w", i+1, err)
			}
			err = simpleFS.addEntry(dir)
			if err != nil {
				return nil, fmt.Errorf("error on line %d could not add directory: %w", i+1, err)
			}
			continue
		}

		sz, err := strconv.ParseUint(fields[0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("error on line %d size must be a valid uint: %w", i+1, err)
		}

		f, err := newFile(fileName(fields[1]), uint(sz))
		if err != nil {
			return nil, fmt.Errorf("error on line %d could not create file: %w", i+1, err)
		}

		err = simpleFS.addEntry(f)
		if err != nil {
			return nil, fmt.Errorf("error on line %d could not add file: %w", i+1, err)
		}

	}

	return simpleFS, nil
}
