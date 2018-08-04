package divideconquer

// EfficientlyImplementPowerFunctionRecursiveAndIterative http://www.techiedelight.com/power-function-implementation-recursive-iterative/
func EfficientlyImplementPowerFunctionRecursiveAndIterative(x int, n uint) int64 {
	return power(x, n)
}

func power(x int, n uint) int64 {
	if n == 0 {
		return 1
	}

	pow := power(x, n/2)
	if n%2 != 0 {
		return int64(x) * pow * pow
	}

	return pow * pow
}
