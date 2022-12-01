package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		result []elfInventory
		err    bool
	}{
		{
			name:   "empty list",
			input:  []string{},
			result: []elfInventory{},
			err:    false,
		},
		{
			name:   "single elf one item",
			input:  []string{"1000"},
			result: []elfInventory{{1000}},
			err:    false,
		},
		{
			name:   "single elf two items",
			input:  []string{"1000", "2000"},
			result: []elfInventory{{1000, 2000}},
			err:    false,
		},
		{
			name:   "multiple elfs",
			input:  []string{"1000", "2000", "", "7000", "8000"},
			result: []elfInventory{{1000, 2000}, {7000, 8000}},
			err:    false,
		},
		{
			name:   "non int input",
			input:  []string{"1000", "2000", "", "7000", "foo"},
			result: nil,
			err:    true,
		},
		{
			name:   "multiple blank lines",
			input:  []string{"1000", "2000", "", "", "7000", "8000"},
			result: []elfInventory{{1000, 2000}, {7000, 8000}},
			err:    false,
		},
		{
			name:   "whitespace empty line lines",
			input:  []string{"1000", "2000", "     \t", "7000", "8000"},
			result: []elfInventory{{1000, 2000}, {7000, 8000}},
			err:    false,
		},
		{
			name:   "closing empty line",
			input:  []string{"1000", "2000", "", "", "7000", "8000", ""},
			result: []elfInventory{{1000, 2000}, {7000, 8000}},
			err:    false,
		},
		{
			name:   "bigger example",
			input:  []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"},
			result: []elfInventory{{1000, 2000, 3000}, {4000}, {5000, 6000}, {7000, 8000, 9000}, {10_000}},
			err:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := parseInput(test.input)
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
