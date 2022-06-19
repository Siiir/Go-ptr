package ptr

import (
	"testing"

	"github.com/Siiir/asserter"
)

func TestCmpBitsAs(t *testing.T) {
	// Construction of asserter
	a, e := asserter.New(make([]uint, 2), func(s string) { t.Error(s) })
	if e != nil {
		panic("e!=nil")
	}

	{
		a.Inc(0)
		// Equality of all bit chunks in the same integer.
		var p int32 = 5
		a.A(CmpBitsAs[uint8](&p, &p))
		a.A(CmpBitsAs[int16](&p, &p))
		a.A(CmpBitsAs[int32](&p, &p))
		// Instruction ```CmpBitsAs[uint64](&p, &p)``` gives appropiate linkage-time error
	}
	{
		a.Inc(0)
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
