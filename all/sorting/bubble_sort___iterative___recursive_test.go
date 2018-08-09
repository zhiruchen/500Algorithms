package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSortIterativeRecursive(t *testing.T) {
	cases := []struct {
		array  []int32
		result []int32
	}{
		{
			array:  []int32{1, 3, 5, 8, 7, 6},
			result: []int32{1, 3, 5, 6, 7, 8},
		},
	}

	for _, c := range cases {
		result := BubbleSortIterativeRecursive(c.array)
		if !reflect.DeepEqual(result, c.result) {
			t.Errorf("expect: %v, get: %v\n", c.result, result)
		}
	}
}
