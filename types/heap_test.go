package types

import (
	"reflect"
	"testing"
)

func TestHeapElements(t *testing.T) {
	cases := []struct {
		array    []int32
		elements []int32
	}{
		{
			array:    []int32{1, 5, 3, 8, 9, 6},
			elements: []int32{9, 8, 6, 1, 5, 3},
		},
	}

	for _, c := range cases {
		h := NewHeap(c.array)
		if !reflect.DeepEqual(h.Elements(), c.elements) {
			t.Errorf("expect: %v, get: %v\n", c.elements, h.Elements())
		}
	}
}
