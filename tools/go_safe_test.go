package tools

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

type Cos struct {
	a int
	s string
	n *Cos
}

func TestGoSafe(t *testing.T) {
	i := 10
	f := func() Result[int] {
		return MakeResult(i*i, nil)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	var res Result[int]
	go GoSafe(context.Background(), &res, &wg, f)
	go GoSafe(context.Background(), nil, &wg, func() Result[any] {
		var a *int
		*a = 10
		return NilResult
	})
	wg.Wait()
	fmt.Println(res.HasValue() && res.Val() == 100)

}

func TestGoSafe1(t *testing.T) {

	f := func() Result[*Cos] {
		c := Cos{
			a: 1,
			s: "hah",
			n: &Cos{
				a: 2,
				s: "hah2",
				n: nil,
			},
		}
		c.a++
		c.n.s = "new"
		return MakeResult(&c, nil)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	var res Result[*Cos]
	go GoSafe(context.Background(), &res, &wg, f)
	wg.Wait()
	Assert(res.HasValue() && res.Val().a == 2)

}

func TestGoSafeNotReturn(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	var local = 1
	fmt.Println("local in test", local)
	go GoSafeNotReturn(context.Background(), func() {
		print("hello")
		wg.Done()
		local++
		fmt.Println("local in goroutine1", local)
	})
	go GoSafeNotReturn(context.Background(), func() {
		defer wg.Done()
		var a [1]int
		var idx = 1
		a[idx] = 10
		fmt.Println("local in goroutine2", local)
	})
	wg.Wait()
	fmt.Println("final local", local)
	print("done")
}
