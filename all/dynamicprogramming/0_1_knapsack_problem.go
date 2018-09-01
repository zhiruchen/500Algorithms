package dynamicprogramming

import (
	"fmt"
	"math"

	"github.com/zhiruchen/500Algorithms/common"
)

// KnapsackProblem http://www.techiedelight.com/0-1-knapsack-problem/
func KnapsackProblem(values, weights []int32, itemNumber int, weightLimit int32) int32 {
	if weightLimit < 0 {
		return math.MinInt32
	}

	if itemNumber < 0 || weightLimit == 0 {
		return 0
	}

	// 包括第n个item + 剩下的item和重量限制(W-W[n])的value
	include := values[itemNumber] + KnapsackProblem(values, weights, itemNumber-1, weightLimit-weights[itemNumber])

	// 排除第n个item的value
	exclude := KnapsackProblem(values, weights, itemNumber-1, weightLimit)

	return common.Max(include, exclude)
}

func KnapsackProblemIterator(values, weights []int32, itemNumber int, weightLimit int32) int32 {
	lookup := make(map[string]int32, len(values))
	return knapsackProblemIterator(values, weights, itemNumber, weightLimit, lookup)
}

func knapsackProblemIterator(values, weights []int32, itemNumber int, weightLimit int32, lookup map[string]int32) int32 {
	if weightLimit < 0 {
		return math.MinInt32
	}

	if itemNumber < 0 || weightLimit == 0 {
		return 0
	}

	key := fmt.Sprintf("%d:%d", itemNumber, weightLimit)
	_, ok := lookup[key]
	if !ok {
		include := values[itemNumber] + knapsackProblemIterator(values, weights, itemNumber-1, weightLimit-weights[itemNumber], lookup)
		exclude := knapsackProblemIterator(values, weights, itemNumber-1, weightLimit, lookup)
		lookup[key] = common.Max(include, exclude)
	}

	return lookup[key]
}
