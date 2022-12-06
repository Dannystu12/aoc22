package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegration(t *testing.T) {
	t.Parallel()
	input := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	duplicates := []byte{
		'p',
		'L',
		'P',
		'v',
		't',
		's',
	}

	scores := []int{
		16,
		38,
		42,
		22,
		20,
		19,
	}

	rucksacks, err := ParseInput(input)
	assert.NoError(t, err)

	for i, rucksack := range rucksacks {
		i := i
		rucksack := rucksack
		t.Run(input[i], func(t *testing.T) {
			t.Parallel()
			duplicate, err := rucksack.GetDuplicate()
			assert.NoError(t, err)
			assert.Equal(t, string(duplicates[i]), string(*duplicate))
			score, err := GetPriorityScore(*duplicate)
			assert.NoError(t, err)
			assert.Equal(t, scores[i], score)
		})

	}

}

func TestParseInput(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		input  []string
		result []rucksack
		err    bool
	}{
		{
			name:   "empty input",
			input:  []string{},
			result: []rucksack{},
			err:    false,
		},
		{
			name:   "nil input",
			input:  nil,
			result: []rucksack{},
			err:    false,
		},
		{
			name: "valid input",
			input: []string{
				"vJwwpW",
				"jzRRNq",
			},
			result: []rucksack{
				{
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
				{
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
			},
			err: false,
		},
		{
			name: "odd line input",
			input: []string{
				"vJrwpW",
				"jqHRNqZ",
			},
			result: nil,
			err:    true,
		},
		{
			name: "works with whitespace",
			input: []string{
				"vJwwpW ",
				"\tjzRRNq",
			},
			result: []rucksack{
				{
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
				{
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
			},
			err: false,
		},
		{
			name: "works with whitespace",
			input: []string{
				"vJwwpW ",
				"\tjzRRNq",
			},
			result: []rucksack{
				{
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
				{
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
			},
			err: false,
		},
		{
			name: "0 overlap",
			input: []string{
				"vJrwpW",
			},
			result: nil,
			err:    true,
		},
		{
			name: "multiple overlap",
			input: []string{
				"vpWwpW",
			},
			result: nil,
			err:    true,
		},
		{
			name: "skip empty lines",
			input: []string{
				"vJwwpW",
				"",
				"jzRRNq",
			},
			result: []rucksack{
				{
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
				{
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
			},
			err: false,
		},
		{
			name: "no runes",
			input: []string{
				"vJww世界",
			},
			result: nil,
			err:    true,
		},
		{
			name: "works with duplicates in line",
			input: []string{
				"vJwwwwpW",
				"jzRRNq",
			},
			result: []rucksack{
				{
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
				{
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
			},
			err: false,
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
				assert.Equal(t, test.result, result)
			}

		})
	}
}
