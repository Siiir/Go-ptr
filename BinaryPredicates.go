package ptr

import "unsafe"

func delibrateLinkageTimeError__InT_HasSmallerSizeThan_CT()
func CmpBitsAs[CT comparable, InT any](p1, p2 *InT) bool {
	if unsafe.Sizeof(*(*InT)(nil)) < unsafe.Sizeof(*(*CT)(nil)) {
		delibrateLinkageTimeError__InT_HasSmallerSizeThan_CT()
	}
	return *(*CT)(unsafe.Pointer(p1)) == *(*CT)(unsafe.Pointer(p2))
}
func Cmp32Bits[T any](p1, p2 *T) bool {
	return CmpBitsAs[uint64](p1, p2)
}
func Cmp64Bits[T any](p1, p2 *T) bool {
	return CmpBitsAs[uint64](p1, p2)
}
