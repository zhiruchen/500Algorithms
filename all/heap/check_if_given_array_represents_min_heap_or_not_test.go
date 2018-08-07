package heap

import (
	"testing"
)

func TestCheckIfGivenArrayRepresentsMinHeapOrNot(t *testing.T) {
	cases := []struct {
		array     []int32
		isMinHeap bool
	}{
		{
			array:     []int32{1, 3, 5, 9, 8, 6, 7},
			isMinHeap: true,
		},
		{
			array:     []int32{1, 3, 5, 9, 8, 2, 0},
			isMinHeap: false,
		},
	}

	for _, c := range cases {
		isMinHeap := CheckIfGivenArrayRepresentsMinHeapOrNot(c.array)
		if isMinHeap != c.isMinHeap {
			t.Errorf("expect %v, get: %v\n", c.isMinHeap, isMinHeap)
		}
	}
}
