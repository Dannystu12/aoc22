package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDuplicate(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		input  rucksack
		result *byte
		err    bool
	}{
		{
			name: "valid input",
			input: rucksack{

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
			result: makeBytePtr(byte('w')),
			err:    false,
		},
		{
			name: "case sensitive",
			input: rucksack{

				A: compartment{
					byte('v'): true,
					byte('J'): true,
					byte('w'): true,
				},
				B: compartment{
					byte('z'): true,
					byte('p'): true,
					byte('W'): true,
				},
			},
			result: nil,
			err:    true,
		},
		{
			name: "no duplicates",
			input: rucksack{

				A: compartment{
					byte('v'): true,
					byte('J'): true,
					byte('w'): true,
				},
				B: compartment{
					byte('z'): true,
					byte('p'): true,
					byte('Q'): true,
				},
			},
			result: nil,
			err:    true,
		},
		{
			name: "multiple duplicates",
			input: rucksack{
				A: compartment{
					byte('v'): true,
					byte('J'): true,
					byte('w'): true,
				},
				B: compartment{
					byte('z'): true,
					byte('J'): true,
					byte('w'): true,
				},
			},
			result: nil,
			err:    true,
		},
		{
			name: "nil compartment A",
			input: rucksack{
				A: nil,
				B: compartment{
					byte('z'): true,
					byte('J'): true,
					byte('w'): true,
				},
			},
			result: nil,
			err:    true,
		},
		{
			name: "nil compartment B",
			input: rucksack{
				A: compartment{
					byte('z'): true,
					byte('J'): true,
					byte('w'): true,
				},
				B: nil,
			},
			result: nil,
			err:    true,
		},
		{
			name: "no duplicate found for false compartment A entry",
			input: rucksack{

				A: compartment{
					byte('v'): true,
					byte('J'): true,
					byte('w'): false,
				},
				B: compartment{
					byte('w'): true,
					byte('p'): true,
					byte('W'): true,
				},
			},
			result: nil,
			err:    true,
		},
		{
			name: "no duplicate found for false compartment B entry",
			input: rucksack{

				A: compartment{
					byte('v'): true,
					byte('J'): true,
					byte('w'): false,
				},
				B: compartment{
					byte('w'): false,
					byte('p'): true,
					byte('W'): true,
				},
			},
			result: nil,
			err:    true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := test.input.GetDuplicate()
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, *test.result, *result)
			}
		})
	}
}

func TestGetPriorityScore(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    byte
		expected int
		err      bool
	}{
		{
			name:     "a is 1",
			input:    byte('a'),
			expected: 1,
			err:      false,
		},
		{
			name:     "z is 26",
			input:    byte('z'),
			expected: 26,
			err:      false,
		},
		{
			name:     "A is 27",
			input:    byte('A'),
			expected: 27,
			err:      false,
		},
		{
			name:     "Z is 52",
			input:    byte('Z'),
			expected: 52,
			err:      false,
		},
		{
			name:     ">Z",
			input:    byte('Z') + 1,
			expected: 0,
			err:      true,
		},
		{
			name:     "<A",
			input:    byte('A') - 1,
			expected: 0,
			err:      true,
		},
		{
			name:     "<a",
			input:    byte('a') - 1,
			expected: 0,
			err:      true,
		},
		{
			name:     ">z",
			input:    byte('z') + 1,
			expected: 0,
			err:      true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := GetPriorityScore(test.input)
			if test.err {
				assert.Error(t, err)
				assert.Equal(t, 0, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    rucksack
		expected compartment
	}{
		{
			name: "basic union",
			input: rucksack{
				A: compartment{
					'a': true,
					'X': true,
				},
				B: compartment{
					'v': true,
					'X': true,
				},
			},
			expected: compartment{
				'v': true,
				'a': true,
				'X': true,
			},
		},
		{
			name: "nil A",
			input: rucksack{
				A: nil,
				B: compartment{
					'v': true,
					'X': true,
				},
			},
			expected: compartment{
				'v': true,
				'X': true,
			},
		},
		{
			name: "nil B",
			input: rucksack{
				A: compartment{
					'v': true,
					'X': true,
				},
				B: nil,
			},
			expected: compartment{
				'v': true,
				'X': true,
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := test.input.GetAll()
			assert.Equal(t, test.expected, result)
		})
	}

}

func makeBytePtr(b byte) *byte {
	return &b
}
