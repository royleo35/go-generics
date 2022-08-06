package types

type Number interface {
	~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64 | ~int | ~uint | ~float32 | ~float64
}

type BuiltIn interface {
	Number | ~string
}

type Hashable interface {
	~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64 | ~int | ~uint | ~string
}

type Integer interface {
	~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64 | ~int | ~uint
}

type Natural interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}
