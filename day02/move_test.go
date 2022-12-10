package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveToScore(t *testing.T) {
	t.Parallel()
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
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := test.input.toScore()
			assert.Equal(t, test.score, result)
		})
	}
}

func TestMoveValidate(t *testing.T) {
	t.Parallel()
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
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, ok := test.input.Validate()
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.ok, ok)
		})
	}
}

func TestBeatenBy(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name     string
		input    move
		expected move
	}{
		{
			name:     "Paper beats rock",
			input:    ROCK,
			expected: PAPER,
		},
		{
			name:     "Rock beats scissors",
			input:    SCISSORS,
			expected: ROCK,
		},
		{
			name:     "Scissors beats paper",
			input:    PAPER,
			expected: SCISSORS,
		},
		{
			name:     "invalid",
			input:    INVALID_MOVE,
			expected: INVALID_MOVE,
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := test.input.BeatenBy()
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestBeats(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name     string
		input    move
		expected move
	}{
		{
			name:     "Paper beats rock",
			input:    PAPER,
			expected: ROCK,
		},
		{
			name:     "Rock beats scissors",
			input:    ROCK,
			expected: SCISSORS,
		},
		{
			name:     "Scissors beats paper",
			input:    SCISSORS,
			expected: PAPER,
		},
		{
			name:     "invalid",
			input:    INVALID_MOVE,
			expected: INVALID_MOVE,
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := test.input.Beats()
			assert.Equal(t, test.expected, result)
		})
	}
}
