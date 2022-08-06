package rbtree

import (
	"github.com/royleo35/go-generics/algorithm/sort"
	"github.com/royleo35/go-generics/tools"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	s := []int{1, 6, 2, 3, 4, 15, 24, 8}
	tree := NewWithNodes(s, NumberCompare[int])
	want := []int{1, 6, 2, 3, 4, 15, 24, 8}
	sort.Sort(want)
	tree.Print()
	res := tree.ToSlice()
	tools.Assert(reflect.DeepEqual(res, want))
}

func TestDelete(t *testing.T) {
	s := []int{1, 6, 2, 3, 4, 15, 24, 8}
	tree := NewWithNumberNodes(s)
	tree.Print()

	// not exist
	tree.Delete(30)
	tree.Print()

	tree.Delete(6)
	tree.Print()
}

func TestIsPowerOf2(t *testing.T) {
	tools.Assert(isPowerOf2(1) && isPowerOf2(2) && isPowerOf2(4) && isPowerOf2(8))
	tools.Assert(!isPowerOf2(3) && !isPowerOf2(5) && !isPowerOf2(6) && !isPowerOf2(7))
}
