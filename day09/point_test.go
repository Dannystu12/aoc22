package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoint_move(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name     string
		p        Point
		d        direction
		expected Point
	}{
		{
			"up",
			Point{0, 0},
			up,
			Point{0, 1},
		},
		{
			"down",
			Point{0, 0},
			down,
			Point{0, -1},
		},
		{
			"left",
			Point{0, 0},
			left,
			Point{-1, 0},
		},
		{
			"right",
			Point{0, 0},
			right,
			Point{1, 0},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.p.move(test.d)
			assert.Equal(t, test.expected, test.p)
		})
	}
}
