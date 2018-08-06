package heap

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// Sort http://www.techiedelight.com/heap-sort-place-place-implementation-c-c/
func Sort(array []int32) []int32 {
	return heapSort(array)
}

func heapSort(array []int32) []int32 {
	h := types.NewHeap(array)
	arrLen := len(array)

	var result = make([]int32, arrLen)
	for i := 1; i <= arrLen; i++ {
		result[i-1] = h.Pop()
	}

	return result
}
