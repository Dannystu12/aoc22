package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name     string
		input    []string
		commands []command
		err      bool
	}{
		{
			name:     "empty input",
			input:    []string{},
			commands: []command{},
			err:      false,
		},
		{
			name:     "empty line",
			input:    []string{"   "},
			commands: nil,
			err:      true,
		},
		{
			name:     "invalid command",
			input:    []string{"foo"},
			commands: nil,
			err:      true,
		},
		{
			name:  "noop command",
			input: []string{"noop"},
			commands: []command{
				noopCommand{},
			},
			err: false,
		},
		{
			name:  "addX command",
			input: []string{"addx 5"},
			commands: []command{
				addXCommand{value: 5},
			},
			err: false,
		},
		{
			name:  "addX command negative",
			input: []string{"addx -5"},
			commands: []command{
				addXCommand{value: -5},
			},
			err: false,
		},
		{
			name:     "addX command invalid arg",
			input:    []string{"addx dfsdfs"},
			commands: nil,
			err:      true,
		},
		{
			name:     "addX command no arg",
			input:    []string{"addx "},
			commands: nil,
			err:      true,
		},
		{
			name:  "basic example",
			input: []string{"noop", "addx 3", "addx -5"},
			commands: []command{
				noopCommand{},
				addXCommand{value: 3},
				addXCommand{value: -5},
			},
			err: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			commands, err := ParseInput(test.input)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.commands, commands)
		})
	}

}
