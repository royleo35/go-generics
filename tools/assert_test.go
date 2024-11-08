package tools

import (
	"fmt"
	"testing"
)

func TestAssert(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	Assert(false)
}

func TestAssertEq(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	AssertEq(10, int64(10))
}
