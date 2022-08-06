package base

import (
	"github.com/royleo35/go-generics/algorithm/math"
	"github.com/royleo35/go-generics/types"
)

// Equal 判断两个slice是否相等
func Equal[T any](s1, s2 []T, equal func(e1, e2 T) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if !equal(s1[i], s2[i]) {
			return false
		}
	}
	return true
}

// Fill 将slice每个元素设置为相同的值
func Fill[T types.BuiltIn](data []T, val T) {
	for i := range data {
		data[i] = val
	}
}

func FillN[T types.BuiltIn](n int, val T) []T {
	if n < 0 {
		n = 0
	}
	res := make([]T, n)
	for i := range res {
		res[i] = val
	}
	return res
}

// LexicoCompare 字典序比较
// 如果s1 < s2 返回true, 否则返回false
func LexicoCompare[T any](s1, s2 []T, less func(e1, e2 T) bool) bool {
	l := math.Min(len(s1), len(s2))
	for i := 0; i < l; i++ {
		if less(s1[i], s2[i]) {
			return true
		}
		if less(s2[i], s1[i]) {
			return false
		}
	}
	return len(s1) < len(s2)
}

// Copy拷贝slice, 是对built-in函数copy的简单封装
func Copy[T any](data []T) []T {
	res := make([]T, len(data))
	copy(res, data)
	return res
}
