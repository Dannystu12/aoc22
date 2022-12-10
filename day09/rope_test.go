package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRope_Move(t *testing.T) {
	t.Parallel()

	r := NewRope()
	moves, err := ParseInput([]string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	})
	assert.NoError(t, err)

	for _, move := range moves {
		r.Move(move)
	}

	expected := rope{
		ropeHead{Point{2, 2}},
		ropeTail{
			Point{1, 2},
			map[Point]bool{
				{0, 0}: true,
				{1, 0}: true,
				{2, 0}: true,
				{3, 0}: true,
				{4, 1}: true,
				{4, 2}: true,
				{3, 2}: true,
				{2, 2}: true,
				{1, 2}: true,
				{4, 3}: true,
				{3, 3}: true,
				{3, 4}: true,
				{2, 4}: true,
			},
		},
	}

	assert.Equal(t, expected, r)
}
