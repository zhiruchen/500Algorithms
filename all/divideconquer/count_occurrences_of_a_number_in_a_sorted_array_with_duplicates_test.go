package divideconquer

import (
	"testing"
)

func TestCountOccurrencesOfANumberInASortedArrayWithDuplicates(t *testing.T) {
	cases := []struct {
		input []int32
		val   int32
		count int32
	}{
		{
			input: []int32{1, 2, 6, 8, 9, 100},
			val:   9,
			count: 1,
		},
		{
			input: []int32{1, 2, 6, 8, 9, 9, 9, 100},
			val:   9,
			count: 3,
		},
		{
			input: []int32{1, 1, 1, 2, 6, 8, 9, 100},
			val:   1,
			count: 3,
		},
		{
			input: []int32{1, 2, 6, 8, 9, 100},
			val:   -1,
			count: -1,
		},
	}

	for _, c := range cases {
		count := CountOccurrencesOfANumberInASortedArrayWithDuplicates(c.input, c.val)
		if count != c.count {
			t.Errorf("Expected %d, get: %d\n", c.count, count)
		}
	}
}
