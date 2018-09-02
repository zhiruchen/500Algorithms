package dynamicprogramming

import (
	"math"

	"github.com/zhiruchen/500Algorithms/common"
)

// LongestIncreasingSubsequenceUsingDynamicProgramming http://www.techiedelight.com/longest-increasing-subsequence-using-dynamic-programming/
func LongestIncreasingSubsequenceUsingDynamicProgramming(seq []int32) int {
	return longestIncreasingSubsequence(seq, 0, len(seq), math.MinInt32)
}

func longestIncreasingSubsequence(seq []int32, i, n int, prev int32) int {
	if n == i {
		return 0
	}

	// exclude
	excluLcsLen := longestIncreasingSubsequence(seq, i+1, n, prev)

	// include current item
	var includeLCSLen = 0
	if seq[i] > prev {
		includeLCSLen = 1 + longestIncreasingSubsequence(seq, i+1, n, seq[i])
	}

	return common.IntMax(includeLCSLen, excluLcsLen)
}

func LongestIncreasingSubsequenceIter(seq []int32) int {
	lenTable := make([]int, len(seq))
	for i := range lenTable {
		lenTable[i] = 1
	}

	for i := 1; i < len(seq); i++ {
		for j := 0; j < i; j++ {
			if seq[j] <= seq[i] {
				lenTable[i] = common.IntMax(lenTable[i], lenTable[j]+1)
			}
		}
	}

	return lenTable[len(seq)-1]
}
