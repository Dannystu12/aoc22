package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  []string
		result treeGrid
		err    bool
	}{
		{
			name:   "empty input",
			input:  []string{},
			result: treeGrid{},
			err:    false,
		},
		{
			name:   "nil input",
			input:  nil,
			result: treeGrid{},
			err:    false,
		},
		{
			name: "lines of different sizes",
			input: []string{
				"123456",
				"123",
			},
			result: nil,
			err:    true,
		},
		{
			name: "blank lines",
			input: []string{
				"123456",
				"",
				"123456",
			},
			result: nil,
			err:    true,
		},
		{
			name: "works with whitespace",
			input: []string{
				" 123456",
				"123456\t",
			},
			result: treeGrid{
				[]treeHeight{1, 2, 3, 4, 5, 6},
				[]treeHeight{1, 2, 3, 4, 5, 6},
			},
			err: false,
		},
		{
			name: "invalid chars",
			input: []string{
				"12d3456",
				"1234567",
			},
			result: nil,
			err:    true,
		},
		{
			name: "basic example",
			input: []string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			result: treeGrid{
				[]treeHeight{3, 0, 3, 7, 3},
				[]treeHeight{2, 5, 5, 1, 2},
				[]treeHeight{6, 5, 3, 3, 2},
				[]treeHeight{3, 3, 5, 4, 9},
				[]treeHeight{3, 5, 3, 9, 0},
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
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.result, result)
		})
	}
}
