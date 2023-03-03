package ascii

import (
	"errors"
	"gobble/pkg/combinator"
	"gobble/pkg/combinator/branch"
	"gobble/pkg/combinator/sequence"
	"gobble/pkg/parser"
	"gobble/pkg/parser/bytes"
)

var (
	// ErrOverflow the parsed digits don't fit in the value type
	ErrOverflow = errors.New("overflow") // the parsed digits don't fit in the value type
)

// UInt8 will parse a number in text form to uint8
func UInt8() parser.Parser[parser.Reader, uint8] {
	return positiveIntParser[uint8]()
}

// UInt16 will parse a number in text form to uint16
func UInt16() parser.Parser[parser.Reader, uint16] {
	return positiveIntParser[uint16]()
}

// UInt32 will parse a number in text form to uint32
func UInt32() parser.Parser[parser.Reader, uint32] {
	return positiveIntParser[uint32]()
}

// UInt64 will parse a number in text form to uint64
func UInt64() parser.Parser[parser.Reader, uint64] {
	return positiveIntParser[uint64]()
}

// Int8 will parse a number in text form to uint8
func Int8() parser.Parser[parser.Reader, int8] {
	return signedIntParser[int8]()
}

// Int16 will parse a number in text form to int16
func Int16() parser.Parser[parser.Reader, int16] {
	return signedIntParser[int16]()
}

// Int32 will parse a number in text form to uint32
func Int32() parser.Parser[parser.Reader, int32] {
	return signedIntParser[int32]()
}

// Int64 will parse a number in text form to int64
func Int64() parser.Parser[parser.Reader, int64] {
	return signedIntParser[int64]()
}

func signedIntParser[T int8 | int16 | int32 | int64]() parser.Parser[parser.Reader, T] {
	return branch.Alt(
		sequence.Preceded(bytes.Byte('-'), intParser(checkedSub[T])),
		positiveIntParser[T](),
	)
}

func positiveIntParser[T uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64]() parser.Parser[parser.Reader, T] {
	return sequence.Preceded(combinator.Optional(bytes.Byte('+')), intParser(checkedAdd[T]))
}

func intParser[T uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64](
	f func(a T, b uint8) (T, error),
) parser.Parser[parser.Reader, T] {
	return combinator.Map(Digit1(), func(digits []byte) (result T, err error) {
		for _, d := range digits {
			result, err = checkedMul(result, 10)
			if err != nil {
				var r T
				return r, err
			}
			result, err = f(result, d-'0')
			if err != nil {
				var r T
				return r, err
			}
		}
		return result, nil
	})
}

func checkedMul[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a T, b T) (T, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}

	result := a * b
	if a == result/b {
		return result, nil
	}

	return 0, ErrOverflow
}

func checkedAdd[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a T, b uint8) (T, error) {
	result := a + T(b)
	if (result > a) == (b > 0) {
		return result, nil
	}

	return 0, ErrOverflow
}

func checkedSub[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a T, b uint8) (T, error) {
	result := a - T(b)
	if (result < a) == (b > 0) {
		return result, nil
	}

	return 0, ErrOverflow
}
