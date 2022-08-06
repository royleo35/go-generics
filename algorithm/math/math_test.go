package math

import (
	"github.com/royleo35/go-generics/tools"
	"testing"
)

func TestMath(t *testing.T) {
	a := Max(1, []int{2, 5}...)
	tools.Assert(a == 5)
	min := Min(1, []int{2, 5}...)
	tools.Assert(min == 1)
	sum := Sum(1, []int{2, 5}...)
	tools.Assert(sum == 8)
	avg := Average(1, []int{2, 5}...)
	tools.Assert(avg-2.76 < 0.1)
}

func TestCutInRange(t *testing.T) {
	min := 0
	max := 5
	tools.Assert(CutInRange(2, min, max) == 2)
	tools.Assert(CutInRange(-1, min, max) == min)
	tools.Assert(CutInRange(10, min, max) == max)

}
