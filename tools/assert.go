package tools

import (
	"fmt"
	"reflect"
	"runtime"
)

func Assert(ok bool) {
	if !ok {
		_, file, line, _ := runtime.Caller(1)
		s := fmt.Sprintf("%v:%v assert failed.", file, line)
		panic(s)
	}
}

func AssertEq(a, b any) {
	if !reflect.DeepEqual(a, b) {
		_, file, line, _ := runtime.Caller(1)
		s := fmt.Sprintf("%v:%v assert failed. first param [ type: %s, value: %v ], not equal to second param [ type: %s, value: %v ]",
			file, line, reflect.TypeOf(a).String(), a, reflect.TypeOf(b).String(), b)
		panic(s)
	}
}
