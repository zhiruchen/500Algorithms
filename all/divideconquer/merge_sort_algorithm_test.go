package divideconquer

import (
	"reflect"
	"testing"
)

func TestMergeSortAlgorithm(t *testing.T) {
	cases := []struct {
		array  []int32
		result []int32
	}{
		{
			array:  []int32{1, 3, 6, 5, 9, 2},
			result: []int32{1, 2, 3, 5, 6, 9},
		},
	}

	for _, c := range cases {
		MergeSortAlgorithm(&c.array)
		if !reflect.DeepEqual(c.array, c.result) {
			t.Errorf("expect: %v, get: %v\n", c.result, c.array)
		}
	}
}
