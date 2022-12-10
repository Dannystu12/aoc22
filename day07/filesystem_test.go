package day07

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
	fs := newSimpleFS(4000)
	expected := simpleFS{
		capacity:        4000,
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
				capacity:        100000,
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			&file{
				"test",
				100,
			},
			simpleFS{
				capacity: 100000,
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
				capacity:        100000,
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			nil,
			simpleFS{
				capacity:        100000,
				entries:         fsEntryMap{},
				currentLocation: []dirName{},
			},
			true,
		},
		{
			"cant add duplicate file",
			simpleFS{
				capacity: 100000,
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
				capacity: 100000,
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
				capacity:        100000,
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
				capacity: 100000,
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
				capacity: 100000,
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
				capacity: 100000,
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
				capacity: 100000,
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
				capacity: 100000,
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{},
				}},
				currentLocation: []dirName{},
			},
			true,
		},
		{
			"cant add file if it would be over capacity",
			simpleFS{
				capacity: 99,
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{},
				}},
				currentLocation: []dirName{},
			},
			&file{
				"testfile",
				100,
			},
			simpleFS{
				capacity: 99,
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{},
				}},
				currentLocation: []dirName{},
			},
			true,
		},
		{
			"can add file to capacity",
			simpleFS{
				capacity: 99,
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{},
				}},
				currentLocation: []dirName{"test"},
			},
			&file{
				"testfile",
				99,
			},
			simpleFS{
				capacity: 99,
				entries: fsEntryMap{"test": &dir{
					"test",
					fsEntryMap{
						"testfile": &file{
							"testfile",
							99,
						},
					},
				}},
				currentLocation: []dirName{"test"},
			},
			false,
		},
		{
			"can add file to sub directory",
			simpleFS{
				capacity: 100000,
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
				capacity: 100000,
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
				capacity: 100000,
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
				capacity: 100000,
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
	t.Parallel()
	count := 0
	exampleFSEntryMap.traverse(func(key FsEntryKey, entry FsEntry) {
		count++
	})

	assert.Equal(t, 6, count)
}

func TestSimpleFS_getSize(t *testing.T) {
	t.Parallel()
	fs := newSimpleFS(100)
	assert.Equal(t, uint(0), fs.getSize())

	err := fs.addEntry(&file{
		name: "foo",
		size: 10,
	})
	assert.NoError(t, err)
	assert.Equal(t, uint(10), fs.getSize())
	err = fs.addEntry(&dir{
		name: "testdir",
		entries: fsEntryMap{
			"foo": &file{
				name: "foo",
				size: 1,
			},
			"bar": &file{
				name: "foo",
				size: 2,
			},
			"baz": &file{
				name: "foo",
				size: 3,
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, uint(16), fs.getSize())
}

func TestSimpleFS_RecommendDirectoryForDeletion(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name     string
		fs       simpleFS
		required uint
		expected *dir
		err      bool
	}{
		{
			"required space greater than capacity",
			simpleFS{
				entries:         fsEntryMap{"test": &dir{"test", fsEntryMap{}}},
				currentLocation: []dirName{"test"},
				capacity:        100,
			},
			1000,
			nil,
			true,
		},
		{
			"no results",
			simpleFS{
				entries:         fsEntryMap{"test": &dir{"test", fsEntryMap{}}, "foo": &file{name: "foo", size: 99}},
				currentLocation: []dirName{"test"},
				capacity:        100,
			},
			99,
			nil,
			true,
		},
		{
			"already enough space",
			simpleFS{
				entries:         fsEntryMap{"test": &dir{"test", fsEntryMap{}}, "foo": &file{name: "foo", size: 99}},
				currentLocation: []dirName{"test"},
				capacity:        100,
			},
			1,
			nil,
			true,
		},
		{
			"gets smallest",
			simpleFS{
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
				currentLocation: []dirName{},
				capacity:        70000000,
			},
			30000000,
			&dir{
				name: "d",
				entries: fsEntryMap{
					"j":     &file{name: "j", size: 4060174},
					"d.log": &file{name: "d.log", size: 8033020},
					"d.ext": &file{name: "d.ext", size: 5626152},
					"k":     &file{name: "k", size: 7214296},
				},
			},
			false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			dir, err := test.fs.RecommendDirectoryForDeletion(test.required)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, dir)
		})
	}
}
