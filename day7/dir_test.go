package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirname_validate(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name  string
		input dirName
		err   bool
	}{
		{
			name:  "empty",
			input: dirName(""),
			err:   true,
		},
		{
			name:  "whitespace",
			input: dirName("   "),
			err:   true,
		},
		{
			name:  "trailing whitespace",
			input: dirName("abc  "),
			err:   true,
		},
		{
			name:  "space in middle",
			input: dirName("a bc2"),
			err:   true,
		},
		{
			name:  "valid lowercase",
			input: dirName("abc"),
			err:   false,
		},
		{
			name:  "valid uppercase",
			input: dirName("ABC"),
			err:   false,
		},
		{
			name:  "valid numbers",
			input: dirName("123"),
			err:   false,
		},
		{
			name:  "valid mix",
			input: dirName("aZ3"),
			err:   false,
		},
		{
			name:  "valid single char",
			input: dirName("a"),
			err:   false,
		},
		{
			name:  "invalid unicode",
			input: dirName("世界"),
			err:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.input.validate()
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDirName_getKey(t *testing.T) {
	dn := dirName("foo")
	expected := FsEntryKey("foo")
	assert.Equal(t, expected, dn.getKey())
}

func TestDir_isFile(t *testing.T) {
	t.Parallel()

	dir, err := newDir("test")
	assert.NoError(t, err)

	assert.True(t, !dir.IsFile())
}

func TestDir_isDir(t *testing.T) {
	t.Parallel()

	dir, err := newDir("test")
	assert.NoError(t, err)

	assert.True(t, dir.IsDir())
}

func TestDir_getName(t *testing.T) {
	t.Parallel()

	dir, err := newDir("test")
	assert.NoError(t, err)

	assert.Equal(t, dir.name, dir.getName())
}

func TestNewDir(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name    string
		dirName dirName
		result  *dir
		err     bool
	}{
		{
			"valid dir name",
			dirName("abc123"),
			makePtr(dir{
				"abc123",
				make(fsEntryMap),
			}),
			false,
		},
		{
			"invalid dir name",
			dirName("abc 123"),
			nil,
			true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := newDir(test.dirName)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.result, result)
			}
		})
	}
}

func TestDir_getSize(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name   string
		d      dir
		result uint
	}{
		{
			"empty",
			dir{},
			0,
		},
		{
			"with files",
			dir{
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
					"foo": &dir{
						"foo",
						make(fsEntryMap),
					},
				},
			},
			180,
		},
		{
			"with nested dir",
			dir{
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
					"foo": &dir{
						"foo",
						fsEntryMap{
							"bar": &file{
								"bar",
								45,
							},
							"baz": &file{
								"baz",
								50,
							},
						},
					},
				},
			},
			180 + 95,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := test.d.GetSize()
			assert.Equal(t, test.result, result)
		})
	}
}

func makePtr[T any](t T) *T {
	return &t
}
