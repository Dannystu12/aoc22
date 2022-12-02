package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTotalScores(t *testing.T) {
	games := Games{
		{
			MyMove:        PAPER,
			OpponentsMove: ROCK,
		},
		{
			MyMove:        ROCK,
			OpponentsMove: PAPER,
		},
		{
			MyMove:        SCISSORS,
			OpponentsMove: SCISSORS,
		},
	}

	expectedScore := 15
	actualScore := games.GetTotalScore()
	assert.Equal(t, expectedScore, actualScore)
}

func TestGetResult(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  Game
		output result
	}{
		{
			name: "Rock win",
			input: Game{
				MyMove:        ROCK,
				OpponentsMove: SCISSORS,
			},
			output: WIN,
		},
		{
			name: "Paper win",
			input: Game{
				MyMove:        PAPER,
				OpponentsMove: ROCK,
			},
			output: WIN,
		},
		{
			name: "Scissors win",
			input: Game{
				MyMove:        SCISSORS,
				OpponentsMove: PAPER,
			},
			output: WIN,
		},
		{
			name: "Rock Draw",
			input: Game{
				MyMove:        ROCK,
				OpponentsMove: ROCK,
			},
			output: DRAW,
		},
		{
			name: "Paper Draw",
			input: Game{
				MyMove:        PAPER,
				OpponentsMove: PAPER,
			},
			output: DRAW,
		},
		{
			name: "Scissors Draw",
			input: Game{
				MyMove:        SCISSORS,
				OpponentsMove: SCISSORS,
			},
			output: DRAW,
		},
		{
			name: "Rock LOSS",
			input: Game{
				MyMove:        ROCK,
				OpponentsMove: PAPER,
			},
			output: LOSS,
		},
		{
			name: "Paper LOSS",
			input: Game{
				MyMove:        PAPER,
				OpponentsMove: SCISSORS,
			},
			output: LOSS,
		},
		{
			name: "Scissors LOSS",
			input: Game{
				MyMove:        SCISSORS,
				OpponentsMove: ROCK,
			},
			output: LOSS,
		},
		{
			name: "Invalid player move",
			input: Game{
				MyMove:        "FOO",
				OpponentsMove: ROCK,
			},
			output: INVALID_RESULT,
		},
		{
			name: "Invalid opponent move",
			input: Game{
				MyMove:        ROCK,
				OpponentsMove: "BAR",
			},
			output: INVALID_RESULT,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result := test.input.GetResult()
			assert.Equal(t, test.output, result)
		})
	}
}

func TestGetScore(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  Game
		output int
	}{
		{
			name: "Rock win",
			input: Game{
				MyMove:        ROCK,
				OpponentsMove: SCISSORS,
			},
			output: 7,
		},
		{
			name: "Paper win",
			input: Game{
				MyMove:        PAPER,
				OpponentsMove: ROCK,
			},
			output: 8,
		},
		{
			name: "Scissors win",
			input: Game{
				MyMove:        SCISSORS,
				OpponentsMove: PAPER,
			},
			output: 9,
		},
		{
			name: "Rock Draw",
			input: Game{
				MyMove:        ROCK,
				OpponentsMove: ROCK,
			},
			output: 4,
		},
		{
			name: "Paper Draw",
			input: Game{
				MyMove:        PAPER,
				OpponentsMove: PAPER,
			},
			output: 5,
		},
		{
			name: "Scissors Draw",
			input: Game{
				MyMove:        SCISSORS,
				OpponentsMove: SCISSORS,
			},
			output: 6,
		},
		{
			name: "Rock LOSS",
			input: Game{
				MyMove:        ROCK,
				OpponentsMove: PAPER,
			},
			output: 1,
		},
		{
			name: "Paper LOSS",
			input: Game{
				MyMove:        PAPER,
				OpponentsMove: SCISSORS,
			},
			output: 2,
		},
		{
			name: "Scissors LOSS",
			input: Game{
				MyMove:        SCISSORS,
				OpponentsMove: ROCK,
			},
			output: 3,
		},
		{
			name: "Invalid player move",
			input: Game{
				MyMove:        "FOO",
				OpponentsMove: ROCK,
			},
			output: 0,
		},
		{
			name: "Invalid opponent move",
			input: Game{
				MyMove:        ROCK,
				OpponentsMove: "BAR",
			},
			output: 0,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result := test.input.GetScore()
			assert.Equal(t, test.output, result)
		})
	}
}
