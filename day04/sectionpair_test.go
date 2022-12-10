package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullyContains(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		input  sectionPair
		result bool
	}{
		{
			name: "does not fully contain",
			input: sectionPair{
				A: sectionRange{Min: 1, Max: 2},
				B: sectionRange{Min: 3, Max: 4},
			},
			result: false,
		},
		{
			name: "exact match no range",
			input: sectionPair{
				A: sectionRange{Min: 3, Max: 3},
				B: sectionRange{Min: 3, Max: 3},
			},
			result: true,
		},
		{
			name: "a contains b",
			input: sectionPair{
				A: sectionRange{Min: 3, Max: 6},
				B: sectionRange{Min: 4, Max: 5},
			},
			result: true,
		},
		{
			name: "b contains a",
			input: sectionPair{
				A: sectionRange{Min: 2, Max: 3},
				B: sectionRange{Min: 1, Max: 5},
			},
			result: true,
		},
		{
			name: "partial overlap",
			input: sectionPair{
				A: sectionRange{Min: 2, Max: 3},
				B: sectionRange{Min: 3, Max: 5},
			},
			result: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := test.input.FullyContains()
			assert.Equal(t, test.result, result)
		})
	}
}

func TestAnyOverlap(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		input  sectionPair
		result bool
	}{
		{
			name: "no overlap",
			input: sectionPair{
				A: sectionRange{Min: 1, Max: 2},
				B: sectionRange{Min: 3, Max: 4},
			},
			result: false,
		},
		{
			name: "exact match no range",
			input: sectionPair{
				A: sectionRange{Min: 3, Max: 3},
				B: sectionRange{Min: 3, Max: 3},
			},
			result: true,
		},
		{
			name: "a contains b",
			input: sectionPair{
				A: sectionRange{Min: 3, Max: 6},
				B: sectionRange{Min: 4, Max: 5},
			},
			result: true,
		},
		{
			name: "b contains a",
			input: sectionPair{
				A: sectionRange{Min: 2, Max: 3},
				B: sectionRange{Min: 1, Max: 5},
			},
			result: true,
		},
		{
			name: "partial overlap",
			input: sectionPair{
				A: sectionRange{Min: 2, Max: 3},
				B: sectionRange{Min: 3, Max: 5},
			},
			result: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := test.input.AnyOverlap()
			assert.Equal(t, test.result, result)
		})
	}
}
