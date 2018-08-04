package divideconquer

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		input []int32
		val   int32
		index int
	}{
		{
			input: []int32{1, 2, 6, 8, 9, 100},
			val:   9,
			index: 4,
		},
		{
			input: []int32{1, 2, 6, 8, 9, 100},
			val:   1,
			index: 0,
		},
		{
			input: []int32{1, 2, 6, 8, 9, 100},
			val:   100,
			index: 5,
		},
		{
			input: []int32{1, 2, 6, 8, 9, 100},
			val:   -1,
			index: -1,
		},
	}

	for _, c := range cases {
		index := BinarySearch(c.input, c.val)
		if index != c.index {
			t.Errorf("Expected %d, get: %d\n", c.index, index)
		}
	}
}
