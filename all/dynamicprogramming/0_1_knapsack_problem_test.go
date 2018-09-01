package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test01KnapsackProblem(t *testing.T) {
	vals := []int32{20, 5, 10, 40, 15, 25}
	weights := []int32{1, 2, 3, 8, 7, 4}
	var weightLimit int32 = 10
	v := KnapsackProblem(vals, weights, len(vals)-1, weightLimit)

	assert.Equal(t, int32(60), v)
}

func TestKnapsackProblemIterator(t *testing.T) {
	vals := []int32{20, 5, 10, 40, 15, 25}
	weights := []int32{1, 2, 3, 8, 7, 4}
	var weightLimit int32 = 10
	v := KnapsackProblemIterator(vals, weights, len(vals)-1, weightLimit)

	assert.Equal(t, int32(60), v)
}
