package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetSumProblem(t *testing.T) {
	arr := []int32{7, 3, 2, 5, 8}
	var sum int32 = 14

	set, has := SubsetSumProblem(arr, sum)
	assert.True(t, has)
	assert.True(t, arrayEqual([]int32{7, 2, 5}, set))
}

func TestRemoveFromSlice(t *testing.T) {
	vals := []int32{1, 2, 3, 5, 6}
	removeFromSlice(&vals, 3)
	assert.Equal(t, []int32{1, 2, 5, 6}, vals)
}

func arrayEqual(a1, a2 []int32) bool {
	if len(a1) != len(a2) {
		return false
	}

Loop:
	for _, v := range a1 {
		for _, v1 := range a2 {
			if v == v1 {
				continue Loop
			}
		}

		return false
	}

	return true
}
