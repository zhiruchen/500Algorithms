package common

func Max(a, b int32) int32 {
	if a >= b {
		return a
	}

	return b
}

func Min(a, b int32) int32 {
	if a <= b {
		return a
	}

	return b
}

func IntMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func IntsMin(vals ...int) int {
	var min = vals[0]

	for _, val := range vals[1:] {
		if val < min {
			min = val
		}
	}

	return min
}
