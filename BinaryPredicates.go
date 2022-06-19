package ptr

import (
	"C"
	"unsafe"
)

/*
	Phantom function to be placed in compile-time-dead code block.
	If compiler does not optimeze it off, the linkage error, containing function name, will appear.
*/
func delibrateLinkageTimeError__InT_HasSmallerSizeThan_CT()

/*
	Compares values pointed by `p1` & `p2`,
	as if they were values of type CT, using equality operator.

	If `InT` is smaller in memory than `CT`,
	function will deliberately fail to compile during linkage.
*/
func CmpBitsAs[CT comparable, InT any](p1, p2 *InT) bool {
	if unsafe.Sizeof(*(*InT)(nil)) < unsafe.Sizeof(*(*CT)(nil)) {
		delibrateLinkageTimeError__InT_HasSmallerSizeThan_CT()
	}
	return *(*CT)(unsafe.Pointer(p1)) == *(*CT)(unsafe.Pointer(p2))
}

// Equivalent to:  CmpBitsAs[uint8](p1, p2)
func Cmp8Bits[T any](p1, p2 *T) bool {
	return CmpBitsAs[uint8](p1, p2)
}

// Equivalent to:  CmpBitsAs[uint16](p1, p2)
func Cmp16Bits[T any](p1, p2 *T) bool {
	return CmpBitsAs[uint16](p1, p2)
}

// Equivalent to:  CmpBitsAs[uint32](p1, p2)
func Cmp32Bits[T any](p1, p2 *T) bool {
	return CmpBitsAs[uint32](p1, p2)
}

// Equivalent to:  CmpBitsAs[uint64](p1, p2)
func Cmp64Bits[T any](p1, p2 *T) bool {
	return CmpBitsAs[uint64](p1, p2)
}

// Equivalent to:  CmpBitsAs[[2]uint64](p1, p2)
func Cmp128Bits[T any](p1, p2 *T) bool {
	return CmpBitsAs[[2]uint64](p1, p2)
}
