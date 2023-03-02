package numeric

import (
	"encoding/binary"
	"gobble/internal/combinator"
	"gobble/internal/parser"
	"gobble/internal/parser/bytes"
)

func Uint8[R parser.Reader]() parser.Parser[R, uint8] {
	return func(in R) (uint8, error) {
		return in.ReadByte()
	}
}

func BEUint16[R parser.Reader]() parser.Parser[R, uint16] {
	return endianUint[R](2, binary.BigEndian.Uint16)
}

func BEUint32[R parser.Reader]() parser.Parser[R, uint32] {
	return endianUint[R](4, binary.BigEndian.Uint32)
}

func BEUint64[R parser.Reader]() parser.Parser[R, uint64] {
	return endianUint[R](8, binary.BigEndian.Uint64)
}

func LEUint16[R parser.Reader]() parser.Parser[R, uint16] {
	return endianUint[R](2, binary.LittleEndian.Uint16)
}

func LEUint32[R parser.Reader]() parser.Parser[R, uint32] {
	return endianUint[R](4, binary.LittleEndian.Uint32)
}

func LEUint64[R parser.Reader]() parser.Parser[R, uint64] {
	return endianUint[R](8, binary.LittleEndian.Uint64)
}

func endianUint[R parser.Reader, T uint16 | uint32 | uint64](l uint, f func([]byte) T) parser.Parser[R, T] {
	return combinator.Map[R, []byte, T](
		bytes.Take[R](l),
		func(bytes []byte) (T, error) {
			return f(bytes), nil
		},
	)
}
