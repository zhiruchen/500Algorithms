package dynamicprogramming

import (
	"math"
)

// MinimumCoinChange https://www.geeksforgeeks.org/find-minimum-number-of-coins-that-make-a-change/
func MinimumCoinChange(coins []int, amount int) int {
	coinNumbers := make([]int, amount+1)
	coinNumbers[0] = 0
	for i := 1; i < len(coinNumbers); i++ {
		coinNumbers[i] = math.MaxInt32
	}

	for v := 1; v <= amount; v++ {
		for i := 0; i < len(coins); i++ {
			if coins[i] <= v {
				count := coinNumbers[v-coins[i]] + 1
				if count < coinNumbers[v] {
					coinNumbers[v] = count
				}
			}
		}
	}

	return coinNumbers[amount]
}
