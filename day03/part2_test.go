package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput2(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		input  []string
		result []rucksackGroup
		err    bool
	}{
		{
			name:   "empty input",
			input:  []string{},
			result: []rucksackGroup{},
			err:    false,
		},
		{
			name:   "nil input",
			input:  nil,
			result: []rucksackGroup{},
			err:    false,
		},
		{
			name: "non group of 3",
			input: []string{
				"vJwwpW",
				"jzRRNq",
				"abcDEc",
				"dsfIOd",
			},
			result: nil,
			err:    true,
		},
		{
			name: "valid input",
			input: []string{
				"vJwwpW",
				"jzRRNq",
				"abcDEc",
			},
			result: []rucksackGroup{
				{
					A: rucksack{
						A: compartment{
							byte('v'): true,
							byte('J'): true,
							byte('w'): true,
						},
						B: compartment{
							byte('w'): true,
							byte('p'): true,
							byte('W'): true,
						},
					},
					B: rucksack{
						A: compartment{
							byte('j'): true,
							byte('z'): true,
							byte('R'): true,
						},
						B: compartment{
							byte('R'): true,
							byte('N'): true,
							byte('q'): true,
						},
					},
					C: rucksack{
						A: compartment{
							byte('a'): true,
							byte('b'): true,
							byte('c'): true,
						},
						B: compartment{
							byte('D'): true,
							byte('E'): true,
							byte('c'): true,
						},
					},
				},
			},
			err: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := ParseInput2(test.input)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.result, result)
			}

		})
	}
}

func TestGetDuplicateRucksackGroups(t *testing.T) {
	t.Parallel()

	input := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}

	expected := byte('r')

	rgs, err := ParseInput2(input)
	assert.NoError(t, err)

	dup, err := rgs[0].GetDuplicate()
	assert.NoError(t, err)
	assert.Equal(t, expected, *dup)
}
