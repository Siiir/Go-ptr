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

func Test_cmpBitsAs(t *testing.T) {
	a := newAsserter(make([]uint, 2), func(s string) { t.Error(s) })
	{
		a.Inc(1)
		// Equality of bit chunks in the same integer.
		var p int32 = 5
		a.A(CmpBitsAs[uint8](&p, &p))
		a.A(CmpBitsAs[int16](&p, &p))
		a.A(CmpBitsAs[int32](&p, &p))
		// Instruction ```CmpBitsAs[uint64](&p, &p)``` gives appropiate compile-time error
	}
	{
		a.Inc(1)
		// Equality of the first two bytes in different integers.

		var p, q uint64 = 0b10111101_11011011_00000000_00000000_00000000_00000000_11011011_10111101, 0b10111101_11011011_11111111_11111111_11111111_11111111_11011011_10111101
		a.Assert(CmpBitsAs[int8](&p, &q))
		a.Assert(CmpBitsAs[uint16](&p, &q))
		a.Assert(!CmpBitsAs[uint32](&p, &q))
		a.Assert(!CmpBitsAs[int64](&p, &q))
		// A proof that cmp64Bits works as expected
		a.Assert(Cmp64Bits(&p, &p))
		a.Assert(!Cmp64Bits(&p, &q))
		// Lets fail assertion
		a.Assert(false)
	}
}
