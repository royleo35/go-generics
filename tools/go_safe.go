package tools

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

var NilResult = MakeResult(any(nil), nil)

type Result[T any] struct {
	val T
	err error
}

func (r *Result[T]) HasValue() bool {
	return r.err == nil
}

func (r *Result[T]) Val() T {
	return r.val
}

func (r *Result[T]) Err() error {
	return r.err
}

func MakeResult[T any](val T, err error) Result[T] {
	return Result[T]{
		val: val,
		err: err,
	}
}

func GoSafe[T any](ctx context.Context, res *Result[T], wg *sync.WaitGroup, f func() Result[T]) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			s := fmt.Sprintf("got panic, err: %+v, stack: %v ", e, *(*string)(unsafe.Pointer(&buf)))
			if res != nil {
				res.err = errors.New(s)
			}
		}
		if wg != nil {
			wg.Done()
		}
	}()
	tmp := f()
	if res != nil {
		*res = tmp
	}
}

func GoSafeNotReturn(ctx context.Context, f func()) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			s := fmt.Sprintf("got panic, err: %+v, stack: %v ", e, *(*string)(unsafe.Pointer(&buf)))
			print(s)
		}
	}()
	f()
}
