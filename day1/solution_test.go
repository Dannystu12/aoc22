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
			name:   "non int inventory",
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

func TestGetMaxCalories(t *testing.T) {

	var tests = []struct {
		name   string
		input  []elfInventory
		result *int
	}{
		{
			name:   "empty list",
			input:  []elfInventory{},
			result: nil,
		},
		{
			name:   "nil inventory",
			input:  nil,
			result: nil,
		},
		{
			name:   "one elf one item",
			input:  []elfInventory{{1000}},
			result: makeIntPtr(1000),
		},
		{
			name:   "one elf multiple items",
			input:  []elfInventory{{1000, 2000}},
			result: makeIntPtr(3000),
		},
		{
			name:   "nil elf",
			input:  []elfInventory{{1000, 2000}, nil, {3000, 5000}},
			result: makeIntPtr(8000),
		},
		{
			name:   "bigger example",
			input:  []elfInventory{{1000, 2000, 3000}, {4000}, {5000, 6000}, {7000, 8000, 9000}, {10_000}},
			result: makeIntPtr(24000),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetMaxCalories(test.input)
			if test.result == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *test.result, *result)
			}

		})
	}
}

func TestGetMaxCaloriesTopN(t *testing.T) {

	var tests = []struct {
		name      string
		inventory []elfInventory
		n         uint
		result    *int
	}{
		{
			name:      "empty list",
			inventory: []elfInventory{},
			n:         3,
			result:    nil,
		},
		{
			name:      "nil inventory",
			inventory: nil,
			n:         3,
			result:    nil,
		},
		{
			name:      "one elf one item top 1",
			inventory: []elfInventory{{1000}},
			n:         1,
			result:    makeIntPtr(1000),
		},
		{
			name:      "one elf one item top 3",
			inventory: []elfInventory{{1000}},
			n:         3,
			result:    makeIntPtr(1000),
		},
		{
			name:      "top 0",
			inventory: []elfInventory{{1000}},
			n:         0,
			result:    nil,
		},
		{
			name:      "one elf multiple items",
			inventory: []elfInventory{{1000, 2000}},
			n:         3,
			result:    makeIntPtr(3000),
		},
		{
			name:      "nil elf",
			inventory: []elfInventory{{1000, 2000}, nil, {3000, 5000}},
			n:         3,
			result:    makeIntPtr(11000),
		},
		{
			name:      "bigger example",
			inventory: []elfInventory{{1000, 2000, 3000}, {4000}, {5000, 6000}, {7000, 8000, 9000}, {10_000}},
			n:         3,
			result:    makeIntPtr(45000),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetMaxCaloriesTopN(test.inventory, test.n)
			if test.result == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *test.result, *result)
			}
		})
	}
}

func makeIntPtr(i int) *int {
	return &i
}
