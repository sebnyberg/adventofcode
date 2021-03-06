package mathx

import "math"

func Min(as ...int) int {
	if len(as) == 0 {
		panic("at least on number is required")
	}
	min := math.MaxInt64
	for _, a := range as {
		if a < min {
			min = a
		}
	}
	return min
}

func Max(as ...int) int {
	if len(as) == 0 {
		panic("at least on number is required")
	}
	max := math.MinInt64
	for _, a := range as {
		if a > max {
			max = a
		}
	}
	return max
}
