package ptr

import (
	"testing"
	"unsafe"

	"github.com/Siiir/asserter/v2"
	usptr "github.com/Siiir/unsafe/ptr"
)

func Test_CmpBitsAs_And_AllAliases(t *testing.T) {
	// Construction of asserter
	a, e := asserter.NewReseted(2, func(s string) { t.Error(s) })
	if e != nil {
		panic(e)
	}

	{
		a.Inc(0) // 1.*
		// Equality of all byte chunks in the same integer.

		var p int32 = 5
		a.A(CmpBitsAs[uint8](&p, &p))
		a.A(CmpBitsAs[int16](&p, &p))
		a.A(CmpBitsAs[int32](&p, &p))
	}
	{
		a.Inc(0) // 2.*
		// Comparisons of byte chunks in 2 similar but different long intigers.
		var p, q uint64 = 0b10111101_11011011_00000000_00000000_00000000_00000000_11011011_10111101, 0b10111101_11011011_11111111_11111111_11111111_11111111_11011011_10111101

		// Equality of the first two bytes in different integers.
		a.IncLast() // ≥2.1
		a.Assert(CmpBitsAs[int8](&p, &q))
		a.Assert(CmpBitsAs[uint16](&p, &q))

		// Inequality of integers themselves.
		// ≥2.3
		a.Assert(!CmpBitsAs[uint32](&p, &q))
		a.Assert(!CmpBitsAs[int64](&p, &q))

		// A proof that cmp64Bits works as expected
		// ≥2.5
		a.Assert(Cmp64Bits(&p, &p))
		a.Assert(!Cmp64Bits(&p, &q))
	}
	{
		a.Inc(0) // 3.*
		// Comparisons of arrays.

		var (
			ba1 = [20]byte{'a', 'b', 'c', 'd', '.', 'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l'}
			ba2 = [17]byte{'Z', 'Y', 'X', 'W', 'a', 'b', 'c', 'd', 'a', 'b', 'c', 'd', 'a', 'b', 'c', 'd', '-'}
		)

		// ≥3.0
		a.Assert(!Cmp128Bits(&ba1, &ba2))
		a.Assert(!Cmp32Bits(&ba1, &ba2))
		{ // ≥3.2
			p1 := (*[5]byte)(unsafe.Pointer(&ba1))
			p2 := (*[4]byte)(unsafe.Pointer(&ba2))
			for _, offset := range [3]uintptr{4, 8, 12} {
				a.Assert(Cmp32Bits(
					p1, usptr.Offset(p2, offset),
				))
			}
		}
		// ≥3.5 − Same array.
		// Same part.
		a.Assert(CmpBitsAs[int32](
			(*[7]uint8)(unsafe.Add(unsafe.Pointer(&ba1), 4)),
			(*[5]uint8)(unsafe.Add(unsafe.Pointer(&ba1), 4)),
		))
		a.Assert(CmpBitsAs[int16](
			(*[2]uint8)(unsafe.Add(unsafe.Pointer(&ba2), 4)),
			(*[2]uint8)(unsafe.Add(unsafe.Pointer(&ba2), 12)),
		))
	}
	{
		a.Inc(0) // 4.*
		// Comparisons with use of arrays types as `CT`.
		a.Assert(CmpBitsAs[[64]uint16](
			&[31]uint64{1, 2}, &[52]uint64{1, 2},
		))
		a.Assert(CmpBitsAs[[64]uint16](
			&[31][2]uint64{{1, 2}}, &[52][7]uint64{{1, 2}},
		))
		a.Assert(!CmpBitsAs[[64]uint8](
			&[16]uint32{1, 2}, &[17]uint32{1, 2, 3},
		))
	}
	{ // Delibrate linkage-time errors
		/*
			var p uint32= 6342
			CmpBitsAs[uint64](&p, &p)

			a.Assert(CmpBitsAs[[64]uint8](
				&[8]uint32{1, 2}, &[17]uint32{1, 2},
			))
		*/
	}
}
