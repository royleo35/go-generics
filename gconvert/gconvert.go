// gconvert 用于实现泛型slice之内 和 map 之内的互相转换，支持链式操作

package gconvert

import (
	"github.com/royleo35/go-generics/gmap"
	"github.com/royleo35/go-generics/gslice"
)

type MapCtx[K comparable, V any] struct {
	m map[K]V
}

type SliceCtx[T any] struct {
	s []T
}

func NewFromSlice[T any](s []T) *SliceCtx[T] {
	return &SliceCtx[T]{s}
}

func NewFromMap[K comparable, V any](m map[K]V) *MapCtx[K, V] {
	return &MapCtx[K, V]{m}
}

func (s *SliceCtx[T]) Filter(f func(v T) bool) *SliceCtx[T] {
	s.s = gslice.Filter(s.s, f)
	return s
}

func (s *SliceCtx[T]) ForEach(f func(v *T)) *SliceCtx[T] {
	gslice.ForEach(s.s, f)
	return s
}

func (s *SliceCtx[T]) Reverse() *SliceCtx[T] {
	gslice.Reverse(s.s)
	return s
}

func (s *SliceCtx[T]) Merge(s2 []T) *SliceCtx[T] {
	s.s = append(s.s, s2...)
	return s
}

func (s *SliceCtx[T]) Copy() *SliceCtx[T] {
	res := NewFromSlice(gslice.Copy(s.s))
	return res
}

func (s *SliceCtx[T]) Connect(others []T) *SliceCtx[T] {
	s.s = gslice.Connect(s.s, others)
	return s
}

func (s *SliceCtx[T]) Do() []T {
	return s.s
}

func (m *MapCtx[K, V]) Copy() *MapCtx[K, V] {
	return NewFromMap(gmap.Copy(m.m))
}

func (m *MapCtx[K, V]) Merge(m2 map[K]V) *MapCtx[K, V] {
	gmap.Update(m.m, m2)
	return m
}

func (m *MapCtx[K, V]) Keys() []K {
	return gmap.Keys(m.m)
}

func (m *MapCtx[K, V]) Values() []V {
	return gmap.Values(m.m)
}

func (m *MapCtx[K, V]) Do() map[K]V {
	return m.m
}
