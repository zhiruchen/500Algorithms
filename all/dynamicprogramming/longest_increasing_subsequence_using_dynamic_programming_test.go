package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncreasingSubsequenceUsingDynamicProgramming(t *testing.T) {
	cases := []struct {
		seq    []int32
		length int
	}{
		{
			seq:    []int32{1, 0, 2, 3, 6, 5},
			length: 4, // 1,2,3,6 | 0,2,3,6
		},
		{
			seq:    []int32{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15},
			length: 6, // 0, 2, 6, 9, 11, 15
		},
	}

	for _, c := range cases {
		length := LongestIncreasingSubsequenceUsingDynamicProgramming(c.seq)
		assert.Equal(t, c.length, length)
	}
}

func TestLongestIncreasingSubsequenceIter(t *testing.T) {
	cases := []struct {
		seq    []int32
		length int
	}{
		{
			seq:    []int32{1, 0, 2, 3, 6, 5},
			length: 4, // 1,2,3,6 | 0,2,3,6
		},
		{
			seq:    []int32{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15},
			length: 6, // 0, 2, 6, 9, 11, 15
		},
	}

	for _, c := range cases {
		length := LongestIncreasingSubsequenceIter(c.seq)
		assert.Equal(t, c.length, length)
	}
}
