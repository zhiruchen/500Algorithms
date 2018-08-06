package heap

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	cases := []struct {
		array  []int32
		result []int32
	}{
		{
			array:  []int32{1, 3, 5, 9, 8, 2, 0},
			result: []int32{9, 8, 5, 3, 2, 1, 0},
		},
	}

	for _, c := range cases {
		result := Sort(c.array)
		if !reflect.DeepEqual(result, c.result) {
			t.Errorf("Expect: %v, get: %v\n", c.result, result)
		}
	}
}
