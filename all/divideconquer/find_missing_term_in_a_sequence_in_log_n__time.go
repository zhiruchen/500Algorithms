package divideconquer

// FindMissingTermInASequenceInLogNTime http://www.techiedelight.com/find-missing-term-sequence-ologn-time/
// 在等差数列里找到丢失的数字
func FindMissingTermInASequenceInLogNTime(array []int32) int32 {
	low := 0
	high := len(array) - 1
	d := (array[high] - array[0]) / int32(len(array))

	for low <= high {
		mid := low + (high-low)/2

		if (array[mid] + d) != array[mid+1] {
			return array[mid] + d
		}

		if (array[mid] - d) != array[mid-1] {
			return array[mid] - d
		}

		if (int32(mid) * d) != (array[mid] - array[0]) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
