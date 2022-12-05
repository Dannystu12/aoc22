package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//
//func ParseInput([]string) (cargo, []move, error) {
//	return nil, nil, nil
//}

func parseMoves(rows []string) ([]move, error) {
	moves := make([]move, 0, len(rows))
	for i, row := range rows {
		row = strings.ToLower(strings.TrimSpace(row))
		fields := strings.Fields(row)
		if len(fields) != 6 {
			return nil, fmt.Errorf("expected 6 fields, got %d on move row %d", len(fields), i+1)
		}

		if fields[0] != "move" || fields[2] != "from" || fields[4] != "to" {
			return nil, fmt.Errorf("expected format 'move X from X to X', got %v on move row %d", row, i+1)
		}

		rawNum, rawFrom, rawTo := fields[1], fields[3], fields[5]
		num, err := strconv.ParseUint(rawNum, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid move value %v on row %d, should be a uint", rawNum, i+1)
		}

		from, err := strconv.ParseUint(rawFrom, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid from value %v on row %d, should be a uint", rawFrom, i+1)
		}

		to, err := strconv.ParseUint(rawTo, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid to value %v on row %d, should be a uint", rawTo, i+1)
		}

		m := move{number: uint(num), from: uint(from), to: uint(to)}
		moves = append(moves, m)
	}
	return moves, nil
}

func parseCargo(rows []string) (cargo, error) {
	if rows == nil || len(rows) == 0 {
		return nil, fmt.Errorf("no rows provided to parse cargo")
	}

	baysRow := rows[len(rows)-1]

	bays, err := parseBay(baysRow)
	if err != nil {
		return nil, err
	}

	result := make(cargo)
	for bay, _ := range bays {
		result[bay] = make([]byte, 0)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no bays detected")
	}

	if len(rows) == 1 {
		return result, nil
	}

	for i := len(rows) - 2; i >= 0; i-- {
		cargoRow := rows[i]
		if len(cargoRow) != len(baysRow) {
			return nil, fmt.Errorf("cargo row %d is not the same length as bays row %d", len(cargoRow), len(baysRow))
		}

		spltRow := strings.Fields(cargoRow)
		for _, item := range spltRow {
			match, _ := regexp.MatchString("^\\[[A-Z]\\]$", item)
			if !match {
				return nil, fmt.Errorf("invalid crate value %s, must be in range A-Z and be enclosed in []", item)
			}
		}

		r, _ := regexp.Compile("[A-Z]]")
		matchIndexes := r.FindAllStringIndex(cargoRow, -1)

		for _, matchIndex := range matchIndexes {
			crate := cargoRow[matchIndex[0]]
			if crate < 'A' || crate > 'Z' {
				return nil, fmt.Errorf("invalid crate value %s, must be in range A-Z", string(crate))
			}

			bay := baysRow[matchIndex[0]]
			bayNum, err := strconv.ParseUint(string(bay), 10, 32)
			if err != nil {
				return nil, fmt.Errorf("invalid bay value %v for crate %s, should be a uint", bay, string(crate))
			}

			validCrateLength := len(rows) - 2 - i
			if len(result[uint(bayNum)]) != validCrateLength {
				return nil, fmt.Errorf("cannot add crate without crate underneath")
			}

			result[uint(bayNum)] = append(result[uint(bayNum)], crate)
		}

	}

	return result, nil
}

func parseBay(baysRow string) (map[uint]bool, error) {
	rawBays := strings.Fields(baysRow)

	result := make(map[uint]bool, 0)
	for _, rawBay := range rawBays {
		bay, err := strconv.ParseUint(rawBay, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid bay value %v, should be a uint", rawBay)
		}

		parsedBay := uint(bay)

		if _, ok := result[parsedBay]; ok {
			return nil, fmt.Errorf("duplicate bay detected %v", rawBay)
		}

		if parsedBay < 1 || parsedBay > 9 {
			return nil, fmt.Errorf("bay numbers must be between 1 and 9 (inclusive), give %d", parsedBay)
		}

		result[parsedBay] = true
	}

	return result, nil
}
