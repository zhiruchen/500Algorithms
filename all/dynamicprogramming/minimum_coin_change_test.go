package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumCoinChange(t *testing.T) {
	coins := []int{1, 5, 10}
	amount := 20
	num := 2
	assert.Equal(t, num, MinimumCoinChange(coins, amount))

	coins = []int{1, 5, 10}
	amount = 5
	num = 1
	assert.Equal(t, num, MinimumCoinChange(coins, amount))

	coins = []int{1, 3, 4}
	amount = 6
	num = 2
	assert.Equal(t, num, MinimumCoinChange(coins, amount))
}
