package ptr

import (
	"C"
	"unsafe"
)

/*
	Phantom function to be placed in compile-time-dead code block.
	If compiler does not optimeze it off, the linkage error, containing function name, will appear.
*/
func delibrateLinkageTimeError__T1_HasSmallerSizeThan_CT()

/*
	Phantom function to be placed in compile-time-dead code block.
	If compiler does not optimeze it off, the linkage error, containing function name, will appear.
*/
func delibrateLinkageTimeError__T2_HasSmallerSizeThan_CT()

/*
	Compares values pointed by `p1` & `p2`,
	as if they were values of type CT, using equality operator.

	If `InT` is smaller in memory than `CT`,
	function will deliberately fail to compile during linkage.
*/
func CmpBitsAs[CT comparable, T1, T2 any](p1 *T1, p2 *T2) bool {
	if unsafe.Sizeof(*(*T1)(nil)) < unsafe.Sizeof(*(*CT)(nil)) {
		delibrateLinkageTimeError__T1_HasSmallerSizeThan_CT()
	}
	if unsafe.Sizeof(*(*T2)(nil)) < unsafe.Sizeof(*(*CT)(nil)) {
		delibrateLinkageTimeError__T2_HasSmallerSizeThan_CT()
	}
	return *(*CT)(unsafe.Pointer(p1)) == *(*CT)(unsafe.Pointer(p2))
}

// Equivalent to:  CmpBitsAs[uint8](p1, p2)
func Cmp8Bits[T1, T2 any](p1 *T1, p2 *T2) bool {
	return CmpBitsAs[uint8](p1, p2)
}

// Equivalent to:  CmpBitsAs[uint16](p1, p2)
func Cmp16Bits[T1, T2 any](p1 *T1, p2 *T2) bool {
	return CmpBitsAs[uint16](p1, p2)
}

// Equivalent to:  CmpBitsAs[uint32](p1, p2)
func Cmp32Bits[T1, T2 any](p1 *T1, p2 *T2) bool {
	return CmpBitsAs[uint32](p1, p2)
}

// Equivalent to:  CmpBitsAs[uint64](p1, p2)
func Cmp64Bits[T1, T2 any](p1 *T1, p2 *T2) bool {
	return CmpBitsAs[uint64](p1, p2)
}

// Equivalent to:  CmpBitsAs[[2]uint64](p1, p2)
func Cmp128Bits[T1, T2 any](p1 *T1, p2 *T2) bool {
	return CmpBitsAs[[2]uint64](p1, p2)
}
