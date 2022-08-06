package glist

import (
	"github.com/royleo35/go-generics/tools"
	"reflect"
	"testing"
)

func TestGList(t *testing.T) {
	// Create a new list and put some numbers in it.
	l := New[int]()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	want := []int{1, 2, 3, 4}
	for idx, e := 0, l.Front(); e != nil; idx, e = idx+1, e.Next() {
		tools.Assert(e.Value == want[idx])
	}
	want2 := []int{2, 3, 4, 5}
	s := l.ForEach(func(e *int) { *e++ }).ToSlice()
	tools.Assert(reflect.DeepEqual(want2, s))
}
