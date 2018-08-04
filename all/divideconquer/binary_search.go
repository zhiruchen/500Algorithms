package divideconquer

// BinarySearch http://www.techiedelight.com/binary-search/
func BinarySearch(array []int32, val int32) int {
	low, high := 0, len(array)-1

	for low <= high {
		mid := low + (high-low)/2

		if array[mid] < val {
			low = mid + 1
		} else if array[mid] > val {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
