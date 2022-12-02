package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  []string
		result []Game
		err    bool
	}{
		{
			name:   "empty input",
			input:  []string{},
			result: []Game{},
			err:    false,
		},
		{
			name:   "nil input",
			input:  nil,
			result: []Game{},
			err:    false,
		},
		{
			name:   "input without space",
			input:  []string{"A Y", "BX"},
			result: nil,
			err:    true,
		},
		{
			name:   "input with double space",
			input:  []string{"A  Y"},
			result: []Game{{MyMove: Paper, OpponentsMove: Rock}},
			err:    false,
		},
		{
			name:   "input without valid player move",
			input:  []string{"A Y", "B Q"},
			result: nil,
			err:    true,
		},
		{
			name:   "input without valid opponent",
			input:  []string{"A Y", "Z Z"},
			result: nil,
			err:    true,
		},
		{
			name:   "input with junk data",
			input:  []string{"A Y", "B X I"},
			result: nil,
			err:    true,
		},
		{
			name:  "valid input basic",
			input: []string{"A Y", "B X"},
			result: []Game{
				{
					MyMove:        Paper,
					OpponentsMove: Rock,
				},
				{
					MyMove:        Rock,
					OpponentsMove: Paper,
				},
			},
			err: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result, err := ParseInput(test.input)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				if test.result == nil {
					assert.Nil(t, result)
				} else {
					assert.Equal(t, test.result, result)
				}

			}
		})
	}
}

func TestStringToPlayerMove(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		result move
		err    bool
	}{
		{
			name:   "empty input",
			input:  "",
			result: INVALID,
			err:    true,
		},
		{
			name:   "valid input Rock",
			input:  "X",
			result: Rock,
			err:    false,
		},
		{
			name:   "valid input Paper",
			input:  "Y",
			result: Paper,
			err:    false,
		},
		{
			name:   "valid input Scissors",
			input:  "Z",
			result: Scissors,
			err:    false,
		},
		{
			name:   "works with lower case input",
			input:  "z",
			result: Scissors,
			err:    false,
		},
		{
			name:   "works with whitespace",
			input:  "  z\t",
			result: Scissors,
			err:    false,
		},
		{
			name:   "fails when invalid string",
			input:  "  z dsfdf",
			result: INVALID,
			err:    true,
		},
		{
			name:   "fails when invalid string",
			input:  "c",
			result: INVALID,
			err:    true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result, err := stringToPlayerMove(test.input)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.result, result)
		})
	}
}

func TestStringToOpponentMove(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		result move
		err    bool
	}{
		{
			name:   "empty input",
			input:  "",
			result: INVALID,
			err:    true,
		},
		{
			name:   "valid input Rock",
			input:  "A",
			result: Rock,
			err:    false,
		},
		{
			name:   "valid input Paper",
			input:  "B",
			result: Paper,
			err:    false,
		},
		{
			name:   "valid input Scissors",
			input:  "C",
			result: Scissors,
			err:    false,
		},
		{
			name:   "works with lower case input",
			input:  "c",
			result: Scissors,
			err:    false,
		},
		{
			name:   "works with whitespace",
			input:  "  c\t",
			result: Scissors,
			err:    false,
		},
		{
			name:   "fails when invalid string",
			input:  "z",
			result: INVALID,
			err:    true,
		},
		{
			name:   "fails when invalid string",
			input:  "  c dsfdf",
			result: INVALID,
			err:    true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result, err := stringToOpponentMove(test.input)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.result, result)
		})
	}
}
