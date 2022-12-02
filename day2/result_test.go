package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResultToScore(t *testing.T) {
	for _, test := range []struct {
		name  string
		input result
		score int
	}{
		{
			name:  "WIN",
			input: WIN,
			score: 6,
		},
		{
			name:  "LOSS",
			input: LOSS,
			score: 0,
		},
		{
			name:  "DRAW",
			input: DRAW,
			score: 3,
		},
		{
			name:  "JUNK",
			input: "JUNK",
			score: 0,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			result := test.input.toScore()
			assert.Equal(t, test.score, result)
		})
	}
}

func TestResultValidate(t *testing.T) {
	for _, test := range []struct {
		name     string
		input    result
		expected result
		ok       bool
	}{
		{
			name:     "WIN",
			input:    WIN,
			expected: WIN,
			ok:       true,
		},
		{
			name:     "LOSS",
			input:    LOSS,
			expected: LOSS,
			ok:       true,
		},
		{
			name:     "DRAW",
			input:    DRAW,
			expected: DRAW,
			ok:       true,
		},
		{
			name:     "JUNK",
			input:    "JUNK",
			expected: INVALID_RESULT,
			ok:       false,
		},
		{
			name:     "INVALID_RESULT",
			input:    INVALID_RESULT,
			expected: INVALID_RESULT,
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
