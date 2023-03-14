// Package numeric provides parsers for recognizing numeric bytes
package numeric

import (
	"encoding/binary"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"math"
)

type (
	endianParserConstraint interface {
		bool | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
	}

	endianParser[T endianParserConstraint] struct {
		byteOrder binary.ByteOrder
	}
)

func (o *endianParser[T]) Parse(in parser.Reader) (result T, err error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	err = binary.Read(in, o.byteOrder, &result)
	if err == io.ErrUnexpectedEOF {
		_, _ = in.Seek(currentOffset, io.SeekStart)
		err = io.EOF
	}
	return result, err
}

func (o *endianParser[T]) ParseBytes(in []byte) (result T, out []byte, err error) {
	out, err = readNumeric(in, o.byteOrder, &result)
	return
}

func readNumeric(in []byte, order binary.ByteOrder, data any) ([]byte, error) {
	size := intDataSize(data)
	if len(in) < size {
		return in, io.EOF
	}

	switch data := data.(type) {
	case *bool:
		*data = in[0] != 0
	case *int8:
		*data = int8(in[0])
	case *uint8:
		*data = in[0]
	case *int16:
		*data = int16(order.Uint16(in[:size]))
	case *uint16:
		*data = order.Uint16(in[:size])
	case *int32:
		*data = int32(order.Uint32(in[:size]))
	case *uint32:
		*data = order.Uint32(in[:size])
	case *int64:
		*data = int64(order.Uint64(in[:size]))
	case *uint64:
		*data = order.Uint64(in[:size])
	case *float32:
		*data = math.Float32frombits(order.Uint32(in[:size]))
	case *float64:
		*data = math.Float64frombits(order.Uint64(in[:size]))
	}

	return in[size:], nil
}

func intDataSize(data any) int {
	switch data.(type) {
	case *bool, *int8, *uint8:
		return 1
	case *int16, *uint16:
		return 2
	case *int32, *uint32, *float32:
		return 4
	case *int64, *uint64, *float64:
		return 8
	}
	return 0
}

// UInt8 returns a 1 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func UInt8() parser.Parser[parser.Reader, uint8] {
	return &endianParser[uint8]{}
}

// Int8 returns a 1 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int8() parser.Parser[parser.Reader, int8] {
	return &endianParser[int8]{}
}

// Uint16BE returns a big endian 2 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func Uint16BE() parser.Parser[parser.Reader, uint16] {
	return &endianParser[uint16]{byteOrder: binary.BigEndian}
}

// Int16BE returns a big endian 2 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int16BE() parser.Parser[parser.Reader, int16] {
	return &endianParser[int16]{byteOrder: binary.BigEndian}
}

// Uint32BE returns a big endian 4 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func Uint32BE() parser.Parser[parser.Reader, uint32] {
	return &endianParser[uint32]{byteOrder: binary.BigEndian}
}

// Int32BE returns a big endian 4 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int32BE() parser.Parser[parser.Reader, int32] {
	return &endianParser[int32]{byteOrder: binary.BigEndian}
}

// Uint64BE returns a big endian 8 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func Uint64BE() parser.Parser[parser.Reader, uint64] {
	return &endianParser[uint64]{byteOrder: binary.BigEndian}
}

// Int64BE returns a big endian 8 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int64BE() parser.Parser[parser.Reader, int64] {
	return &endianParser[int64]{byteOrder: binary.BigEndian}
}

// Uint16LE returns a little endian 2 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func Uint16LE() parser.Parser[parser.Reader, uint16] {
	return &endianParser[uint16]{byteOrder: binary.LittleEndian}
}

// Int16LE returns a little endian 2 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int16LE() parser.Parser[parser.Reader, int16] {
	return &endianParser[int16]{byteOrder: binary.LittleEndian}
}

// Uint32LE returns a little endian 4 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func Uint32LE() parser.Parser[parser.Reader, uint32] {
	return &endianParser[uint32]{byteOrder: binary.LittleEndian}
}

// Int32LE returns a little endian 4 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int32LE() parser.Parser[parser.Reader, int32] {
	return &endianParser[int32]{byteOrder: binary.LittleEndian}
}

// UInt64LE returns a little endian 8 byte unsigned integer. io.EOF is returned if the input contains too few bytes
func UInt64LE() parser.Parser[parser.Reader, uint64] {
	return &endianParser[uint64]{byteOrder: binary.LittleEndian}
}

// Int64LE returns a little endian 8 byte signed integer. io.EOF is returned if the input contains too few bytes
func Int64LE() parser.Parser[parser.Reader, int64] {
	return &endianParser[int64]{byteOrder: binary.LittleEndian}
}

// Float32BE returns a big endian 4 byte floating point number. io.EOF is returned if the input contains too few bytes
func Float32BE() parser.Parser[parser.Reader, float32] {
	return &endianParser[float32]{byteOrder: binary.BigEndian}
}

// Float32LE returns a little endian 4 byte floating point number. io.EOF is returned if the input contains too few bytes
func Float32LE() parser.Parser[parser.Reader, float32] {
	return &endianParser[float32]{byteOrder: binary.LittleEndian}
}

// Float64BE returns a big endian 8 byte floating point number. io.EOF is returned if the input contains too few bytes
func Float64BE() parser.Parser[parser.Reader, float64] {
	return &endianParser[float64]{byteOrder: binary.BigEndian}
}

// Float64LE returns a little endian 8 byte floating point number. io.EOF is returned if the input contains too few bytes
func Float64LE() parser.Parser[parser.Reader, float64] {
	return &endianParser[float64]{byteOrder: binary.LittleEndian}
}

// Bool returns true if the next byte in the input isn't zero. io.EOF is returned if the input contains too few bytes
func Bool() parser.Parser[parser.Reader, bool] {
	return &endianParser[bool]{}
}
