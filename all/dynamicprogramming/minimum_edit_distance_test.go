package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumEditDistance(t *testing.T) {
	s1 := "abcde"
	s2 := "xyz"
	d := 5
	assert.Equal(t, d, MinimumEditDistance(s1, s2))

	s1 = "abc"
	s2 = "abc"
	d = 0
	assert.Equal(t, d, MinimumEditDistance(s1, s2))

	s1 = "abc"
	s2 = "abz"
	d = 1
	assert.Equal(t, d, MinimumEditDistance(s1, s2))

	s1 = "great"
	s2 = "gzreyatly"
	d = 4
	assert.Equal(t, d, MinimumEditDistance(s1, s2))
}
