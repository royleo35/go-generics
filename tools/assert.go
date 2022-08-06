package tools

import (
	"fmt"
	"runtime"
)

func Assert(ok bool) {
	if !ok {
		_, file, line, _ := runtime.Caller(1)
		s := fmt.Sprintf("%v:%v assert failed.", file, line)
		panic(s)
	}
}
