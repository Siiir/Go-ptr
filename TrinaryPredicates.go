package ptr

import (
	"C"
)

// Returns logical value of `val`∈⟨`leftInclusiveBound`,`rightInclusiveBound`⟩
func InRange[T any](val, leftInclusiveBound, rightInclusiveBound *T) bool {
	v := newUintptr(val)
	return (newUintptr(leftInclusiveBound) <= v && v <= newUintptr(rightInclusiveBound))
}
