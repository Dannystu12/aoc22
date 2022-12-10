package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCpu_ProcessCommand(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name     string
		c        cpu
		cmd      command
		expected cpu
	}{
		{
			name: "noop",
			c:    *NewCPU(),
			cmd:  noopCommand{},
			expected: cpu{
				cycle:    1,
				x:        1,
				xHistory: []int{1},
			},
		},
		{
			name: "addX positive",
			c: cpu{
				cycle:    1,
				x:        1,
				xHistory: []int{1},
			},
			cmd: addXCommand{value: 3},
			expected: cpu{
				cycle:    3,
				x:        4,
				xHistory: []int{1, 1, 1},
			},
		},
		{
			name: "addX negative",
			c: cpu{
				cycle:    3,
				x:        4,
				xHistory: []int{1, 1, 1},
			},
			cmd: addXCommand{value: -5},
			expected: cpu{
				cycle:    5,
				x:        -1,
				xHistory: []int{1, 1, 1, 4, 4},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.c.ProcessCommand(test.cmd)
			assert.Equal(t, test.expected, test.c)
		})
	}
}

func TestCpu_GetSignalStrength(t *testing.T) {
	t.Parallel()

	input := []string{
		"addx 15",
		"addx -11",
		"addx 6",
		"addx -3",
		"addx 5",
		"addx -1",
		"addx -8",
		"addx 13",
		"addx 4",
		"noop",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx -35",
		"addx 1",
		"addx 24",
		"addx -19",
		"addx 1",
		"addx 16",
		"addx -11",
		"noop",
		"noop",
		"addx 21",
		"addx -15",
		"noop",
		"noop",
		"addx -3",
		"addx 9",
		"addx 1",
		"addx -3",
		"addx 8",
		"addx 1",
		"addx 5",
		"noop",
		"noop",
		"noop",
		"noop",
		"noop",
		"addx -36",
		"noop",
		"addx 1",
		"addx 7",
		"noop",
		"noop",
		"noop",
		"addx 2",
		"addx 6",
		"noop",
		"noop",
		"noop",
		"noop",
		"noop",
		"addx 1",
		"noop",
		"noop",
		"addx 7",
		"addx 1",
		"noop",
		"addx -13",
		"addx 13",
		"addx 7",
		"noop",
		"addx 1",
		"addx -33",
		"noop",
		"noop",
		"noop",
		"addx 2",
		"noop",
		"noop",
		"noop",
		"addx 8",
		"noop",
		"addx -1",
		"addx 2",
		"addx 1",
		"noop",
		"addx 17",
		"addx -9",
		"addx 1",
		"addx 1",
		"addx -3",
		"addx 11",
		"noop",
		"noop",
		"addx 1",
		"noop",
		"addx 1",
		"noop",
		"noop",
		"addx -13",
		"addx -19",
		"addx 1",
		"addx 3",
		"addx 26",
		"addx -30",
		"addx 12",
		"addx -1",
		"addx 3",
		"addx 1",
		"noop",
		"noop",
		"noop",
		"addx -9",
		"addx 18",
		"addx 1",
		"addx 2",
		"noop",
		"noop",
		"addx 9",
		"noop",
		"noop",
		"noop",
		"addx -1",
		"addx 2",
		"addx -37",
		"addx 1",
		"addx 3",
		"noop",
		"addx 15",
		"addx -21",
		"addx 22",
		"addx -6",
		"addx 1",
		"noop",
		"addx 2",
		"addx 1",
		"noop",
		"addx -10",
		"noop",
		"noop",
		"addx 20",
		"addx 1",
		"addx 2",
		"addx 2",
		"addx -6",
		"addx -11",
		"noop",
		"noop",
		"noop",
	}

	clockCycleSignalStrengths := map[int]int{
		20:  420,
		60:  1140,
		100: 1800,
		140: 2940,
		180: 2880,
		220: 3960,
	}

	c := NewCPU()
	commands, err := ParseInput(input)
	assert.Nil(t, err)

	for _, cmd := range commands {
		c.ProcessCommand(cmd)
	}

	for cycle, strength := range clockCycleSignalStrengths {
		signalStrength, ok := c.GetSignalStrength(cycle)
		assert.True(t, ok)
		assert.Equal(t, strength, signalStrength)
	}

}
