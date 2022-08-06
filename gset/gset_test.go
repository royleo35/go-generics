package gset

import (
	"fmt"
	"github.com/royleo35/go-generics/tools"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := New([]int64{1, 2, 2})
	tools.Assert(s.Contain(2))
	s.Add(3)
	tools.Assert(len(s) == 3)
	s.Delete(3)
	s.Delete(4)
	tools.Assert(len(s) == 2) // 1 2
	c := s.Copy()
	tools.Assert(len(c) == 2)
	data := s.ToSlice()
	fmt.Println(data)
	tools.Assert(len(data) == 2)
	s1 := New([]int64{2, 5})
	s2 := s.UnsafeMerge(s1)
	tools.Assert(len(s2) == 3 && len(s) == 3)
}

func TestSetMethod(t *testing.T) {
	s := New([]int64{1, 2, 2})
	tools.Assert(s.Contain(2))
	s2 := New([]int64{2, 1, 2})
	tools.Assert(s.Equal(s2))

	s = New([]int64{})
	s2 = New([]int64{})
	tools.Assert(s.Equal(s2))

}
