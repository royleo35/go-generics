package hashset

import (
	"fmt"
	"github.com/royleo35/go-generics/gslice"
	"github.com/royleo35/go-generics/hashmap"
)

type HashSet[T any] struct {
	set    *hashmap.HashMap[T, struct{}]
	fmtKey func(T) string
}

func defFmtKey[T comparable](v T) string {
	return fmt.Sprintf("%v", v)
}

func New[T any](s []T, fmtKey func(T) string) *HashSet[T] {
	set := hashmap.NewHashMap[T, struct{}](len(s), fmtKey)
	gslice.ForEach(s, func(v *T) { set.Set(*v, struct{}{}) })
	return &HashSet[T]{set: set, fmtKey: fmtKey}
}

func NewDef[T comparable](s []T) *HashSet[T] {
	return New[T](s, defFmtKey[T])
}

func (s *HashSet[T]) Contains(v ...T) bool {
	for _, val := range v {
		if _, ok := s.set.Get(val); !ok {
			return false
		}
	}
	return true
}

func (s *HashSet[T]) Add(v T) {
	s.set.Set(v, struct{}{})
}

func (s *HashSet[T]) Delete(v T) {
	s.set.Delete(v)
}

func (s *HashSet[T]) Copy() *HashSet[T] {
	return New[T](s.set.Keys(), s.fmtKey)
}

func (s *HashSet[T]) Update(other *HashSet[T]) {
	if other == nil {
		return
	}
	for it := other.set.First(); it != nil; it = it.Next() {
		s.Add(it.Key())
	}
}

func (s *HashSet[T]) Intersection(other *HashSet[T]) *HashSet[T] {
	if other == nil {
		return nil
	}
	if s.set.Size() > other.set.Size() {
		return other.Intersection(s)
	}
	res := New[T](nil, s.fmtKey)
	for it := s.set.First(); it != nil; it = it.Next() {
		if _, ok := other.set.Get(it.Key()); ok {
			res.Add(it.Key())
		}
	}
	return res
}

// Difference s-other, means which contains in s, but not contains in other
func (s *HashSet[T]) Difference(other *HashSet[T]) *HashSet[T] {
	inter := s.Intersection(other)
	res := New[T](nil, s.fmtKey)
	for it := s.set.First(); it != nil; it = it.Next() {
		if _, ok := inter.set.Get(it.Key()); !ok {
			res.Add(it.Key())
		}
	}
	return res
}

func (s *HashSet[T]) Size() int {
	return s.set.Size()
}

func (s *HashSet[T]) ToSlice() []T {
	return s.set.Keys()
}

func (s *HashSet[T]) Equal(other *HashSet[T]) bool {
	if s.set.Size() != other.set.Size() {
		return false
	}
	for it := s.set.First(); it != nil; it = it.Next() {
		if _, ok := other.set.Get(it.Key()); !ok {
			return false
		}
	}
	return true
}
