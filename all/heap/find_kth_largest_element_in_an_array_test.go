package heap

import (
	"testing"
)

func TestFindKthLargestElementInAnArray(t *testing.T) {
	cases := []struct {
		array []int32
		k     int
		val   int32
	}{
		{
			array: []int32{1, 3, 2, 6, 5, 9, 10, 0},
			k:     5,
			val:   3,
		},
		{
			array: []int32{1, 3, 2, 6, 5, 9, 10, 0},
			k:     1,
			val:   10,
		},
		{
			array: []int32{1, 3, 2, 6, 5, 9, 10, 0},
			k:     3,
			val:   6,
		},
	}

	for _, c := range cases {
		val := FindKthLargestElementInAnArray(c.array, c.k)
		if val != c.val {
			t.Errorf("Expect: %d. get: %d\n", c.val, val)
		}
	}
}
