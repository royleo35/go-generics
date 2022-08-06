package sort

import (
	"reflect"
	"testing"

	"github.com/royleo35/go-generics/tools"
)

func TestSort(t *testing.T) {
	s := []int{1, 3, 2}
	Sort(s)
	tools.Assert(reflect.DeepEqual(s, []int{1, 2, 3}))
	SortDesc(s)
	tools.Assert(reflect.DeepEqual(s, []int{3, 2, 1}))
	SortWrap(s)
	tools.Assert(reflect.DeepEqual(s, []int{1, 2, 3}))
}

func TestSortByCopy(t *testing.T) {
	s := []int{1, 3, 2}
	scopy := []int{1, 3, 2}
	want := []int{1, 2, 3}
	res := SortByCopyDef(s)
	tools.Assert(reflect.DeepEqual(s, scopy) && reflect.DeepEqual(res, want))
	res = SortByCopy(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	tools.Assert(reflect.DeepEqual(s, scopy) && reflect.DeepEqual(res, want))
}
