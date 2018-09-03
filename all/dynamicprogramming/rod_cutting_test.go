package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRodCutting(t *testing.T) {
	price := []int32{1, 5, 8, 9, 10, 17, 17, 20}
	rodLen := 4

	maxVal := RodCutting(price, rodLen)
	assert.Equal(t, int32(10), maxVal)
}

func TestRodCuttingIter(t *testing.T) {
	price := []int32{1, 5, 8, 9, 10, 17, 17, 20}
	rodLen := 4

	maxVal := RodCuttingIter(price, rodLen)
	assert.Equal(t, int32(10), maxVal)
}
