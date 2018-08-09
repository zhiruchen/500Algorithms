package divideconquer

// FindFrequencyOfEachElementInASortedArrayContainingDuplicates http://www.techiedelight.com/find-frequency-element-sorted-array-containing-duplicates/
func FindFrequencyOfEachElementInASortedArrayContainingDuplicates(sortedArray []int32) map[int32]int {
	var freq = make(map[int32]int)
	findFrequency(sortedArray, len(sortedArray), freq)
	return freq
}

func findFrequency(arr []int32, n int, counter map[int32]int) {
	if arr[0] == arr[n-1] {
		counter[arr[0]] += n
		return
	}

	mid := n / 2
	findFrequency(arr, mid, counter)
	findFrequency(arr[mid:], n-mid, counter)
}
