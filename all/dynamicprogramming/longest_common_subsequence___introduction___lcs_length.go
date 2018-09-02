package dynamicprogramming

import (
	"github.com/zhiruchen/500Algorithms/common"
)

// LongestCommonSubsequenceIntroductionLCSLength http://www.techiedelight.com/longest-common-subsequence/
func LongestCommonSubsequenceIntroductionLCSLength(x, y string) int {
	return lcs(x, y, len(x), len(y))
}

func lcs(x, y string, m, n int) int {
	// base case
	if m == 0 || n == 0 {
		return 0
	}

	if x[m-1] == y[n-1] {
		return lcs(x, y, m-1, n-1) + 1
	}

	return common.IntMax(lcs(x, y, m, n-1), lcs(x, y, m-1, n))
}

func LongestCommonSubsequenceLengthIter(x, y string) int {
	lengthTable := make([][]int, len(x))
	for i := range lengthTable {
		for j := 0; j < len(y); j++ {
			lengthTable[i] = append(lengthTable[i], 0)
		}
	}

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			if x[i] == y[j] {
				if i >= 1 && j >= 1 {
					lengthTable[i][j] = lengthTable[i-1][j-1] + 1
				} else {
					lengthTable[i][j] = 1
				}
			} else {
				if i >= 1 && j >= 1 {
					lengthTable[i][j] = common.IntMax(lengthTable[i-1][j], lengthTable[i][j-1])
				}
			}
		}
	}

	return lengthTable[len(x)-1][len(y)-1]
}
