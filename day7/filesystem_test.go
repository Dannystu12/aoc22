package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleFSEntryMap = fsEntryMap{
	"foo": &dir{
		"foo",
		fsEntryMap{
			"bar": &file{
				"bar",
				90,
			},
			"baz": &file{
				"baz",
				90,
			},
			"quz": &dir{
				"bar",
				fsEntryMap{
					"quz": &file{
						"quz",
						50,
					},
				},
			},
		},
	},
	"bar": &file{
		"bar",
		90,
	},
}

func TestNewSimpleFS(t *testing.T) {
	t.Parallel()
	fs := newSimpleFS()
	expected := simpleFS{
		entries:         make(fsEntryMap),
		currentLocation: []dirName{},
	}
	assert.NotNil(t, fs)
	assert.Equal(t, expected, *fs)
}

func TestSimpleFS_getCurrentDirectoryContents(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name     string
		input    simpleFS
		expected fsEntryMap
		err      bool
	}{
		{
			"empty",
			simpleFS{
				entries:         make(fsEntryMap),
				currentLocation: []dirName{},
			},
			fsEntryMap{},
			false,
		},
		{
			"gets files and directories at root",
			simpleFS{
				entries:         exampleFSEntryMap,
				currentLocation: []dirName{},
			},
			exampleFSEntryMap,
			false,
		},
		{
			"access non existent directory",
			simpleFS{
				entries:         exampleFSEntryMap,
				currentLocation: []dirName{dirName("ljlkjlkjkljklj")},
			},
			nil,
			true,
		},
		{
			"access file as directory",
			simpleFS{
				entries:         exampleFSEntryMap,
				currentLocation: []dirName{dirName("bar")},
			},
			nil,
			true,
		},
		{
			"access nested directory",
			simpleFS{
				entries:         exampleFSEntryMap,
				currentLocation: []dirName{dirName("foo"), dirName("quz")},
			},
			fsEntryMap{"quz": &file{
				"quz",
				50,
			}},
			false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual, err := test.input.getCurrentDirectoryContents()
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, actual)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			}
		})
	}

}

func TestSimpleFS_addEntry(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name     string
		fs       simpleFS
		input    FsEntry
		expected simpleFS
		err      bool
	}{
		{
			"add file",
			simpleFS{
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			&file{
				"test",
				100,
			},
			simpleFS{
				entries: fsEntryMap{"test": &file{
					"test",
					100,
				}},
				currentLocation: []dirName{},
			},
			false,
		},
		{
			"nil entry",
			simpleFS{
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			nil,
			simpleFS{
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			true,
		},
		{
			"cant add duplicate file",
			simpleFS{
				entries: fsEntryMap{"test": &file{
					"test",
					100,
				}},
				currentLocation: []dirName{},
			},
			&file{
				"test",
				123,
			},
			simpleFS{
				entries: fsEntryMap{"test": &file{
					"test",
					100,
				}},
				currentLocation: []dirName{},
			},
			true,
		},
		{
			"add directory",
			simpleFS{
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			&dir{
				"testdir",
				fsEntryMap{
					"test": &file{
						"test",
						100,
					},
				},
			},
			simpleFS{
				entries: fsEntryMap{"testdir": &dir{
					"testdir",
					fsEntryMap{
						"test": &file{
							"test",
							100,
						},
					}}},
				currentLocation: []dirName{},
			},
			false,
		},
		{
			"cant add dir if file name is same",
			simpleFS{
				entries: fsEntryMap{"test": &file{
					"test",
					100,
				}},
				currentLocation: []dirName{},
			},
			&dir{
				"test",
				fsEntryMap{},
			},
			simpleFS{
				entries: fsEntryMap{"test": &file{
					"test",
					100,
				}},
				currentLocation: []dirName{},
			},
			true,
		},
		{
			"cant add file if dir name is same",
			simpleFS{
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{},
				}},
				currentLocation: []dirName{},
			},
			&file{
				"test",
				100,
			},
			simpleFS{
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{},
				}},
				currentLocation: []dirName{},
			},
			true,
		},
		{
			"can add file to sub directory",
			simpleFS{
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{},
				}},
				currentLocation: []dirName{"test"},
			},
			&file{
				"test",
				100,
			},
			simpleFS{
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{"test": &file{
						"test",
						100,
					}},
				}},
				currentLocation: []dirName{"test"},
			},
			false,
		},
		{
			"error on invalid current location",
			simpleFS{
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{},
				}},
				currentLocation: []dirName{"test123"},
			},
			&file{
				"test",
				100,
			},
			simpleFS{
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{},
				}},
				currentLocation: []dirName{"test123"},
			},
			true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.fs.addEntry(test.input)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, test.fs)
		})
	}
}

func TestSimpleFS_cd(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name     string
		fs       simpleFS
		input    string
		expected simpleFS
		err      bool
	}{
		{
			"cant cd to parent when at root",
			simpleFS{
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			"..",
			simpleFS{
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			true,
		},
		{
			"cant cd to root",
			simpleFS{
				entries:         fsEntryMap{"test": &dir{"test", fsEntryMap{}}},
				currentLocation: []dirName{"test"},
			},
			"/",
			simpleFS{
				entries:         fsEntryMap{"test": &dir{"test", fsEntryMap{}}},
				currentLocation: []dirName{},
			},
			false,
		},
		{
			"can cd up a level",
			simpleFS{
				entries:         fsEntryMap{"test": &dir{"test", fsEntryMap{}}},
				currentLocation: []dirName{"test"},
			},
			"..",
			simpleFS{
				entries:         fsEntryMap{"test": &dir{"test", fsEntryMap{}}},
				currentLocation: []dirName{},
			},
			false,
		},
		{
			"cant cd to non existent directory",
			simpleFS{
				entries: fsEntryMap{"test": &file{
					"test",
					100,
				}},
				currentLocation: []dirName{},
			},
			"test",
			simpleFS{
				entries: fsEntryMap{"test": &file{
					"test",
					100,
				}},
				currentLocation: []dirName{},
			},
			true,
		},
		{
			"can cd to directory",
			simpleFS{
				entries:         fsEntryMap{"test": &dir{"test", fsEntryMap{}}},
				currentLocation: []dirName{},
			},
			"test",
			simpleFS{
				entries:         fsEntryMap{"test": &dir{"test", fsEntryMap{}}},
				currentLocation: []dirName{"test"},
			},
			false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.fs.cd(test.input)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, test.fs)
		})
	}
}

func TestSimpleFS_traverse(t *testing.T) {
	count := 0
	exampleFSEntryMap.traverse(func(key FsEntryKey, entry FsEntry) {
		count++
	})

	assert.Equal(t, 6, count)
}
