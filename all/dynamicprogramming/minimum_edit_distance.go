package dynamicprogramming

import "github.com/zhiruchen/500Algorithms/common"

// MinimumEditDistance https://www.geeksforgeeks.org/edit-distance-dp-5/
// https://github.com/youngwind/blog/issues/106
func MinimumEditDistance(s1, s2 string) int {
	dist := make([][]int, len(s1)+1)
	for i := range dist {
		for j := 0; j <= len(s2); j++ {
			dist[i] = append(dist[i], 0)
		}
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 {
				// If first string is empty, only option is to
				// insert all characters of second string
				dist[i][j] = j
			} else if j == 0 {
				// If second string is empty, only option is to
				// remove all characters of second string
				dist[i][j] = i
			} else if s1[i-1] == s2[j-1] {
				// If last characters are same, ignore last char
				// and recur for remaining string
				dist[i][j] = dist[i-1][j-1]
			} else {
				// If the last character is different, consider
				// all possibilities and find the minimum
				dist[i][j] = 1 + common.IntsMin(dist[i-1][j], dist[i][j-1], dist[i-1][j-1])
			}
		}
	}

	return dist[len(s1)][len(s2)]
}
