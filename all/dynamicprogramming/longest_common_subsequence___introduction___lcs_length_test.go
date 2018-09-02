package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequenceIntroductionLCSLength(t *testing.T) {
	cases := []struct {
		x, y   string
		lcsLen int
	}{
		{
			x:      "ABCBDAB",
			y:      "BDCABA",
			lcsLen: 4,
		},
	}

	for _, c := range cases {
		result := LongestCommonSubsequenceIntroductionLCSLength(c.x, c.y)
		assert.Equal(t, c.lcsLen, result)
	}
}
func TestLongestCommonSubsequenceLengthIter(t *testing.T) {
	cases := []struct {
		x, y   string
		lcsLen int
	}{
		{
			x:      "ABCBDAB",
			y:      "BDCABA",
			lcsLen: 4,
		},
	}

	for _, c := range cases {
		result := LongestCommonSubsequenceLengthIter(c.x, c.y)
		assert.Equal(t, c.lcsLen, result)
	}
}
