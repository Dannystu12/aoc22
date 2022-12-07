package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name   string
		input  []string
		result *simpleFS
		err    bool
	}{
		{
			name:  "empty",
			input: []string{},
			result: &simpleFS{
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			err: false,
		},
		{
			name:   "empty line",
			input:  []string{"  "},
			result: nil,
			err:    true,
		},
		{
			name:   "unsupported command",
			input:  []string{"$ cat"},
			result: nil,
			err:    true,
		},
		{
			name:   "bad cd",
			input:  []string{"$ cd .."},
			result: nil,
			err:    true,
		},
		{
			name:   "bad cd 2",
			input:  []string{"$ cd foo"},
			result: nil,
			err:    true,
		},
		{
			name:  "ls works",
			input: []string{"$ ls"},
			result: &simpleFS{
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			err: false,
		},
		{
			name:   "must be in ls mode to ls",
			input:  []string{"dir foo"},
			result: nil,
			err:    true,
		},
		{
			name:   "cant have duplicates",
			input:  []string{"$ ls", "dir foo", "123 foo"},
			result: nil,
			err:    true,
		},
		{
			name:   "junk",
			input:  []string{"fdsdf"},
			result: nil,
			err:    true,
		},
		{
			name:   "no cd",
			input:  []string{"$ cd"},
			result: nil,
			err:    true,
		},
		{
			name:   "too many cd",
			input:  []string{"$ cd test llll"},
			result: nil,
			err:    true,
		},
		{
			name:   "too many ls",
			input:  []string{"$ ls lkjd"},
			result: nil,
			err:    true,
		},
		{
			name:   "ls mode reset on command",
			input:  []string{"$ ls", "dir foo", "$ cd foo", "dir bar"},
			result: nil,
			err:    true,
		},
		{
			name:   "invalid dir entry line",
			input:  []string{"$ ls", "dir foo cat", "$ cd foo", "dir bar"},
			result: nil,
			err:    true,
		},
		{
			name:   "invalid file entry line",
			input:  []string{"$ ls", "123 foo cat", "$ cd foo", "dir bar"},
			result: nil,
			err:    true,
		},
		{
			name:   "invalid file entry size",
			input:  []string{"$ ls", "ddd file"},
			result: nil,
			err:    true,
		},
		{
			name: "full example",
			input: []string{
				"$ cd /",
				"$ ls",
				"dir a",
				"14848514 b.txt",
				"8504156 c.dat",
				"dir d",
				"$ cd a",
				"$ ls",
				"dir e",
				"29116 f",
				"2557 g",
				"62596 h.lst",
				"$ cd e",
				"$ ls",
				"584 i",
				"$ cd ..",
				"$ cd ..",
				"$ cd d",
				"$ ls",
				"4060174 j",
				"8033020 d.log",
				"5626152 d.ext",
				"7214296 k",
			}, result: &simpleFS{
				entries: fsEntryMap{
					"a": &dir{
						name: "a",
						entries: fsEntryMap{
							"e": &dir{
								"e",
								fsEntryMap{
									"i": &file{name: "i", size: 584},
								},
							},
							"f":     &file{name: "f", size: 29116},
							"g":     &file{name: "g", size: 2557},
							"h.lst": &file{name: "h.lst", size: 62596},
						},
					},
					"b.txt": &file{name: "b.txt", size: 14848514},
					"c.dat": &file{name: "c.dat", size: 8504156},
					"d": &dir{
						name: "d",
						entries: fsEntryMap{
							"j":     &file{name: "j", size: 4060174},
							"d.log": &file{name: "d.log", size: 8033020},
							"d.ext": &file{name: "d.ext", size: 5626152},
							"k":     &file{name: "k", size: 7214296},
						},
					},
				},
				currentLocation: []dirName{"d"},
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
