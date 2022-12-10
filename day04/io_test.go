package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name     string
		input    []string
		expected []sectionPair
		err      bool
	}{
		{
			name:     "empty",
			input:    []string{},
			expected: []sectionPair{},
			err:      false,
		},
		{
			name:     "nil",
			input:    nil,
			expected: []sectionPair{},
			err:      false,
		},
		{
			name: "valid input",
			input: []string{
				"2-2,6-8",
				"2-3,4-5",
			},
			expected: []sectionPair{
				{
					A: sectionRange{Min: 2, Max: 2},
					B: sectionRange{Min: 6, Max: 8},
				},
				{
					A: sectionRange{Min: 2, Max: 3},
					B: sectionRange{Min: 4, Max: 5},
				},
			},
			err: false,
		},
		{
			name: "missing comma",
			input: []string{
				"2-46-8",
			},
			expected: nil,
			err:      true,
		},
		{
			name: "missing dash",
			input: []string{
				"2-4,68",
			},
			expected: nil,
			err:      true,
		},
		{
			name: "works with whitespace",
			input: []string{
				"2-4,6-8   ",
				"\t2-3,4-5",
			},
			expected: []sectionPair{
				{
					A: sectionRange{Min: 2, Max: 4},
					B: sectionRange{Min: 6, Max: 8},
				},
				{
					A: sectionRange{Min: 2, Max: 3},
					B: sectionRange{Min: 4, Max: 5},
				},
			},
			err: false,
		},
		{
			name: "second entry must be gte",
			input: []string{
				"4-1,6-8",
				"2-3,4-5",
			},
			expected: nil,
			err:      true,
		},
		{
			name: "no empty lines",
			input: []string{
				"1-1,6-8",
				"",
				"2-3,4-5",
			},
			expected: nil,
			err:      true,
		},
		{
			name: "non int in range",
			input: []string{
				"1-1,6-foo",
			},
			expected: nil,
			err:      true,
		},
		{
			name: "non int in range",
			input: []string{
				"bar-1,6-8",
			},
			expected: nil,
			err:      true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := ParseInput(test.input)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})

	}
}
