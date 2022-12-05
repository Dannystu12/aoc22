package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMoves(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result []move
		err    bool
	}{
		{
			name: "valid input",
			input: []string{
				"move 1 from 2 to 1",
				"move 3 from 1 to 3",
				"move 2 from 2 to 1",
				"move 1 from 1 to 2",
			},
			result: []move{
				{number: 1, from: 2, to: 1},
				{number: 3, from: 1, to: 3},
				{number: 2, from: 2, to: 1},
				{number: 1, from: 1, to: 2},
			},
			err: false,
		},
		{
			name:   "nil input",
			input:  nil,
			result: []move{},
			err:    false,
		},
		{
			name:   "empty input",
			input:  []string{},
			result: []move{},
			err:    false,
		},
		{
			name: "junk input",
			input: []string{
				"tjkljkljlkj",
			},
			result: nil,
			err:    true,
		},
		{
			name: "non uint val num",
			input: []string{
				"move -1 from 2 to 1",
			},
			result: nil,
			err:    true,
		},
		{
			name: "non uint val from",
			input: []string{
				"move 1 from -2 to 1",
			},
			result: nil,
			err:    true,
		},
		{
			name: "non uint val to",
			input: []string{
				"move 1 from 2 to -1",
			},
			result: nil,
			err:    true,
		},
		{
			name: "case insensitive",
			input: []string{
				"MOVE 1 frOm 2 to 1",
			},
			result: []move{
				{number: 1, from: 2, to: 1},
			},
			err: false,
		},
		{
			name: "works with whitespace",
			input: []string{
				" move 1 from 2   to 1\t",
			},
			result: []move{
				{number: 1, from: 2, to: 1},
			},
			err: false,
		},
		{
			name: "expect format move",
			input: []string{
				"foo 1 from 2 to 1",
			},
			result: nil,
			err:    true,
		},
		{
			name: "expect format from",
			input: []string{
				"move 1 foo 2 to 1",
			},
			result: nil,
			err:    true,
		},
		{
			name: "expect format to",
			input: []string{
				"move 1 from 2 foo 1",
			},
			result: nil,
			err:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			moves, err := parseMoves(test.input)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, moves)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.result, moves)
			}
		})
	}
}

func TestParseCargo(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result cargo
		err    bool
	}{
		{
			name: "no crates",
			input: []string{
				" 1   2   3 ",
			},
			result: cargo{
				1: []byte{},
				2: []byte{},
				3: []byte{},
			},
			err: false,
		},
		{
			name: "bay rows must be ordered",
			input: []string{
				" 1   2   1 ",
			},
			result: nil,
			err:    true,
		},
		{
			name: "bays must be 1-9",
			input: []string{
				" 1   2   10 ",
			},
			result: nil,
			err:    true,
		},
		{
			name:   "empty input",
			input:  []string{},
			result: nil,
			err:    true,
		},
		{
			name:   "nil input",
			input:  nil,
			result: nil,
			err:    true,
		},
		{
			name: "junk bays",
			input: []string{
				" hkjlj ",
			},
			result: nil,
			err:    true,
		},
		{
			name: "non uint bays",
			input: []string{
				" 1 -2 ",
			},
			result: nil,
			err:    true,
		},
		{
			name: "whitespace works",
			input: []string{
				" \t1   2    3 ",
			},
			result: cargo{
				1: []byte{},
				2: []byte{},
				3: []byte{},
			},
			err: false,
		},
		{
			name: "no bays",
			input: []string{
				"",
			},
			result: nil,
			err:    true,
		},
		{
			name: "duplicate bays",
			input: []string{
				" 1 2 1",
			},
			result: nil,
			err:    true,
		},
		{
			name: "valid input with crates",
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3 ",
			},
			result: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
			err: false,
		},
		{
			name: "works with cargo row one less than bay row",
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3",
			},
			result: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
			err: false,
		},
		{
			name: "pads cargo rows",
			input: []string{
				"    [D]",
				"[N] [C]",
				"[Z] [M] [P]",
				" 1   2   3",
			},
			result: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
			err: false,
		},
		{
			name: "floating crate",
			input: []string{
				"    [D]    ",
				"[N]        ",
				"[Z] [M] [P]",
				" 1   2   3 ",
			},
			result: nil,
			err:    true,
		},
		{
			name: "crate in non existent bay",
			input: []string{
				"[Z] [M] [P] [I]",
				" 1   2   3 ",
			},
			result: nil,
			err:    true,
		},
		{
			name: "empty bay",
			input: []string{
				"[Z] [M]    ",
				" 1   2   3 ",
			},
			result: cargo{
				1: []byte{'Z'},
				2: []byte{'M'},
				3: []byte{},
			},
			err: false,
		},
		{
			name: "Non A-Z crate",
			input: []string{
				"[Z] [M] [9]",
				" 1   2   3 ",
			},
			result: nil,
			err:    true,
		},
		{
			name: "junk crate line",
			input: []string{
				"jkljkljl   ",
				"[Z] [M] [9]",
				" 1   2   3 ",
			},
			result: nil,
			err:    true,
		},
		{
			name: "no bracket crate line",
			input: []string{
				" Z  [M] [P]",
				" 1   2   3 ",
			},
			result: nil,
			err:    true,
		},
		{
			name: "mis-aligned crate",
			input: []string{
				" [Z] [M]  [P]",
				" 1    2    3 ",
			},
			result: nil,
			err:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cargo, err := parseCargo(test.input)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, cargo)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.result, cargo)
			}
		})
	}
}
