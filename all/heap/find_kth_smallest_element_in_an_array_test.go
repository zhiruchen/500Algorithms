package heap

import (
	"testing"
)

func TestFindKthSmallestElementInAnArray(t *testing.T) {
	cases := []struct {
		array []int32
		k     int
		val   int32
	}{
		{
			array: []int32{1, 3, 2, 6, 5, 9, 10, 0},
			k:     5,
			val:   5,
		},
		{
			array: []int32{1, 3, 2, 6, 5, 9, 10, 0},
			k:     1,
			val:   0,
		},
		{
			array: []int32{1, 3, 2, 6, 5, 9, 10, 0},
			k:     3,
			val:   2,
		},
	}

	for _, c := range cases {
		val := FindKthSmallestElementInAnArray(c.array, c.k)
		if val != c.val {
			t.Errorf("Expect: %d. get: %d\n", c.val, val)
		}
	}
}
