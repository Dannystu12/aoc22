package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerformMove(t *testing.T) {
	var tests = []struct {
		name     string
		cargo    cargo
		move     move
		err      bool
		endState cargo
	}{
		{
			name: "valid single move",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
			move: move{
				number: 1,
				from:   1,
				to:     3,
			},

			err: false,
			endState: cargo{
				1: []byte{'Z'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P', 'N'},
			},
		},
		{
			name: "move to same bay",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
			move: move{
				number: 1,
				from:   1,
				to:     1,
			},

			err: false,
			endState: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
		},
		{
			name: "valid multi move",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
			move: move{
				number: 3,
				from:   2,
				to:     1,
			},

			err: false,
			endState: cargo{
				1: []byte{'Z', 'N', 'D', 'C', 'M'},
				2: []byte{},
				3: []byte{'P'},
			},
		},
		{
			name: "no crate move",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
			move: move{
				number: 1,
				from:   3,
				to:     1,
			},
			err: true,
			endState: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
		},
		{
			name: "move too many",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
			move: move{
				number: 4,
				from:   2,
				to:     1,
			},
			err: true,
			endState: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
		},
		{
			name: "move from non existent bay",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
			move: move{
				number: 1,
				from:   4,
				to:     1,
			},
			err: true,
			endState: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
		},
		{
			name: "move to non existent bay",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
			move: move{
				number: 1,
				from:   1,
				to:     4,
			},
			err: true,
			endState: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
		},
		{
			name: "zero move",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
			move: move{
				number: 0,
				from:   1,
				to:     3,
			},
			err: false,
			endState: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.cargo.PerformMove(test.move, false)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.endState, test.cargo)
		})
	}

	t.Run("preserved order", func(t *testing.T) {

		c := cargo{
			1: []byte{'Z', 'N'},
			2: []byte{'M', 'C', 'D'},
			3: []byte{'P'},
		}
		m := move{
			number: 3,
			from:   2,
			to:     1,
		}
		endState := cargo{
			1: []byte{'Z', 'N', 'M', 'C', 'D'},
			2: []byte{},
			3: []byte{'P'},
		}
		err := c.PerformMove(m, true)
		assert.NoError(t, err)
		assert.Equal(t, c, endState)

	})
}

func TestGetMessage(t *testing.T) {
	var tests = []struct {
		name    string
		cargo   cargo
		message string
	}{
		{
			name: "basic cargo",
			cargo: cargo{1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'}},
			message: "NDP",
		},
		{
			name:    "empty cargo",
			cargo:   cargo{},
			message: "",
		},
		{
			name: "empty bay",
			cargo: cargo{
				1: []byte{'Z', 'N', 'M', 'C', 'D'},
				2: []byte{},
				3: []byte{'P'},
			},
			message: "DP",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			message := test.cargo.GetMessage()
			assert.Equal(t, test.message, message)
		})
	}
}
