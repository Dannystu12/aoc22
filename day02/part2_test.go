package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput2(t *testing.T) {
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
			result: []Game{{MyMove: ROCK, OpponentsMove: ROCK}},
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
			input: []string{"A Y", "B X", "C Z"},
			result: []Game{
				{
					MyMove:        ROCK,
					OpponentsMove: ROCK,
				},
				{
					MyMove:        ROCK,
					OpponentsMove: PAPER,
				},
				{
					MyMove:        ROCK,
					OpponentsMove: SCISSORS,
				},
			},
			err: false,
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := ParseInput2(test.input)
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

func TestStringToOutcome(t *testing.T) {
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

func TestGetPlayerMove(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name    string
		oppMove move
		outcome result
		result  move
	}{
		{
			name:    "rock win",
			oppMove: ROCK,
			outcome: WIN,
			result:  PAPER,
		},
		{
			name:    "rock draw",
			oppMove: ROCK,
			outcome: DRAW,
			result:  ROCK,
		},
		{
			name:    "rock loss",
			oppMove: ROCK,
			outcome: LOSS,
			result:  SCISSORS,
		},
		{
			name:    "paper win",
			oppMove: PAPER,
			outcome: WIN,
			result:  SCISSORS,
		},
		{
			name:    "paper draw",
			oppMove: PAPER,
			outcome: DRAW,
			result:  PAPER,
		},
		{
			name:    "paper loss",
			oppMove: PAPER,
			outcome: LOSS,
			result:  ROCK,
		},
		{
			name:    "scissors win",
			oppMove: SCISSORS,
			outcome: WIN,
			result:  ROCK,
		},
		{
			name:    "scissors draw",
			oppMove: SCISSORS,
			outcome: DRAW,
			result:  SCISSORS,
		},
		{
			name:    "scissors loss",
			oppMove: SCISSORS,
			outcome: LOSS,
			result:  PAPER,
		},
		{
			name:    "bad opp move",
			oppMove: "FOO",
			outcome: LOSS,
			result:  INVALID_MOVE,
		},
		{
			name:    "bad outcome",
			oppMove: SCISSORS,
			outcome: INVALID_RESULT,
			result:  INVALID_MOVE,
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := getPlayerMove(test.oppMove, test.outcome)
			assert.Equal(t, test.result, result)
		})
	}
}
