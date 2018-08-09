package divideconquer

import (
	"reflect"
	"testing"
)

func TestFindFrequencyOfEachElementInASortedArrayContainingDuplicates(t *testing.T) {
	cases := []struct {
		arr  []int32
		freq map[int32]int
	}{
		{
			arr: []int32{1, 2, 2, 3, 5, 5, 6},
			freq: map[int32]int{
				int32(1): 1,
				int32(2): 2,
				int32(3): 1,
				int32(5): 2,
				int32(6): 1,
			},
		},
		{
			arr: []int32{2, 3, 4},
			freq: map[int32]int{
				int32(2): 1,
				int32(3): 1,
				int32(4): 1,
			},
		},
	}

	for _, c := range cases {
		freq := FindFrequencyOfEachElementInASortedArrayContainingDuplicates(c.arr)
		if !reflect.DeepEqual(freq, c.freq) {
			t.Errorf("Expect: %v, get: %v", c.freq, freq)
		}
	}
}
