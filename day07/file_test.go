package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilename_validate(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name  string
		input fileName
		err   bool
	}{
		{
			name:  "empty",
			input: "",
			err:   true,
		},
		{
			name:  "whitespace",
			input: "   ",
			err:   true,
		},
		{
			name:  "trailing whitespace",
			input: "abc  ",
			err:   true,
		},
		{
			name:  "space in middle",
			input: "a bc2",
			err:   true,
		},
		{
			name:  "valid lowercase",
			input: "abc",
			err:   false,
		},
		{
			name:  "valid uppercase",
			input: "ABC",
			err:   false,
		},
		{
			name:  "valid numbers",
			input: "123",
			err:   false,
		},
		{
			name:  "valid mix",
			input: "aZ3",
			err:   false,
		},
		{
			name:  "valid single char",
			input: "a",
			err:   false,
		},
		{
			name:  "invalid unicode",
			input: "世界",
			err:   true,
		},
		{
			name:  "valid with extension",
			input: "test.txt",
			err:   false,
		},
		{
			name:  "invalid no extension text",
			input: "test.",
			err:   true,
		},
		{
			name:  "invalid no prefix",
			input: ".test",
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

func TestFileName_getKey(t *testing.T) {
	fn := fileName("foo")
	expected := FsEntryKey("foo")
	assert.Equal(t, expected, fn.getKey())
}

func TestFile_isFile(t *testing.T) {
	t.Parallel()

	f, err := newFile("test", 2)
	assert.NoError(t, err)

	assert.True(t, f.IsFile())
}

func TestFile_isDir(t *testing.T) {
	t.Parallel()

	f, err := newFile("test", 90)
	assert.NoError(t, err)

	assert.True(t, !f.IsDir())
}

func TestFile_getName(t *testing.T) {
	t.Parallel()

	f, err := newFile("test", 33)
	assert.NoError(t, err)

	assert.Equal(t, f.name, f.getName())
}

func TestNewFile(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		fName  fileName
		size   uint
		result *file
		err    bool
	}{
		{
			"valid file name",
			"abcD.123",
			99,
			makePtr(file{
				"abcD.123",
				99,
			}),
			false,
		},
		{
			"invalid file name",
			"abc..sl ",
			99,
			nil,
			true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := newFile(test.fName, test.size)
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

func TestFile_getSize(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name   string
		f      file
		result uint
	}{
		{
			"empty",
			file{},
			0,
		},
		{
			"with size",
			file{
				"foo",
				99,
			},
			99,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := test.f.GetSize()
			assert.Equal(t, test.result, result)
		})
	}
}
