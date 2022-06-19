package ptr

import (
	"testing"
	"unsafe"
)

func TestInRange(t *testing.T) {
	// Definitions
	const tQuantity = 5
	ArgsTab := [tQuantity][3]uintptr{
		{90, 20, 1000},
		{200000, 200000, 200000},
		{200000, 10000000, 0},
		{200001, 200000, 200000},
		{200001, 200001, 200000},
	}
	ExpectTab := [tQuantity]bool{true, true, false, false, false}

	// Algorithm
	for i := 0; i < tQuantity; i++ {
		// Converted args go to `a` array
		var a [3]*float64
		for j, uPtr := range ArgsTab[i] {
			a[j] = (*float64)(unsafe.Pointer(uPtr))
		}

		//Now we can run our function
		if ex, got := ExpectTab[i], InRange(a[0], a[1], a[2]); ex != got {
			t.Errorf("\nTest case %d failed. Expected != Got"+
				"\n%v != %v", i, ex, got)
		}
	}
}
