package character

import (
	"errors"
	"gobble/internal/parser"
	"gobble/internal/parser/bytes"
)

var (
	// UInt8 will parse a number in text form to a number
	UInt8 = intParser[uint8](checkedAdd[uint8])
	// UInt16 will parse a number in text form to a number
	UInt16 = intParser[uint16](checkedAdd[uint16])
	// UInt32 will parse a number in text form to a number
	UInt32 = intParser[uint32](checkedAdd[uint32])
	// UInt64 will parse a number in text form to a number
	UInt64 = intParser[uint64](checkedAdd[uint64])
	// Int8 will parse a number in text form to a number
	Int8 = signedIntParser[int8]()
	// Int16 will parse a number in text form to a number
	Int16 = signedIntParser[int16]()
	// Int32 will parse a number in text form to a number
	Int32 = signedIntParser[int32]()
	// Int64 will parse a number in text form to a number
	Int64 = signedIntParser[int64]()
)

func signedIntParser[T int8 | int16 | int32 | int64]() parser.Parser[parser.Reader, T] {
	return func(in parser.Reader) (T, error) {
		s, err := bytes.OneOf('-', '+')(in)
		if err == nil && s == '-' {
			return intParser(checkedSub[T])(in)
		}
		return intParser(checkedAdd[T])(in)
	}
}

func intParser[T uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64](
	f func(a T, b uint8) (T, error),
) parser.Parser[parser.Reader, T] {
	return func(in parser.Reader) (T, error) {
		var result T
		b, err := bytes.Digit1()(in)
		if err != nil {
			return 0, err
		}

		result += T(b - '0')
		for b, err := bytes.Digit1()(in); err == nil; b, err = bytes.Digit1()(in) {
			v, err := checkedMul(result, 10)
			if err != nil {
				return result, nil
			}
			v, err = f(v, b)
			if err != nil {
				return result, nil
			}
			result = v
		}

		return result, nil
	}
}

func checkedMul[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a T, b T) (T, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}

	result := a * b
	if a == result/b {
		return result, nil
	}

	return 0, errors.New("overflow")
}

func checkedAdd[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a T, b uint8) (T, error) {
	result := a + T(b)
	if (result > a) == (b > 0) {
		return result, nil
	}

	return 0, errors.New("overflow")
}

func checkedSub[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a T, b uint8) (T, error) {
	result := a - T(b)
	if (result < a) == (b > 0) {
		return result, nil
	}

	return 0, errors.New("overflow")
}
