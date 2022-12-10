package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongRope_MoveBig(t *testing.T) {
	t.Parallel()

	r, err := NewLongRope(Point{11, 5}, 8)
	assert.NoError(t, err)

	moves, err := ParseInput([]string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	})

	assert.NoError(t, err)

	for _, move := range moves {
		r.Move(move)
	}

	expected := longRope{
		head: ropeHead{
			Point{0, 20},
		},
		knots: []Point{
			{0, 19},
			{0, 18},
			{0, 17},
			{0, 16},
			{0, 15},
			{0, 14},
			{0, 13},
			{0, 12},
		},
		tail: ropeTail{
			Point: Point{0, 11},
			history: map[Point]bool{
				{11, 5}:  true,
				{12, 6}:  true,
				{13, 7}:  true,
				{12, 8}:  true,
				{13, 9}:  true,
				{14, 10}: true,
				{15, 10}: true,
				{16, 10}: true,
				{17, 9}:  true,
				{18, 8}:  true,
				{19, 7}:  true,
				{20, 6}:  true,
				{21, 5}:  true,
				{20, 4}:  true,
				{19, 3}:  true,
				{18, 2}:  true,
				{17, 1}:  true,
				{16, 0}:  true,
				{15, 0}:  true,
				{14, 0}:  true,
				{13, 0}:  true,
				{12, 0}:  true,
				{11, 0}:  true,
				{10, 0}:  true,
				{9, 0}:   true,
				{8, 1}:   true,
				{7, 2}:   true,
				{6, 3}:   true,
				{5, 4}:   true,
				{4, 5}:   true,
				{3, 6}:   true,
				{2, 7}:   true,
				{1, 8}:   true,
				{0, 9}:   true,
				{0, 10}:  true,
				{0, 11}:  true,
			},
		},
	}

	assert.Equal(t, expected, *r)

}

func TestGetNextPos(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name     string
		p1       Point
		p2       Point
		expected Point
	}{
		{
			name:     "no move same pos",
			p1:       Point{3, 3},
			p2:       Point{3, 3},
			expected: Point{3, 3},
		},
		{
			name:     "no move top left",
			p1:       Point{3, 3},
			p2:       Point{2, 4},
			expected: Point{3, 3},
		},
		{
			name:     "no move top middle",
			p1:       Point{3, 3},
			p2:       Point{3, 4},
			expected: Point{3, 3},
		},
		{
			name:     "no move top right",
			p1:       Point{3, 3},
			p2:       Point{4, 4},
			expected: Point{3, 3},
		},
		{
			name:     "no move left",
			p1:       Point{3, 3},
			p2:       Point{2, 3},
			expected: Point{3, 3},
		},
		{
			name:     "no move right",
			p1:       Point{3, 3},
			p2:       Point{4, 3},
			expected: Point{3, 3},
		},
		{
			name:     "no move bottom left",
			p1:       Point{3, 3},
			p2:       Point{2, 2},
			expected: Point{3, 3},
		},
		{
			name:     "no move bottom middle",
			p1:       Point{3, 3},
			p2:       Point{3, 2},
			expected: Point{3, 3},
		},
		{
			name:     "no move bottom right",
			p1:       Point{3, 3},
			p2:       Point{4, 2},
			expected: Point{3, 3},
		},
		{
			name:     "move  right",
			p1:       Point{3, 3},
			p2:       Point{5, 3},
			expected: Point{4, 3},
		},
		{
			name:     "move left",
			p1:       Point{3, 3},
			p2:       Point{1, 3},
			expected: Point{2, 3},
		},
		{
			name:     "move up",
			p1:       Point{3, 3},
			p2:       Point{3, 5},
			expected: Point{3, 4},
		},
		{
			name:     "move down",
			p1:       Point{3, 3},
			p2:       Point{3, 1},
			expected: Point{3, 2},
		},
		{
			name:     "move top right",
			p1:       Point{3, 0},
			p2:       Point{4, 2},
			expected: Point{4, 1},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := getNextPos(test.p1, test.p2)
			assert.Equal(t, test.expected, result)
		})
	}
}
