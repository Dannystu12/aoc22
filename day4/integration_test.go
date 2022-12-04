package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegration(t *testing.T) {
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	pairs, err := ParseInput(input)
	assert.NoError(t, err)

	t.Run("fully contains", func(t *testing.T) {
		expected := 2
		count := 0
		for _, pair := range pairs {
			if pair.FullyContains() {
				count++
			}
		}

		assert.Equal(t, expected, count)
	})

	t.Run("any overlap", func(t *testing.T) {
		expected := 4
		count := 0
		for _, pair := range pairs {
			if pair.AnyOverlap() {
				count++
			}
		}

		assert.Equal(t, expected, count)
	})

}
