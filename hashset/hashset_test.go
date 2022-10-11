package hashset

import (
	"fmt"
	"github.com/royleo35/go-generics/tools"
	"reflect"
	"testing"
)

type InfoTest struct {
	id   int64
	name string
}

func fmtKey(i InfoTest) string {
	return fmt.Sprintf("%v", i.id)
}

func TestHashSetCopyUpdateEqual(t *testing.T) {
	s2 := New[InfoTest]([]InfoTest{{20, "nisfiqe"}}, fmtKey)
	s3 := s2.Copy()
	tools.Assert(s2 != s3 && s2.Equal(s3))
	s4 := New[InfoTest]([]InfoTest{{20, "20"}}, fmtKey)
	s2.Update(s4)
	tools.Assert(reflect.DeepEqual(s2.ToSlice(), []InfoTest{{20, "20"}}))
	s2.Delete(InfoTest{20, ""})
	tools.Assert(s2.Size() == 0)
	s5 := New[InfoTest]([]InfoTest{{50, "50"}}, fmtKey)
	s2.Update(s5)
	tools.Assert(reflect.DeepEqual(s2.ToSlice(), []InfoTest{{50, "50"}}))
}

func TestHashSet2(t *testing.T) {
	l := []InfoTest{
		{10, "hi"},
		{10, "nhif"},
		{20, "roy"},
	}
	s := New[InfoTest](l, fmtKey)
	tools.Assert(s.Size() == 2)
	keys := s.ToSlice()
	fmt.Println(keys)
	tools.Assert(s.Contains(InfoTest{10, "2478235"}))
	s.Add(InfoTest{30, "hihi"})
	tools.Assert(s.Size() == 3)
	s.Add(InfoTest{20, "hihi"})
	tools.Assert(s.Size() == 3)
	s.Delete(InfoTest{30, ""})
	tools.Assert(s.Size() == 2)
	s2 := New[InfoTest]([]InfoTest{{20, "nisfiqe"}}, fmtKey)
	reflect.DeepEqual(s.Intersection(s2).ToSlice(), []InfoTest{{20, "roy"}})
	diff := s.Difference(s2).ToSlice()
	tools.Assert(len(diff) == 1 && diff[0].id == 10 && diff[0].name == "nhif")

}
