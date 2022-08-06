package numeric

import "github.com/royleo35/go-generics/types"

// Accmulate 计算init + sum(data)
func Accmulate[T types.Number](data []T, init T) T {
	for _, item := range data {
		init += item
	}
	return init
}

// AdjacentDifference 计算相邻元素的差
func AdjacentDifference[T types.Number](data []T) []T {
	l := len(data)
	res := make([]T, l)
	if l == 0 {
		return res
	}
	res[0] = data[0]
	for i := 1; i < l; i++ {
		res[i] = data[i] - data[i-1]
	}
	return res
}

// InnerProduct 计算两个向量的内积
func InnerProduct[T types.Number](vec1, vec2 []T) T {
	if len(vec1) != len(vec2) {
		panic("length error")
	}
	res := T(0)
	for i := range vec1 {
		res += vec1[i] * vec2[i]
	}
	return res
}

// PartialSum 计算前i个元素累计和
func PartialSum[T types.Number](data []T) []T {
	l := len(data)
	res := make([]T, len(data))
	if l == 0 {
		return res
	}
	res[0] = data[0]
	for i := 1; i < l; i++ {
		res[i] = res[i-1] + data[i]
	}
	return res
}

// Power 计算幂，即val^n
func Power[T types.Number, Int types.Integer](val T, n Int) T {
	if n < 0 {
		panic("n is must nature number")
	}
	if n == 0 {
		return T(1)
	}
	// n为偶数
	for (n & 1) == 0 {
		n >>= 1
		val = val * val
	}
	res := val
	n >>= 1
	for n != 0 {
		val = val * val
		if (n & 1) != 0 {
			res = res * val
		}
		n >>= 1
	}
	return res
}
