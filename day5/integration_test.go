package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegration(t *testing.T) {
	t.Parallel()
	input :=
		`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	cargo, moves, err := ParseInput(input)
	assert.NoError(t, err)

	for _, m := range moves {
		err = cargo.PerformMove(m, false)
		assert.NoError(t, err)
	}

	message := cargo.GetMessage()
	assert.Equal(t, "CMZ", message)

	cargo, moves, err = ParseInput(input)
	assert.NoError(t, err)

	for _, m := range moves {
		err = cargo.PerformMove(m, true)
		assert.NoError(t, err)
	}

	message = cargo.GetMessage()
	assert.Equal(t, "MCD", message)
}
