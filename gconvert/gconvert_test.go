package gconvert

import (
	"github.com/royleo35/go-generics/gslice"
	"github.com/royleo35/go-generics/tools"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4}
	s1 := []int{100, 101}
	sc := NewFromSlice(s)
	// reverse -> 4 3 2 1
	// foreach +1 -> 5 4 3 2
	// filter >2 -> 5 4 3
	// merge s1  -> 5 4 3 100 101
	want := []int{5, 4, 3, 100, 101}
	snew := sc.Reverse().ForEach(func(v *int) { *v++ }).Filter(func(v int) bool { return v > 2 }).Merge(s1).Do()
	tools.Assert(reflect.DeepEqual(snew, want))
}

func TestMap(t *testing.T) {
	// case1 map-> map
	m1 := map[int]string{1: "1", 2: "2"}
	m2 := map[int]string{3: "3"}
	mc := NewFromMap(m1)
	// copy m1 will not change
	// merge -> 1:"1" 2:"2" 3:"3"
	wantm1 := map[int]string{1: "1", 2: "2"}
	want := map[int]string{1: "1", 2: "2", 3: "3"}
	mnew := mc.Copy().Merge(m2).Do()
	tools.Assert(reflect.DeepEqual(mnew, want) && reflect.DeepEqual(m1, wantm1))

	// case2 map->keys
	wantKeys := []int{1, 2, 3}
	keys := NewFromMap(m1).Copy().Merge(m2).Keys()
	tools.Assert(gslice.SortEqual(wantKeys, keys))

	// case2 map->values
	wantValues := []string{"1", "2", "3"}
	values := NewFromMap(m1).Merge(m2).Values()
	tools.Assert(gslice.SortEqual(wantValues, values))

}
