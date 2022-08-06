package math

import "github.com/royleo35/go-generics/types"

func Max[T types.Number](v T, more ...T) T {
	m := v
	for _, item := range more {
		if item > m {
			m = item
		}
	}
	return m
}

func Min[T types.Number](v T, more ...T) T {
	m := v
	for _, item := range more {
		if item < m {
			m = item
		}
	}
	return m
}

func CutInRange[T types.Number](v T, min T, max T) T {
	if v < min {
		v = min
	} else if v > max {
		v = max
	}
	return v
}

func Sum[T types.Number](v T, more ...T) T {
	sum := v
	for _, item := range more {
		sum += item
	}
	return sum
}

func Average[T types.Number](v T, more ...T) float64 {
	return float64(Sum(v, more...)) / float64(1+len(more))
}
