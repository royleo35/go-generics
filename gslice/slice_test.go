package gslice

import (
	"github.com/royleo35/go-generics/hputil"
	"github.com/royleo35/go-generics/tools"
	"reflect"
	strconv "strconv"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int64{1, 2, 2}
	b := []int64{2, 5}
	m := MergeSlice(a, b)
	tools.Assert(len(m) == 3)
	inter := IntersectionSlice(a, b)
	tools.Assert(len(inter) == 1 && inter[0] == 2)
	diff := DifferenceSlice(a, b)
	tools.Assert(len(diff) == 1 && diff[0] == 1)

	ss := Unique([]string{"a", "a"})
	tools.Assert(len(ss) == 1 && ss[0] == "a")

	s1 := []int{1, 2, 2}
	s2 := []int{2, 1, 2}
	tools.Assert(SetEqual(s1, s2))
}

func TestMap(t *testing.T) {
	// int to string
	a := []int{1, 2, 3}
	res := Map(a, func(e int) string { return strconv.Itoa(e) })
	tools.Assert(len(res) == 3 && res[0] == "1")

	// i*i
	res1 := Map(a, func(i int) int { return i * i })
	tools.Assert(len(res) == 3 && res1[1] == 4)

	// copy
	res2 := Map(a, func(i int) int { return i })
	tools.Assert(len(res) == 3 && res2[2] == 3)
}

func TestConnect(t *testing.T) {
	a := []int{1}
	b := []int{2, 3}
	c := []int{4, 5}
	res := Connect(a, b, c)
	tools.Assert(len(res) == 5 && res[1] == 2)
}

func TestToMapValues(t *testing.T) {
	type Info struct {
		Id   int64
		Name string
	}
	a := []*Info{
		{1, "nihao"},
		{2, "hha"},
	}
	m := ToMapValues(a, func(t *Info) int64 {
		return t.Id
	})
	tools.Assert(m[1].Name == "nihao")
}

func TestFilter(t *testing.T) {
	a := []int{1, 2, 3, 5}
	want := []int{1, 3, 5}
	res := Filter(a, func(e int) bool { return e&1 == 1 })
	tools.Assert(reflect.DeepEqual(res, want))
}

func TestToMap(t *testing.T) {
	type Info struct {
		Id   int64
		Name string
	}
	a := []*Info{
		{1, "nihao"},
		{2, "hha"},
	}
	m := ToMap(a, func(t *Info) (int64, int64) {
		return t.Id, t.Id
	})
	tools.Assert(m[1] == 1)
}

func TestFilterLen(t *testing.T) {
	a := []int{1, 1, 2, 1, 4}
	res := FilterLen(a, func(e int) bool { return e == 1 })
	tools.Assert(res == 3)
}

func TestElemToSlice(t *testing.T) {
	e := 10
	e2 := tools.PtrOf("hello")
	tools.Assert(len(ElemToSlice(e)) == 1 && ElemToSlice(e)[0] == e)
	tools.Assert(len(ElemToSlice(e2)) == 1 && ElemToSlice(e2)[0] == e2)
}

func TestReverse(t *testing.T) {
	a := []int{1, 2, 3}
	want := []int{3, 2, 1}
	res := ReverseByCopy(a)
	tools.Assert(reflect.DeepEqual(want, res))
	Reverse(a)
	tools.Assert(hputil.DeepEqual(want, a))
}

func TestToTreeMapValues(t *testing.T) {
	vs := []struct {
		id   int64
		name string
	}{
		{1, "hello"},
		{3, "nihao"},
		{2, "haha"},
	}
	tm := ToTreeMapValues(vs, func(v struct {
		id   int64
		name string
	}) int64 {
		return v.id
	})
	v, ok := tm.Get(1)
	tools.Assert(ok && v.name == "hello")
}
