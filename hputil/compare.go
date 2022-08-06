package hputil

import (
	"unsafe"
)

type tflag uint8
type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype
type rtype struct {
	size       uintptr
	ptrdata    uintptr // number of bytes in the type that can contain pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      tflag   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldAlign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte   // garbage collection data
	str       nameOff // string form
	ptrToThis typeOff // type for pointer to this type, may be zero
}

type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}

func unpackEface(i interface{}) *emptyInterface {
	e := (*emptyInterface)(unsafe.Pointer(&i))
	return e
}

func DeepEqual(obj1, obj2 interface{}) bool {
	if obj1 == nil || obj2 == nil {
		return obj1 == obj2
	}
	//if reflect.TypeOf(obj1) != reflect.TypeOf(obj2) {
	//	return false
	//}
	fn := unpackEface(&obj1).typ.equal
	return fn(unsafe.Pointer(&obj1), unsafe.Pointer(&obj2))
}
