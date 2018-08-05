package divideconquer

// MergeSortAlgorithm http://www.techiedelight.com/merge-sort/
func MergeSortAlgorithm(array *[]int32) {
	mergeSort(array, 0, len(*array)-1)
}

func mergeSort(arr *[]int32, low, high int) {
	if high == low {
		return
	}

	mid := low + (high-low)/2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	mergeSubArrays(arr, low, mid, high)
}

func mergeSubArrays(arr *[]int32, low, mid, high int) {
	k, i, j := 0, low, mid+1

	var aux = make([]int32, (high - low + 1))

	for i <= mid && j <= high {
		if (*arr)[i] < (*arr)[j] {
			aux[k] = (*arr)[i]
			i++
		} else {
			aux[k] = (*arr)[j]
			j++
		}
		k++
	}

	for i <= mid {
		aux[k] = (*arr)[i]
		k++
		i++
	}

	for j <= high {
		aux[k] = (*arr)[j]
		k++
		j++
	}

	var x = 0
	for i := low; i <= high; i++ {
		(*arr)[i] = aux[x]
		x++
	}
}
