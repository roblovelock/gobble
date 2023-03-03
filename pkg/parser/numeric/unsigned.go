package numeric

import (
	"encoding/binary"
	"gobble/pkg/combinator"
	"gobble/pkg/parser"
	"gobble/pkg/parser/bytes"
)

func Uint8() parser.Parser[parser.Reader, uint8] {
	return func(in parser.Reader) (uint8, error) {
		return in.ReadByte()
	}
}

func BEUint16() parser.Parser[parser.Reader, uint16] {
	return endianUint(2, binary.BigEndian.Uint16)
}

func BEUint32() parser.Parser[parser.Reader, uint32] {
	return endianUint(4, binary.BigEndian.Uint32)
}

func BEUint64() parser.Parser[parser.Reader, uint64] {
	return endianUint(8, binary.BigEndian.Uint64)
}

func LEUint16() parser.Parser[parser.Reader, uint16] {
	return endianUint(2, binary.LittleEndian.Uint16)
}

func LEUint32() parser.Parser[parser.Reader, uint32] {
	return endianUint(4, binary.LittleEndian.Uint32)
}

func LEUint64() parser.Parser[parser.Reader, uint64] {
	return endianUint(8, binary.LittleEndian.Uint64)
}

func endianUint[T uint16 | uint32 | uint64](l uint, f func([]byte) T) parser.Parser[parser.Reader, T] {
	return combinator.Map[parser.Reader, []byte, T](
		bytes.Take(l),
		func(bytes []byte) (T, error) {
			return f(bytes), nil
		},
	)
}
