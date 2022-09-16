package gheap

import (
	"github.com/royleo35/go-generics/tools"
	"github.com/royleo35/go-generics/types"
)

func MinTopK[T any](s []T, k int, less func(e1, e2 T) bool) []T {
	if k <= 0 {
		panic("k must be > 0")
	}
	pq := &fixedPriorityQueue[T]{
		Heap: New[T](less),
		k:    k,
	}
	pq.PushSlice(s)
	return pq.data
}

func MaxTopK[T any](s []T, k int, less func(e1, e2 T) bool) []T {
	f := func(e1, e2 T) bool {
		return less(e2, e1)
	}
	return MinTopK[T](s, k, f)
}

func MinDefTopK[T types.BuiltIn](s []T, k int) []T {
	return MinTopK[T](s, k, tools.Less[T])
}

func MaxDefTopK[T types.BuiltIn](s []T, k int) []T {
	return MaxTopK[T](s, k, tools.Less[T])
}

type fixedPriorityQueue[T any] struct {
	*Heap[T]
	k int
}

func (pq *fixedPriorityQueue[T]) PushSlice(s []T) {
	for _, v := range s {
		pq.Push(v)
	}
}
func (pq *fixedPriorityQueue[T]) Push(v T) {
	pq.Heap.Push(v)
	for pq.Len() > pq.k {
		pq.Pop()
	}
}

func (pq *fixedPriorityQueue[T]) Pop() T {
	return pq.Heap.Pop()
}
