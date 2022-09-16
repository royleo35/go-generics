package gheap

import (
	"fmt"
	"math/rand"
	"testing"
)

var s = func() []int {
	var res = make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		res = append(res, i)
	}
	return res
}()

func TestMinTopK(t *testing.T) {
	rand.Shuffle(100, func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	res := MinDefTopK(s, 10)
	fmt.Println(res)
	// [90 93 91 94 95 92 96 98 99 97]
}

func TestMaxTopK(t *testing.T) {
	rand.Shuffle(100, func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	res := MaxDefTopK(s, 10)
	fmt.Println(res)
	// [9 8 4 7 6 0 2 5 3 1]
}
