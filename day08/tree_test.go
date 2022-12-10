package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeHeight_isValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		height treeHeight
		err    bool
	}{
		{
			name:   "negative is invalid",
			height: -1,
			err:    true,
		},
		{
			name:   "zero is valid",
			height: 0,
			err:    false,
		},
		{
			name:   "9 is valid",
			height: 9,
			err:    false,
		},
		{
			name:   "10 is invalid",
			height: 10,
			err:    true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.height.isValid()
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestTreeGrid_isValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		grid treeGrid
		err  bool
	}{
		{
			name: "empty grid is valid",
			grid: treeGrid{},
			err:  false,
		},
		{
			name: "grid with one empty row",
			grid: treeGrid{[]treeHeight{}},
			err:  false,
		},
		{
			name: "valid grid",
			grid: treeGrid{[]treeHeight{0, 1, 2, 3, 4, 5}, []treeHeight{5, 4, 3, 2, 1, 0}},
			err:  false,
		},

		{
			name: "not all rows same length",
			grid: treeGrid{[]treeHeight{0, 1, 2, 3, 4, 5}, []treeHeight{5, 4, 3, 1, 0}},
			err:  true,
		},
		{
			name: "not all valid heights",
			grid: treeGrid{[]treeHeight{0, 1, 2, 3, 4, 5}, []treeHeight{5, 4, 10, 1, 0}},
			err:  true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.grid.isValid()
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestTreeGrid_GetVisibleTrees(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		grid treeGrid
		want []treeHeight
	}{
		{
			name: "empty grid",
			grid: treeGrid{},
			want: []treeHeight{},
		},
		{
			name: "grid with one empty row",
			grid: treeGrid{[]treeHeight{}},
			want: []treeHeight{},
		},
		{
			name: "one row",
			grid: treeGrid{[]treeHeight{0, 1, 2, 3, 4, 5}},
			want: []treeHeight{0, 1, 2, 3, 4, 5},
		},
		{
			name: "one column",
			grid: treeGrid{
				[]treeHeight{0},
				[]treeHeight{1},
				[]treeHeight{2},
				[]treeHeight{3},
				[]treeHeight{4},
				[]treeHeight{5},
			},
			want: []treeHeight{0, 1, 2, 3, 4, 5},
		},
		{
			name: "returns all outer values",
			grid: treeGrid{
				[]treeHeight{0, 6, 7, 8, 9},
				[]treeHeight{1, 0, 0, 0, 5},
				[]treeHeight{2, 0, 0, 0, 4},
				[]treeHeight{3, 0, 0, 0, 2},
				[]treeHeight{4, 0, 0, 0, 1},
				[]treeHeight{5, 9, 8, 7, 0},
			},
			want: []treeHeight{0, 6, 7, 8, 9, 1, 5, 2, 4, 3, 2, 4, 1, 5, 9, 8, 7, 0},
		},
		{
			name: "basic example",
			grid: treeGrid{
				[]treeHeight{3, 0, 3, 7, 3},
				[]treeHeight{2, 5, 5, 1, 2},
				[]treeHeight{6, 5, 3, 3, 2},
				[]treeHeight{3, 3, 5, 4, 9},
				[]treeHeight{3, 5, 3, 9, 0},
			},
			want: []treeHeight{3, 0, 3, 7, 3, 2, 5, 5, 2, 6, 5, 3, 2, 3, 5, 9, 3, 5, 3, 9, 0},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.grid.GetVisibleTrees()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestTreeGrid_getTree(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		grid treeGrid
		x    int
		y    int
		want treeHeight
		ok   bool
	}{
		{
			name: "empty grid",
			grid: treeGrid{},
			x:    0,
			y:    0,
			want: 0,
			ok:   false,
		},
		{
			name: "empty row",
			grid: treeGrid{[]treeHeight{}},
			x:    0,
			y:    0,
			want: 0,
			ok:   false,
		},
		{
			name: "valid coords",
			grid: treeGrid{[]treeHeight{1}},
			x:    0,
			y:    0,
			want: 1,
			ok:   true,
		},
		{
			name: "negative x",
			grid: treeGrid{[]treeHeight{1}},
			x:    -1,
			y:    0,
			want: 0,
			ok:   false,
		},
		{
			name: "negative y",
			grid: treeGrid{[]treeHeight{1}},
			x:    0,
			y:    -1,
			want: 0,
			ok:   false,
		},
		{
			name: "out of bounds y",
			grid: treeGrid{[]treeHeight{1}},
			x:    0,
			y:    3,
			want: 0,
			ok:   false,
		},
		{
			name: "out of bounds x",
			grid: treeGrid{[]treeHeight{1}},
			x:    3,
			y:    0,
			want: 0,
			ok:   false,
		},
		{
			name: "out of bounds x",
			grid: treeGrid{[]treeHeight{1, 2, 3}, []treeHeight{4, 5, 6}},
			x:    2,
			y:    1,
			want: 6,
			ok:   true,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, ok := test.grid.getTree(test.x, test.y)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.ok, ok)
		})
	}
}

func TestTreeGrid_GetScenicScores(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		grid treeGrid
		want [][]int
	}{
		{
			name: "empty grid",
			grid: treeGrid{},
			want: nil,
		},
		{
			name: "empty row",
			grid: treeGrid{{}},
			want: nil,
		},
		{
			name: "one row",
			grid: treeGrid{{1, 2, 3, 4, 5}},
			want: [][]int{{0, 0, 0, 0, 0}},
		},
		{
			name: "one column",
			grid: treeGrid{
				{1},
				{1}},
			want: [][]int{{0}, {0}},
		},
		{
			name: "2 row 2 column",
			grid: treeGrid{
				{1, 2},
				{2, 1}},
			want: [][]int{{0, 0}, {0, 0}},
		},
		{
			name: "basic example",
			grid: treeGrid{
				[]treeHeight{3, 0, 3, 7, 3},
				[]treeHeight{2, 5, 5, 1, 2},
				[]treeHeight{6, 5, 3, 3, 2},
				[]treeHeight{3, 3, 5, 4, 9},
				[]treeHeight{3, 5, 3, 9, 0},
			},
			want: [][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 4, 1, 0},
				{0, 6, 1, 2, 0},
				{0, 1, 8, 3, 0},
				{0, 0, 0, 0, 0},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := test.grid.GetScenicScores()
			assert.Equal(t, test.want, result)
		})
	}
}
