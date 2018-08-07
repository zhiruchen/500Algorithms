package heap

// CheckIfGivenArrayRepresentsMinHeapOrNot http://www.techiedelight.com/check-given-array-represents-min-heap-not/
func CheckIfGivenArrayRepresentsMinHeapOrNot(array []int32) bool {
	newArray := []int32{-1}
	newArray = append(newArray, array...)
	return checkMinHeap(newArray, 1)
}

func checkMinHeap(array []int32, i int) bool {
	// 如果 i 是一个叶子节点，返回true
	if 2*i+1 > len(array) {
		return true
	}

	var left, right bool
	left = array[i] <= array[2*i] && checkMinHeap(array, 2*i)

	right = 2*i+1 == len(array)
	right = right || (array[i] <= array[2*i+1] && checkMinHeap(array, 2*i+1))

	return left && right
}
