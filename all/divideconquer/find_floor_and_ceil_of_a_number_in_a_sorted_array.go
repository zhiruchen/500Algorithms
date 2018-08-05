package divideconquer

// FindFloorAndCeilOfANumberInASortedArray http://www.techiedelight.com/find-floor-ceil-number-sorted-array/
// 在一个排序数组里面寻找x的floor和ceil
// floor: 数组里面小于等于x的最大的数, ceil: 数组里面大于等于x的最小的数
func FindFloorAndCeilOfANumberInASortedArray(array []int32, x int32) (floor, ceil int32) {
	low := 0
	high := len(array) - 1
	return getCeil(array, low, high, x), getFloor(array, low, high, x)
}

func getCeil(array []int32, low, high int, x int32) int32 {
	for low <= high {
		if array[low] >= x {
			return array[low]
		}

		if array[high] < x {
			return -1
		}

		mid := low + (high-low)/2
		if array[mid] == x {
			return array[mid]
		} else if array[mid] > x {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return -1
}

func getFloor(array []int32, low, high int, x int32) int32 {
	for low <= high {
		if x < array[low] {
			return -1
		}

		if x >= array[high] {
			return array[high]
		}

		mid := low + (high-low)/2
		if low == mid {
			return array[low]
		}

		if array[mid] == x {
			return array[mid]
		} else if array[mid] < x {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return -1
}
