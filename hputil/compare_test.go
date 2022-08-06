package hputil

import (
	"github.com/royleo35/go-generics/tools"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	// contain nil
	tools.Assert(!DeepEqual(nil, 10) && !reflect.DeepEqual(nil, 10))
	tools.Assert(!DeepEqual(10, nil) && !reflect.DeepEqual(10, nil))
	tools.Assert(DeepEqual(nil, nil) && reflect.DeepEqual(nil, nil))

	// not same type
	a := []int{1, 2}
	b := []int{1, 2}
	type info struct {
		Id   int
		Name *string
	}
	i := &info{
		Id:   100,
		Name: tools.PtrOf("hello"),
	}
	i2 := &info{
		Id:   100,
		Name: tools.PtrOf("hello"),
	}
	tools.Assert(!DeepEqual(a, 10) && !reflect.DeepEqual(a, 10))
	tools.Assert(!DeepEqual(a, i) && !reflect.DeepEqual(a, i))

	// equal
	tools.Assert(DeepEqual(a, b) && reflect.DeepEqual(a, b))
	tools.Assert(DeepEqual(i, i2) && reflect.DeepEqual(i, i2))

	// large data
	ld := [1]byte{}
	ld2 := [1 << 20]byte{}
	tools.Assert(!DeepEqual(&ld, &ld2) && !reflect.DeepEqual(&ld, &ld2))

}
