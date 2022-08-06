package gmap

import (
	"github.com/royleo35/go-generics/gslice"
)

func Copy[Key comparable, Val any](m map[Key]Val) map[Key]Val {
	res := make(map[Key]Val)
	for k, v := range m {
		res[k] = v
	}
	return res
}

// Merge m1 and m2
func Merge[Key comparable, Val any](m1, m2 map[Key]Val) map[Key]Val {
	res := make(map[Key]Val)
	for _, m := range [...]map[Key]Val{m1, m2} {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// Update m2 into m1
func Update[Key comparable, Val any](m1, m2 map[Key]Val) {
	if m1 == nil {
		m1 = make(map[Key]Val)
	}
	for k, v := range m2 {
		m1[k] = v
	}
}

func Keys[Key comparable, Val any](m map[Key]Val) []Key {
	res := make([]Key, len(m))
	idx := 0
	for k := range m {
		res[idx] = k
		idx++
	}
	return res
}

func Values[Key comparable, Val any](m map[Key]Val) []Val {
	res := make([]Val, len(m))
	idx := 0
	for _, v := range m {
		res[idx] = v
		idx++
	}
	return res
}

// Equal calc if keys of two map is same
func KeysEqual[Key comparable, Val any](m1, m2 map[Key]Val) bool {
	if len(m1) != len(m2) {
		return false
	}
	return gslice.SetEqual(Keys(m1), Keys(m2))
}

func Convert[K1, K2 comparable, V1, V2 any](m1 map[K1]V1, f func(K1, V1) (K2, V2)) map[K2]V2 {
	res := make(map[K2]V2, len(m1))
	for k, v := range m1 {
		k2, v2 := f(k, v)
		res[k2] = v2
	}
	return res
}

func ConvertValues[K1 comparable, V1, V2 any](m1 map[K1]V1, f func(V1) V2) map[K1]V2 {
	res := make(map[K1]V2, len(m1))
	for k, v := range m1 {
		res[k] = f(v)
	}
	return res
}
