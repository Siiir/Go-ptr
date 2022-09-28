package ptr

import (
	"C"
)

// Returns logical value of `val`∈⟨`leftInclusiveBound`,`rightInclusiveBound`⟩
func InRange[T any](val, leftInclusiveBound, rightInclusiveBound *T) bool {
	v := NewUintptr(val)
	return (NewUintptr(leftInclusiveBound) <= v && v <= NewUintptr(rightInclusiveBound))
}
