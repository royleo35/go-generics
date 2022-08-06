package treemap

import (
	"fmt"
	"github.com/royleo35/go-generics/tools"
	"reflect"
	"testing"
)

func TestTreeMapGetSet(t *testing.T) {
	tm := NewDef[int, string]()
	tm.Set(10, "hi")
	v, ok := tm.Get(10)
	tools.Assert(ok && v == "hi")

	tm.Set(10, "haha")
	v, ok = tm.Get(10)
	tools.Assert(ok && v == "haha")
}

func TestTreeMapKeysValues(t *testing.T) {
	tm := New[int, string](defCmpKey[int])
	tm.Set(10, "10")
	tm.Set(11, "100")
	tm.Set(12, "12")
	keys := tm.Keys()
	values := tm.Values()
	tools.Assert(reflect.DeepEqual(keys, []int{10, 11, 12}))
	tools.Assert(reflect.DeepEqual(values, []string{"10", "100", "12"}))
}

func TestTreeMapIter(t *testing.T) {
	tm := New[int, string](defCmpKey[int])
	tm.Set(10, "10")
	tm.Set(11, "100")
	tm.Set(18, "18")
	tm.Set(12, "12")
	keys := []int{10, 11, 12, 18}
	values := []string{"10", "100", "12", "18"}
	idx := 0
	for it := tm.First(); it != nil; it = it.Next() {
		fmt.Printf("k: %v, v: %v\n", it.Key(), it.Value())
		tools.Assert(it.Key() == keys[idx] && it.Value() == values[idx])
		idx++
	}
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
	tm := NewDefByValues(vs, func(v struct {
		id   int64
		name string
	}) int64 {
		return v.id
	})
	v, ok := tm.Get(1)
	tools.Assert(ok && v.name == "hello")
}
