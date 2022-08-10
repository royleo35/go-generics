package internal

import (
	"testing"
	"unsafe"
)

func assert(cond bool) {
	if !cond {
		panic("ops")
	}
}

func TestMemCmpValue(t *testing.T) {
	v1 := any(uint64(1))
	v2 := any(int64(1))
	res := MemCmpValue(v1, v2)
	assert(res != 0)

	// 0 bytes
	v := struct{}{}
	res = MemCmpValue(v, v)
	assert(res == 0)
	// 1bytes
	res = MemCmpValue(byte(1), byte(2))
	assert(res < 0)
	// 2bytes
	res = MemCmpValue(uint16(1), uint16(2))
	assert(res < 0)
	// 3bytes
	res = MemCmpValue([3]byte{1, 2, 3}, [3]byte{1, 2, 1})
	assert(res > 0)
	// 4bits
	res = MemCmpValue(int32(10), int32(9))
	assert(res > 0)

	// 8bits
	res = MemCmpValue(int64(10), int64(9))
	assert(res > 0)

	// 8bits chan ptr
	var ch chan int
	var ch2 chan int
	assert(MemCmpValue(ch, ch2) == 0)

	// 8bits map ptr
	var m1 map[int]int
	var m2 map[int]int
	assert(SizeOf[map[int]int]() == 8)
	assert(MemCmpValue(m1, m2) == 0)

	// 9bits
	res = MemCmpValue([9]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}, [9]byte{1, 2, 3, 4, 5, 6, 7, 8, 3})
	assert(res > 0)
	assert(MemCmpValue([9]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}, [9]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}) == 0)

	//16bbit any
	var inter1 any
	var inter2 any
	assert(SizeOf[any]() == 16)
	assert(MemCmpValue(inter2, inter1) == 0)

}

func TestMemCopy(t *testing.T) {
	a := [9]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := [9]byte{}
	MemCopy(unsafe.Pointer(&b), unsafe.Pointer(&a), SizeOfObj(a))
	assert(a == b)
}

func BenchmarkSizeOf1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SizeOf[[1e6]byte]()
	}
}

func BenchmarkSizeOf2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SizeOf[[1]byte]()
	}
}

func BenchmarkSizeOfObj1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SizeOfObj([1e6]byte{})
	}
}

func BenchmarkSizeOfObj2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SizeOfObj([1]byte{})
	}
}
