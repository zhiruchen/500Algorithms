package divideconquer

import (
	"testing"
)

func TestFindMissingTermInASequenceInLogNTime(t *testing.T) {
	cases := []struct {
		array []int32
		val   int32
	}{
		{
			array: []int32{1, 2, 3, 4, 6},
			val:   5,
		},
		{
			array: []int32{2, 4, 6, 8, 10, 14},
			val:   12,
		},
		{
			array: []int32{1, 3, 5, 7, 11},
			val:   9,
		},
		{
			array: []int32{3, 6, 9, 12, 18},
			val:   15,
		},
	}

	for _, c := range cases {
		val := FindMissingTermInASequenceInLogNTime(c.array)
		if val != c.val {
			t.Errorf("expect %d, get: %d\n", c.val, val)
		}
	}
}
