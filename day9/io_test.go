package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name  string
		input []string
		moves []move
		err   bool
	}{
		{
			name:  "empty input",
			input: []string{},
			moves: []move{},
			err:   false,
		},
		{
			name:  "nil input",
			input: nil,
			moves: []move{},
			err:   false,
		},
		{
			name: "basic example",
			input: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
				"L 5",
				"R 2",
			},
			moves: []move{
				{right, 4},
				{up, 4},
				{left, 3},
				{down, 1},
				{right, 4},
				{down, 1},
				{left, 5},
				{right, 2},
			},
			err: false,
		},
		{
			name: "empty line",
			input: []string{
				"R 4",
				"U 4",
				"",
				"D 1",
			},
			moves: nil,
			err:   true,
		},
		{
			name: "works with lower case",
			input: []string{
				"r 4",
			},
			moves: []move{{right, 4}},
			err:   false,
		},
		{
			name: "invalid direction",
			input: []string{
				"z 4",
			},
			moves: nil,
			err:   true,
		},
		{
			name: "amount must be positive",
			input: []string{
				"R -4",
			},
			moves: nil,
			err:   true,
		},
		{
			name: "not a number",
			input: []string{
				"R jkljkl",
			},
			moves: nil,
			err:   true,
		},
		{
			name: "line too long",
			input: []string{
				"R 4 R 9",
			},
			moves: nil,
			err:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			moves, err := ParseInput(test.input)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.moves, moves)
		})
	}
}
