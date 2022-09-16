package gheap

import "testing"

func TestMinHeap(t *testing.T) {
	h := NewDefMinHeap[int]()
	h.PushSlice([]int{1, 3, 2, 4, 5})
	for h.Len() > 0 {
		println(h.Pop())
	}
	// out: 1 2 3 4 5
}

func TestMaxHeap(t *testing.T) {
	h := NewDefMaxHeap[int]()
	h.PushSlice([]int{1, 3, 2, 4, 5})
	for h.Len() > 0 {
		println(h.Pop())
	}
	// out: 5 4 3 2 1
}
