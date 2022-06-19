package unsafe

import (
	"testing"
	"unsafe"
)

func TestCmpNBytes(t *testing.T) {
	const tcQuantity = 7
	argTab := [tcQuantity]struct {
		a [2][]byte
		n uintptr
	}{
		{[2][]byte{[]byte("a\bcde\f"), []byte("abcde")}, 4},
		{[2][]byte{[]byte("abc"), []byte("def")}, 3},
		{[2][]byte{[]byte("XHKr"), []byte("XHKr")}, 1},

		{[2][]byte{[]byte("ajklsdfasdfkljklasjkasdfflkasdfaklfiwe492104|X/epDDaLFDJAKLDAIWALwd"), []byte("ajklsdfasdfkljklasjkasdfflkasdfaklfiwe492104|X/epDDaLFDJAKLDAIWALwd")}, 37},
		{
			[2][]byte{
				[]byte("C¨\x9ap=\x80\x03\x0c\x0f\x83?\x12Ü\xd5&\x99w\\\xb9\xf8m`\xc6\xe8\x91f¤\x00}\xd0)t\x9e\x05q\xb6}ýßWx&\x11\x8c\xd9"),
				[]byte("C¨\x9ap=\x80\x03\x0c\x0f\x83?\x12Ü\xd5&\x99w\\\xb9\xf8m`\xc6\xe8\x91f¤\x00}\xd0)t\x9e\x05q\xb6}ýßWx&\x11\x8c\xd9"),
			},
			42,
		},
		{[2][]byte{[]byte("\x1fÝnC\xefÎ5_"), []byte("\x1fÝnD\xefÎ5_")}, 4},

		{[2][]byte{[]byte(""), []byte("")}, 0},
	}
	expectTab := [tcQuantity]bool{
		false, false, true,

		true, true, false,

		true,
	}

	for tcInd, args := range argTab {
		n, p1, p2 := args.n, unsafe.Pointer(&args.a[0][0]), unsafe.Pointer(&args.a[1][0])
	}
}
