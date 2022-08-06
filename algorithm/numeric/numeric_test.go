package numeric

import (
	"fmt"
	"testing"
)

func TestPower(t *testing.T) {
	res := Power(2, 7)
	res1 := Power(2, 8)
	res2 := Power(2, 10)

	fmt.Println(res, res1, res2)
}
