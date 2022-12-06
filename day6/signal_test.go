package day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignalIsValid(t *testing.T) {
	var tests = []struct {
		name   string
		signal Signal
		valid  bool
	}{
		{
			name:   "empty Signal",
			signal: Signal(""),
			valid:  false,
		},
		{
			name:   "too short",
			signal: Signal("abc"),
			valid:  false,
		},
		{
			name:   "right size",
			signal: Signal("abcd"),
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
		t.Run(test.name, func(t *testing.T) {
			ok := test.signal.IsValid()
			assert.Equal(t, test.valid, ok)
		})
	}
}

func TestSignalGetMarker(t *testing.T) {
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
		t.Run(test.name, func(t *testing.T) {
			marker, ok := test.signal.GetMarker()
			assert.Equal(t, test.marker, marker)
			assert.Equal(t, test.ok, ok)
		})
	}

}
