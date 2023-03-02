package bits

import (
	"gobble/internal/parser"
	"io"
	"math/bits"
)

func Tag[T uint8 | uint16 | uint32 | uint64 | uint](n uint8, tag T) parser.Parser[parser.BitReader, T] {
	var t T
	l := bits.Len64(uint64(t) - 1)

	if int(n) > l {
		return func(in parser.BitReader) (T, error) {
			return t, ErrBitsOverflow
		}
	}

	return func(in parser.BitReader) (T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		b, _, err := in.ReadBits(n)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return t, err
		}
		if T(b) != tag {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return t, parser.ErrNotMatched
		}

		return tag, nil

	}
}
