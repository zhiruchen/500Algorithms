package divideconquer

// CountOccurrencesOfANumberInASortedArrayWithDuplicates http://www.techiedelight.com/count-occurrences-number-sorted-array-duplicates/
// 在一个排序数组中指定的数字出现的次数
func CountOccurrencesOfANumberInASortedArrayWithDuplicates(array []int32, val int32) int32 {
	first := binarySearch(array, val, true)
	last := binarySearch(array, val, false)
	if first == -1 {
		return -1
	}

	return int32(last-first) + 1
}

func binarySearch(array []int32, val int32, searchFirst bool) int {
	low, high := 0, len(array)-1
	var result = -1

	for low <= high {
		mid := low + (high-low)/2

		if array[mid] < val {
			low = mid + 1
		} else if array[mid] > val {
			high = mid - 1
		} else {
			result = mid
			if searchFirst {
				high = mid - 1 // 如果搜索第一次出现的位置， 则在左边搜索
			} else {
				low = mid + 1 // 否则在👉搜索
			}
		}
	}

	return result
}
