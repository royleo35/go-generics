package gheap

import (
	"github.com/royleo35/go-generics/tools"
	"github.com/royleo35/go-generics/types"
)

type Heap[T any] struct {
	data []T
	less func(e1, e2 T) bool
}

func NewDefMinHeap[T types.BuiltIn]() *Heap[T] {
	return New[T](tools.Less[T])
}

func NewDefMinHeapWith[T types.BuiltIn](s []T) *Heap[T] {
	h := NewDefMinHeap[T]()
	h.PushSlice(s)
	return h
}

func NewDefMaxHeap[T types.BuiltIn]() *Heap[T] {
	return New[T](tools.Greater[T])
}

func NewDefMaxHeapWith[T types.BuiltIn](s []T) *Heap[T] {
	h := NewDefMaxHeap[T]()
	h.PushSlice(s)
	return h
}

func New[T any](less func(e1, e2 T) bool) *Heap[T] {
	return &Heap[T]{
		data: nil,
		less: less,
	}
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Empty() bool {
	return h.Len() == 0
}

func (h *Heap[T]) Top() T {
	if h.Empty() {
		panic("heap is empty")
	}
	return h.data[0]
}

func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.up()
}

func (h *Heap[T]) PushSlice(s []T) {
	for _, v := range s {
		h.Push(v)
	}
}

func (h *Heap[T]) up() {
	j := len(h.data) - 1
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(h.data[j], h.data[i]) {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		j = i
	}
}
func (h *Heap[T]) Pop() T {
	if len(h.data) == 0 {
		panic("heap is empty!")
	}
	n := len(h.data) - 1
	h.data[0], h.data[n] = h.data[n], h.data[0]
	h.down()
	ret := h.data[n]
	h.data = h.data[:n]
	return ret
}

func (h *Heap[T]) down() {
	i := 0
	n := len(h.data) - 1
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(h.data[j2], h.data[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(h.data[j], h.data[i]) {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}
}
