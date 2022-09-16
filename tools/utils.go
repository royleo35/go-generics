package tools

import (
	"github.com/royleo35/go-generics/types"
	"unsafe"
)

func PtrOf[T any](v T) *T {
	return &v
}

func ForEach[T any](data []T, do func(e *T)) {
	for i := range data {
		do(&data[i])
	}
}

func IfThen[T any](cond bool, trueVal T, falseVal T) T {
	if cond {
		return trueVal
	}
	return falseVal
}

func Dereference[T any](p *T) (val T) {
	if p == nil {
		return
	}
	return *p
}

// PtrCast 指针强转，不安全，转换失败会panic
func PtrCast[T1 any, T2 any](p *T1) *T2 {
	return (*T2)(unsafe.Pointer(p))
}

func Less[T types.BuiltIn](e1, e2 T) bool {
	return e1 < e2
}

func Greater[T types.BuiltIn](e1, e2 T) bool {
	return e1 > e2
}
