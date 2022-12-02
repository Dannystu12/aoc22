package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveToScore(t *testing.T) {
	for _, test := range []struct {
		name  string
		input move
		score int
	}{
		{
			name:  "ROCK",
			input: ROCK,
			score: 1,
		},
		{
			name:  "PAPER",
			input: PAPER,
			score: 2,
		},
		{
			name:  "SCISSORS",
			input: SCISSORS,
			score: 3,
		},
		{
			name:  "Invalid",
			input: INVALID_MOVE,
			score: 0,
		},
		{
			name:  "Junk",
			input: "Junk",
			score: 0,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result := test.input.toScore()
			assert.Equal(t, test.score, result)
		})
	}
}

func TestMoveValidate(t *testing.T) {
	for _, test := range []struct {
		name     string
		input    move
		expected move
		ok       bool
	}{
		{
			name:     "ROCK",
			input:    ROCK,
			expected: ROCK,
			ok:       true,
		},
		{
			name:     "PAPER",
			input:    PAPER,
			expected: PAPER,
			ok:       true,
		},
		{
			name:     "SCISSORS",
			input:    SCISSORS,
			expected: SCISSORS,
			ok:       true,
		},
		{
			name:     "JUNK",
			input:    "JUNK",
			expected: INVALID_MOVE,
			ok:       false,
		},
		{
			name:     "INVALID_MOVE",
			input:    INVALID_MOVE,
			expected: INVALID_MOVE,
			ok:       false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result, ok := test.input.Validate()
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.ok, ok)
		})
	}
}
