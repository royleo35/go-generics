package gmap

import (
	"fmt"
	"github.com/royleo35/go-generics/tools"
	"testing"
)

func TestKV(t *testing.T) {
	m1 := map[string]string{"1": "2"}
	m2 := map[int]int{3: 4}
	tools.Assert(Keys(m1)[0] == "1")
	tools.Assert(Keys(m2)[0] == 3)
	tools.Assert(Values(m1)[0] == "2")
	tools.Assert(Values(m2)[0] == 4)
}

func TestMerge(t *testing.T) {
	m1 := map[string]string{"1": "2"}
	m2 := map[string]string{"3": "4"}
	m3 := Merge(m1, m2)
	tools.Assert(m3["1"] == "2" && m3["3"] == "4")
	Update(m1, m2)
	tools.Assert(m1["1"] == "2" && m1["3"] == "4")
}

func TestEqual(t *testing.T) {
	m1 := map[string]string{"1": "2", "3": "4"}
	m2 := map[string]string{"1": "2", "3": "4"}
	tools.Assert(KeysEqual(m1, m2))
	m2["5"] = "5"
	tools.Assert(!KeysEqual(m1, m2))

}

func TestCopy(t *testing.T) {
	m1 := map[string]string{"1": "2", "3": "4"}
	m2 := Copy(m1)
	tools.Assert(KeysEqual(m1, m2))
}

func TestConvert(t *testing.T) {
	var v struct{}
	m1 := map[int]struct{}{1: v, 2: v}
	m2 := Convert(m1, func(k1 int, v1 struct{}) (int, string) {
		return k1, fmt.Sprintf("hello_%v", k1)
	})
	tools.Assert(m2[1] == "hello_1" && m2[2] == "hello_2")
}

func TestConvertValues(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2}
	m2 := ConvertValues(m1, func(v1 int) string {
		return fmt.Sprintf("%d", v1)
	})
	tools.Assert(m2[1] == "1" && m2[2] == "2")

}
