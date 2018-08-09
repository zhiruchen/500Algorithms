package sorting

// BubbleSortIterativeRecursive http://www.techiedelight.com/bubble-sort-iterative-recursive/
func BubbleSortIterativeRecursive(array []int32) []int32 {
	arrLen := len(array)

	// 外层循环需要跑length-1次，因为当n-1个元素冒到正确的位置的时候，最后一个也到了正确的位置上
	for i := 0; i < arrLen-1; i++ {

		// n 个元素交换次数n-1次, 并且排除已完成排序的i个元素
		for j := 0; j < arrLen-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
