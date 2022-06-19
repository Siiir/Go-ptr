package ptr

import "unsafe"

// Casts *T into uintptr
func NewUintptr[T any](val *T) uintptr {
	return uintptr(unsafe.Pointer(val))
}
