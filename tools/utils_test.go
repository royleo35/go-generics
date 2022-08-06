package tools

import (
	"fmt"
	"reflect"
	"testing"
)

type MyCos struct {
	a int
	b string
	c *float64
}

func TestIfThen(t *testing.T) {
	a := 10
	b := IfThen(a > 5, "ok", "bad")
	Assert(b == "ok")
	fmt.Println(b)
}

func TestForEach(t *testing.T) {
	a := []int{1, 2, 3}
	addOne := func(i *int) {
		(*i)++
	}
	ForEach(a, addOne)
	fmt.Println(a)
	Assert(reflect.DeepEqual(a, []int{2, 3, 4}))

	b := struct {
		s string
	}{"hds"}
	c := &b
	d := &c
	Assert(b.s == "hds")
	(*(*d)).s = "1"
	Assert(b.s == "1")
	(*d).s = "2" // ok <=>  (*(*d)).s
	Assert(b.s == "2")
	//d.s = "3" // 编译通不过 ，二维指针至少需要解引用一次
	//Assert(b.s == "3")

	ss := []*MyCos{
		{
			a: 1,
			b: "23",
			c: nil,
		},
	}
	ForEach(ss, func(e **MyCos) {
		(*e).a++
	})
	Assert(ss[0].a == 2)
}

func TestPtrOf(t *testing.T) {
	a := 10
	b := 10.1
	c := "sdhiaf"
	d := &MyCos{
		a: 1,
		b: "2",
		c: &b,
	}
	Assert(*PtrOf(a) == a && *PtrOf(b) == b && *PtrOf(c) == c && *PtrOf(d) == d)
}

func TestDeref(t *testing.T) {
	a := 10
	b := "djis"
	type Info struct {
		Id   int
		Name string
	}
	i := Info{
		Id:   10,
		Name: "he",
	}
	Assert(Dereference(&a) == 10)
	Assert(Dereference((*int)(nil)) == 0)
	Assert(Dereference(&b) == b)
	Assert(Dereference((*string)(nil)) == "")
	Assert(Dereference(&i) == i)
	Assert(Dereference((*Info)(nil)) == Info{})
}

func TestPtrCast(t *testing.T) {
	a := 10
	type MyInt int
	_ = PtrCast[int, MyInt](&a)
}
