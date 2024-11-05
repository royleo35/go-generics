package gsync

import (
	"fmt"
	"github.com/royleo35/go-generics/tools"
	"testing"
)

func TestSyncMap(t *testing.T) {
	m := NewSyncMap[string, string](0)
	m.Store("1", "1")
	m.Store("2", "2")
	var val string
	var ok bool
	val, ok = m.Load("1")
	tools.Assert(val == "1" && ok)
	val, ok = m.Load("2")
	tools.Assert(ok && val == "2")
	tools.Assert(m.Delete("1") == true)
	val, ok = m.Load("1")
	tools.Assert(!ok && val == "")

	m.Store("3", "3")
	m1 := map[string]string{}
	m.Range(func(k string, v string) bool {
		m1[k] = v
		return true
	})
	fmt.Println(m1)
	want := map[string]string{
		"2": "2",
		"3": "3",
	}
	tools.AssertEq(m1, want)
	tools.AssertEq(m.CopyMap(), want)

}
