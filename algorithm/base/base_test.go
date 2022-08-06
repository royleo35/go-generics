package base

import (
	"fmt"
	"github.com/royleo35/go-generics/tools"
	"testing"
)

func TestEqual(t *testing.T) {
	s1 := []int{1, 2}
	s2 := Copy(s1)
	p1 := &s1
	p2 := &s2
	tools.Assert(Equal(s1, s2, func(a, b int) bool { return a == b }) && (p1 != p2))
}

func TestFill(t *testing.T) {
	s1 := make([]int, 10)
	Fill(s1, 1001)
	fmt.Println(s1)
	tools.Assert(s1[1] == 1001)
}

func TestFillN(t *testing.T) {
	a := 1.1
	res := FillN(10, a)
	fmt.Println(res)
	tools.Assert(res[1] == a)
}

func TestLexicoCompare(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"a", "B"}
	less := func(s1, s2 string) bool { return s1 < s2 }
	tools.Assert(!LexicoCompare(a, b, less))
	tools.Assert(LexicoCompare(b, a, less))
}
