package heap

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// FindKthLargestElementInAnArray http://www.techiedelight.com/find-kth-largest-element-array/
// Find K’th largest element in an array
func FindKthLargestElementInAnArray(array []int32, k int) int32 {
	if len(array) == 1 {
		return array[0]
	}

	h := types.NewHeap(array)

	for i := 1; i <= k-1; i++ {
		h.Pop()
	}

	return h.Pop()
}
