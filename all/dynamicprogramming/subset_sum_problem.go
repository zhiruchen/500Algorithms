package dynamicprogramming

// SubsetSumProblem http://www.techiedelight.com/subset-sum-problem/
func SubsetSumProblem(array []int32, sum int32) ([]int32, bool) {
	set := &[]int32{}
	r := findSumSet(array, sum, len(array)-1, set)
	if r {
		return *set, true
	}

	return nil, false
}

func findSumSet(arr []int32, sum int32, n int, set *[]int32) bool {
	// base case
	if sum == 0 {
		return true
	}

	if n < 0 || sum < 0 {
		return false
	}

	// 包含当前item
	*set = append(*set, arr[n])
	include := findSumSet(arr, sum-arr[n], n-1, set)
	if include {
		return true
	}

	// 不包含item
	removeFromSlice(set, arr[n])
	exclude := findSumSet(arr, sum, n-1, set)
	return exclude
}

func removeFromSlice(arr *[]int32, val int32) {
	for i, v := range *arr {
		if v == val {
			left := (*arr)[i+1:]
			*arr = (*arr)[0:i]
			*arr = append(*arr, left...)
		}
	}
}
