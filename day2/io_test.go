package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name   string
		input  []string
		result Games
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
			result: []Game{{MyMove: PAPER, OpponentsMove: ROCK}},
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
					MyMove:        PAPER,
					OpponentsMove: ROCK,
				},
				{
					MyMove:        ROCK,
					OpponentsMove: PAPER,
				},
			},
			err: false,
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
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
	t.Parallel()
	for _, test := range []struct {
		name   string
		input  string
		result move
		err    bool
	}{
		{
			name:   "empty input",
			input:  "",
			result: INVALID_MOVE,
			err:    true,
		},
		{
			name:   "valid input ROCK",
			input:  "X",
			result: ROCK,
			err:    false,
		},
		{
			name:   "valid input PAPER",
			input:  "Y",
			result: PAPER,
			err:    false,
		},
		{
			name:   "valid input SCISSORS",
			input:  "Z",
			result: SCISSORS,
			err:    false,
		},
		{
			name:   "works with lower case input",
			input:  "z",
			result: SCISSORS,
			err:    false,
		},
		{
			name:   "works with whitespace",
			input:  "  z\t",
			result: SCISSORS,
			err:    false,
		},
		{
			name:   "fails when invalid string",
			input:  "  z dsfdf",
			result: INVALID_MOVE,
			err:    true,
		},
		{
			name:   "fails when invalid string",
			input:  "c",
			result: INVALID_MOVE,
			err:    true,
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
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
	t.Parallel()
	for _, test := range []struct {
		name   string
		input  string
		result move
		err    bool
	}{
		{
			name:   "empty input",
			input:  "",
			result: INVALID_MOVE,
			err:    true,
		},
		{
			name:   "valid input ROCK",
			input:  "A",
			result: ROCK,
			err:    false,
		},
		{
			name:   "valid input PAPER",
			input:  "B",
			result: PAPER,
			err:    false,
		},
		{
			name:   "valid input SCISSORS",
			input:  "C",
			result: SCISSORS,
			err:    false,
		},
		{
			name:   "works with lower case input",
			input:  "c",
			result: SCISSORS,
			err:    false,
		},
		{
			name:   "works with whitespace",
			input:  "  c\t",
			result: SCISSORS,
			err:    false,
		},
		{
			name:   "fails when invalid string",
			input:  "z",
			result: INVALID_MOVE,
			err:    true,
		},
		{
			name:   "fails when invalid string",
			input:  "  c dsfdf",
			result: INVALID_MOVE,
			err:    true,
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
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
