package myslice

import (
	"fmt"
	"unsafe"
)

type Slice[T any] struct {
	data     unsafe.Pointer
	len      int
	cap      int
	elemSize int // must in last position
}

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func sizeof[T any]() int {
	var dummy T
	return int(unsafe.Sizeof(dummy))
}

func memcopy(dst, src unsafe.Pointer, bytes int) {
	// assert(bytes >= 0, "bytes must >= 0")
	l8 := bytes >> 3
	for i := 0; i < l8; i++ {
		*(*uint64)(unsafe.Pointer(uintptr(dst) + uintptr(i*8))) = *(*uint64)(unsafe.Pointer(uintptr(src) + uintptr(i*8)))
	}
	for i := 0; i < bytes&7; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(dst) + uintptr(i))) = *(*uint8)(unsafe.Pointer(uintptr(src) + uintptr(i)))
	}
}

func nextPower(n int) int {
	num := uint64(n)
	res := 0
	for half, cnt := 32, 0; num != 0 && cnt <= 5; cnt++ {
		bits := half >> cnt
		high, low := num>>(bits), num&((1<<bits)-1)
		if high != 0 {
			res += bits
			num = high
		} else {
			num = low
		}
	}
	return 1 << (res + 1)
}

func nextCap(n int) int {
	if n < 1024 {
		return n << 1
	}
	return int(float64(n) * 1.25)
}

func malloc(bytes int) unsafe.Pointer {
	data := make([]byte, bytes)
	return unsafe.Pointer(&data[0])
}

func New[T any](_len, _cap int) *Slice[T] {
	assert(_len >= 0 && _cap >= 0 && _cap >= _len, "len or cap param error")
	s := &Slice[T]{
		data:     nil,
		len:      _len,
		cap:      _cap,
		elemSize: sizeof[T](),
	}
	if _cap == 0 {
		return s
	}
	s.data = malloc(_cap * s.elemSize)
	return s
}

func NewWithValues[T any](e ...T) *Slice[T] {
	s := New[T](len(e), len(e))
	if len(e) == 0 {
		return s
	}
	memcopy(s.ptrOf(0), unsafe.Pointer(&e[0]), len(e)*s.elemSize)
	return s
}

func (s *Slice[T]) ptrOf(i int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(s.data) + uintptr(s.elemSize*i))
}

func (s *Slice[T]) At(i int) T {
	if i < 0 || i >= s.len {
		panic(fmt.Sprintf("index out of slice[%d] at index %d", s.len, i))
	}
	// assert(i >= 0 && i < s.len, fmt.Sprintf("index out of slice[%d] at index %d", s.len, i))
	return *(*T)(s.ptrOf(i))
}

func (s *Slice[T]) Set(i int, val T) *Slice[T] {
	if i < 0 || i >= s.len {
		panic(fmt.Sprintf("index out of slice[%d] at index %d", s.len, i))
	}
	// assert(i >= 0 && i < s.len, fmt.Sprintf("index out of slice[%d] at index %d", s.len, i))
	*(*T)(s.ptrOf(i)) = val
	return s
}

func (s *Slice[T]) shadow(dst *Slice[T]) {
	s.data = dst.data
	s.len = dst.len
	s.cap = dst.cap
}

func (s *Slice[T]) Append(e ...T) *Slice[T] {
	l := len(e)
	if l == 0 {
		return s
	}
	needGrowUp := s.len+l > s.cap
	if !needGrowUp {
		memcopy(s.ptrOf(s.len), unsafe.Pointer(&e[0]), l*s.elemSize)
		s.len += l
		return s
	}
	newS := New[T](0, nextCap(s.len+l))
	// copy old
	if s.len != 0 {
		memcopy(newS.ptrOf(0), s.ptrOf(0), s.len*s.elemSize)
	}
	// copy new
	memcopy(newS.ptrOf(s.len), unsafe.Pointer(&e[0]), l*s.elemSize)
	newS.len = s.len + l
	s.shadow(newS)
	return newS
}

func (s *Slice[T]) Cut(left, right int) *Slice[T] {
	return s.Slice(left, right)
}

func (s *Slice[T]) Slice(left, right int) *Slice[T] {
	assert(left < right, "slice index left must <= right")
	assert(left >= 0 && right <= s.len,
		fmt.Sprintf("slice index must be in [0, %d), but given [%d, %d)", s.len, left, right))
	res := &Slice[T]{
		data:     s.ptrOf(left),
		len:      right - left,
		cap:      s.cap - left,
		elemSize: s.elemSize,
	}
	return res
}

func (s *Slice[T]) Len() int {
	return s.len
}

func (s *Slice[T]) Cap() int {
	return s.cap
}

func (s *Slice[T]) Copy() *Slice[T] {
	res := *s
	return &res
}

func (s *Slice[T]) ToSlice() []T {
	return *(*[]T)(unsafe.Pointer(s))
}

// for debug
func (s *Slice[T]) bytes() []byte {
	res := &Slice[T]{
		data: s.data,
		len:  s.len * s.elemSize,
		cap:  s.cap * s.elemSize,
	}
	return *(*[]byte)(unsafe.Pointer(res))
}

func bytes[T any](s []T) []byte {
	ss := *(*Slice[T])(unsafe.Pointer(&s))
	elemSize := sizeof[T]()
	res := &Slice[T]{
		data: ss.data,
		len:  ss.len * elemSize,
		cap:  ss.cap * elemSize,
	}
	return *(*[]byte)(unsafe.Pointer(res))
}
