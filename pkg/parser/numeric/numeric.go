// Package numeric provides parsers for recognizing numeric bytes
package numeric

import (
	"encoding/binary"
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
)

// UInt8 returns a 1 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func UInt8() parser.Parser[parser.Reader, uint8] {
	return bytes.One()
}

// Int8 returns a 1 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int8() parser.Parser[parser.Reader, int8] {
	return func(in parser.Reader) (int8, error) {
		b, err := in.ReadByte()
		return int8(b), err
	}
}

// UInt16BE returns a big endian 2 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func UInt16BE() parser.Parser[parser.Reader, uint16] {
	return endian(2, binary.BigEndian.Uint16)
}

// Int16BE returns a big endian 2 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int16BE() parser.Parser[parser.Reader, int16] {
	return endian(2, cast[uint16, int16](binary.BigEndian.Uint16))
}

// UInt32BE returns a big endian 4 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func UInt32BE() parser.Parser[parser.Reader, uint32] {
	return endian(4, binary.BigEndian.Uint32)
}

// Int32BE returns a big endian 4 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int32BE() parser.Parser[parser.Reader, int32] {
	return endian(4, cast[uint32, int32](binary.BigEndian.Uint32))
}

// UInt64BE returns a big endian 8 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func UInt64BE() parser.Parser[parser.Reader, uint64] {
	return endian(8, binary.BigEndian.Uint64)
}

// Int64BE returns a big endian 8 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int64BE() parser.Parser[parser.Reader, int64] {
	return endian(8, cast[uint64, int64](binary.BigEndian.Uint64))
}

// UInt16LE returns a little endian 2 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func UInt16LE() parser.Parser[parser.Reader, uint16] {
	return endian(2, binary.LittleEndian.Uint16)
}

// Int16LE returns a little endian 2 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int16LE() parser.Parser[parser.Reader, int16] {
	return endian(2, cast[uint16, int16](binary.LittleEndian.Uint16))
}

// UInt32LE returns a little endian 4 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func UInt32LE() parser.Parser[parser.Reader, uint32] {
	return endian(4, binary.LittleEndian.Uint32)
}

// Int32LE returns a little endian 4 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int32LE() parser.Parser[parser.Reader, int32] {
	return endian(4, cast[uint32, int32](binary.LittleEndian.Uint32))
}

// UInt64LE returns a little endian 8 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func UInt64LE() parser.Parser[parser.Reader, uint64] {
	return endian(8, binary.LittleEndian.Uint64)
}

// Int64LE returns a little endian 8 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int64LE() parser.Parser[parser.Reader, int64] {
	return endian(8, cast[uint64, int64](binary.LittleEndian.Uint64))
}

func cast[U uint16 | uint32 | uint64, S int16 | int32 | int64](f func(b []byte) U) func([]byte) S {
	return func(b []byte) S {
		return S(f(b))
	}
}

func endian[T uint16 | uint32 | uint64 | int16 | int32 | int64](l uint, f func([]byte) T) parser.Parser[parser.Reader, T] {
	return modifier.Map[parser.Reader, []byte, T](
		bytes.Take(l),
		func(bytes []byte) (T, error) {
			return f(bytes), nil
		},
	)
}
