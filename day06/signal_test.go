package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignalIsValid(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		signal Signal
		valid  bool
	}{
		{
			name:   "empty Signal",
			signal: Signal(""),
			valid:  true,
		},
		{
			name:   "valid Signal",
			signal: Signal("kljljkkljlkjkl"),
			valid:  true,
		},
		{
			name:   "no uppercase ",
			signal: Signal("kljljkklJlkjkl"),
			valid:  false,
		},
		{
			name:   "no whitespace",
			signal: Signal("kljljkk ljlkjkl"),
			valid:  false,
		},
		{
			name:   "no whitespace",
			signal: Signal("\tkljljkkljlkjkl"),
			valid:  false,
		},
		{
			name:   "no punc",
			signal: Signal("kljljkkljlk;jkl"),
			valid:  false,
		},
		{
			name:   "no utf8",
			signal: Signal("kljljkkljl你好kjkl"),
			valid:  false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ok := test.signal.IsValid()
			assert.Equal(t, test.valid, ok)
		})
	}
}

func TestSignalGetStartOfPacketMarker(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		signal Signal
		marker int
		ok     bool
	}{
		{
			name:   "empty Signal",
			signal: Signal(""),
			marker: 0,
			ok:     false,
		},
		{
			name:   "invalid Signal",
			signal: Signal("jkljkljkl;jklj lkjkljK\t"),
			marker: 0,
			ok:     false,
		},
		{
			name:   "valid Signal 1",
			signal: Signal("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			marker: 5,
			ok:     true,
		},
		{
			name:   "valid Signal 2",
			signal: Signal("nppdvjthqldpwncqszvftbrmjlhg"),
			marker: 6,
			ok:     true,
		},
		{
			name:   "valid Signal 3",
			signal: Signal("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			marker: 10,
			ok:     true,
		},
		{
			name:   "valid Signal 4",
			signal: Signal("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			marker: 11,
			ok:     true,
		},
		{
			name:   "too short",
			signal: Signal("zcf"),
			marker: 0,
			ok:     false,
		},
		{
			name:   "min valid",
			signal: Signal("zcfq"),
			marker: 4,
			ok:     true,
		},
		{
			name:   "long invalid",
			signal: Signal("zzzzzzz"),
			marker: 0,
			ok:     false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			marker, ok := test.signal.GetStartOfPacketMarker()
			assert.Equal(t, test.marker, marker)
			assert.Equal(t, test.ok, ok)
		})
	}

}

func TestSignalGetStartOfMessageMarker(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		signal Signal
		marker int
		ok     bool
	}{
		{
			name:   "empty Signal",
			signal: Signal(""),
			marker: 0,
			ok:     false,
		},
		{
			name:   "invalid Signal",
			signal: Signal("jkljkljkl;jklj lkjkljK\t"),
			marker: 0,
			ok:     false,
		},
		{
			name:   "valid Signal 1",
			signal: Signal("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			marker: 23,
			ok:     true,
		},
		{
			name:   "valid Signal 2",
			signal: Signal("nppdvjthqldpwncqszvftbrmjlhg"),
			marker: 23,
			ok:     true,
		},
		{
			name:   "valid Signal 3",
			signal: Signal("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			marker: 29,
			ok:     true,
		},
		{
			name:   "valid Signal 4",
			signal: Signal("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			marker: 26,
			ok:     true,
		},
		{
			name:   "valid Signal 5",
			signal: Signal("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			marker: 19,
			ok:     true,
		},
		{
			name:   "too short",
			signal: Signal("abcdefghijklm"),
			marker: 0,
			ok:     false,
		},
		{
			name:   "min valid",
			signal: Signal("abcdefghijklmn"),
			marker: 14,
			ok:     true,
		},
		{
			name:   "long invalid",
			signal: Signal("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"),
			marker: 0,
			ok:     false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			marker, ok := test.signal.GetStartOfMessageMarker()
			assert.Equal(t, test.marker, marker)
			assert.Equal(t, test.ok, ok)
		})
	}

}
