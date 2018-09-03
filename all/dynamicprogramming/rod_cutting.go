package dynamicprogramming

import (
	"math"

	"github.com/zhiruchen/500Algorithms/common"
)

// RodCutting http://www.techiedelight.com/rot-cutting/
func RodCutting(price []int32, rodLen int) int32 {
	if rodLen == 0 {
		return 0
	}

	var maxValue int32 = math.MinInt32
	for i := 1; i <= rodLen; i++ {
		profit := price[i-1] + RodCutting(price, rodLen-i)

		if profit > maxValue {
			maxValue = profit
		}
	}

	return maxValue
}

func RodCuttingIter(price []int32, rodLen int) int32 {
	t := make([]int32, rodLen+1)

	for i := 1; i <= rodLen; i++ {
		//
		for j := 1; j <= i; j++ {
			t[i] = common.Max(t[i], price[j-1]+t[i-j])
		}
	}

	return t[rodLen]
}
