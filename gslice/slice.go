package gslice

import (
	"github.com/royleo35/go-generics/algorithm/sort"
	"github.com/royleo35/go-generics/gset"
	"github.com/royleo35/go-generics/treemap"
	"github.com/royleo35/go-generics/types"
)

func ReverseByCopy[T any](data []T) []T {
	s := Copy(data)
	sort.Reverse(s)
	return s
}

func Reverse[T any](data []T) {
	sort.Reverse(data)
}

func Unique[T comparable](data []T) []T {
	return gset.New(data).ToSlice()
}

func MergeSlice[T comparable](data1, data2 []T) []T {
	return gset.New(data1).UnsafeMerge(gset.New(data2)).ToSlice()
}

func IntersectionSlice[T comparable](data1, data2 []T) []T {
	return gset.New(data1).Intersection(gset.New(data2)).ToSlice()
}

func DifferenceSlice[T comparable](data1, data2 []T) []T {
	return gset.New(data1).Difference(gset.New(data2)).ToSlice()
}

func SortEqual[T types.BuiltIn](data1, data2 []T) bool {
	if len(data1) != len(data1) {
		return false
	}
	return gset.New(data1).UnsafeEqual(gset.New(data2))
}

func SetEqual[T comparable](data1, data2 []T) bool {
	return gset.New(data1).Equal(gset.New(data2))
}

func Map[T1 any, T2 any](data []T1, do func(e T1) T2) []T2 {
	res := make([]T2, len(data))
	for i := range data {
		res[i] = do(data[i])
	}
	return res
}

func Copy[T any](s []T) []T {
	res := make([]T, len(s))
	copy(res, s)
	return res
	// return append(([]T)(nil), s...)
}

func Connect[T any](s1 []T, other ...[]T) []T {
	res := Copy(s1)
	for _, v := range other {
		res = append(res, v...)
	}
	return res
}

func ToMapValues[T any, K comparable](s []T, f func(T) K) map[K]T {
	res := make(map[K]T, len(s))
	for _, v := range s {
		res[f(v)] = v
	}
	return res
}

func ToTreeMapValues[K types.BuiltIn, V any](values []V, f func(V) K) *treemap.TreeMap[K, V] {
	return treemap.NewDefByValues(values, f)
}

func ToMap[T any, K comparable, V any](s []T, f func(T) (K, V)) map[K]V {
	res := make(map[K]V, len(s))
	for _, v := range s {
		k, val := f(v)
		res[k] = val
	}
	return res
}

// Filter 筛选出f(e) == true 的元素
func Filter[T any](s []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// FilterLen 计算按照条件过滤之后剩余的数量
func FilterLen[T any](s []T, f func(T) bool) int {
	cnt := 0
	for _, v := range s {
		if f(v) {
			cnt++
		}
	}
	return cnt
}

func ElemToSlice[T any](e T) []T {
	return []T{e}
}

func ForEach[T any](data []T, do func(e *T)) {
	for i := range data {
		do(&data[i])
	}
}
