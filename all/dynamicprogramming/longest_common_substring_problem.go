package dynamicprogramming

// LongestCommonSubstringProblem http://www.techiedelight.com/longest-common-substring-problem/
func LongestCommonSubstringProblem(a, b string) int {
	lengthTable := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			lengthTable[i] = append(lengthTable[i], 0)
		}
	}

	var l = 0

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				if i >= 1 && j >= 1 {
					lengthTable[i][j] = lengthTable[i-1][j-1] + 1
				} else {
					lengthTable[i][j] = 1
				}

				if lengthTable[i][j] > l {
					l = lengthTable[i][j]
				}
			}
		}
	}

	return l
}
