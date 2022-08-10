package rbtree

import (
	"fmt"
	"github.com/royleo35/go-generics/algorithm/sort"
	"github.com/royleo35/go-generics/internal"
	"github.com/royleo35/go-generics/tools"
	"reflect"
	"testing"
	"unsafe"
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

func efaceData(v any) []byte {
	res := *(*[16]byte)(unsafe.Pointer(&v))
	return res[:]
}

func TestNewDef(t *testing.T) {
	m := NewDef[any]()
	var a = any(1)
	var b = any(int64(2))
	var c = any(float32(3.1))
	aa := efaceData(a)
	bb := efaceData(b)
	cc := efaceData(c)
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	tools.Assert(internal.MemCmpValue(b, a) > 0 && internal.MemCmpValue(a, c) > 0)
	m.Insert(a)
	m.Insert(b)
	m.Insert(c)
	m.Print()
	tools.Assert(m.Contain(a) && m.Contain(b) && m.Contain(c))
}
