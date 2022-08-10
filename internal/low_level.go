package internal

import "unsafe"

func SizeOfObj[T any](v T) int {
	return int(unsafe.Sizeof(v))
}
func SizeOf[T any]() int {
	var dummy T
	return SizeOfObj[T](dummy)
}

func MemCopy(dst, src unsafe.Pointer, bytes int) {
	// assert(bytes >= 0, "bytes must >= 0")
	l8 := bytes >> 3
	for i := 0; i < l8; i++ {
		*(*uint64)(unsafe.Pointer(uintptr(dst) + uintptr(i*8))) = *(*uint64)(unsafe.Pointer(uintptr(src) + uintptr(i*8)))
	}
	for i, off := 0, l8<<3; i < bytes&7; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(dst) + uintptr(off+i))) = *(*uint8)(unsafe.Pointer(uintptr(src) + uintptr(off+i)))
	}
}

func MemCmpValue[T any](v1, v2 T) int {
	return memcmp(unsafe.Pointer(&v1), unsafe.Pointer(&v2), SizeOf[T]())
}

// memcmp 比较两段内存大小，要求内存长度必须相同;先每次比较8字节，剩余部分逐字节比较,注意8字节比较时与计算机大小端有关系
func memcmp(data1, data2 unsafe.Pointer, bytes int) int {
	l8 := bytes >> 3
	for i := 0; i < l8; i++ {
		left := *(*uint64)(unsafe.Pointer(uintptr(data1) + uintptr(i*8)))
		right := *(*uint64)(unsafe.Pointer(uintptr(data2) + uintptr(i*8)))
		if left > right {
			return 1
		} else if left < right {
			return -1
		}
	}
	for i, off := 0, l8<<3; i < bytes&7; i++ {
		left := *(*byte)(unsafe.Pointer(uintptr(data1) + uintptr(off+i)))
		right := *(*byte)(unsafe.Pointer(uintptr(data2) + uintptr(off+i)))
		if left > right {
			return 1
		} else if left < right {
			return -1
		}
	}
	return 0
}

func Malloc(bytes int) unsafe.Pointer {
	data := make([]byte, bytes)
	return unsafe.Pointer(&data[0])
}
