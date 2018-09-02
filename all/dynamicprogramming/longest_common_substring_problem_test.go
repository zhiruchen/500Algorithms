package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubstringProblem(t *testing.T) {
	cases := []struct {
		s1  string
		s2  string
		lcs int
	}{
		{
			s1:  "BAC",
			s2:  "BABA",
			lcs: 2,
		},
		{
			s1:  "ABABC",
			s2:  "BABCA",
			lcs: 4,
		},
	}

	for _, c := range cases {
		len := LongestCommonSubstringProblem(c.s1, c.s2)
		assert.Equal(t, c.lcs, len)
	}
}
