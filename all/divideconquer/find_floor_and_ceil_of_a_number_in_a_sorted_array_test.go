package divideconquer

import (
	"testing"
)

func TestFindFloorAndCeilOfANumberInASortedArray(t *testing.T) {
	cases := []struct {
		array []int32
		val   int32
		ceil  int32
		floor int32
	}{
		{
			array: []int32{1, 4, 6, 8, 9},
			val:   0,
			ceil:  1,
			floor: -1,
		},
		{
			array: []int32{1, 4, 6, 8, 9},
			val:   1,
			ceil:  1,
			floor: 1,
		},
		{
			array: []int32{1, 4, 6, 8, 9},
			val:   2,
			ceil:  4,
			floor: 1,
		},
		{
			array: []int32{1, 4, 6, 8, 9},
			val:   4,
			ceil:  4,
			floor: 4,
		},
		{
			array: []int32{1, 4, 6, 8, 9},
			val:   5,
			ceil:  6,
			floor: 4,
		},
	}

	for _, c := range cases {
		ceil, floor := FindFloorAndCeilOfANumberInASortedArray(c.array, c.val)
		if ceil != c.ceil || c.floor != floor {
			t.Errorf("expect ceil: %d, get: %d; expect floor: %d, get: %d\n", c.ceil, ceil, c.floor, floor)
		}
	}
}
