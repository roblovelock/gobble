package ascii

import (
	"errors"
	"github.com/roblovelock/gobble/pkg/combinator/branch"
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/combinator/sequence"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
)

var (
	// ErrOverflow the parsed digits don't fit in the value type
	ErrOverflow = errors.New("overflow") // the parsed digits don't fit in the value type
)

type (
	signedIntConstraint interface {
		int8 | int16 | int32 | int64
	}
	unsignedIntConstraint interface {
		uint8 | uint16 | uint32 | uint64
	}
	intConstraint interface {
		signedIntConstraint | unsignedIntConstraint
	}
)

// UInt8 will parse a number in text form to uint8
func UInt8() parser.Parser[parser.Reader, uint8] {
	return unsignedIntParser[uint8]()
}

// UInt16 will parse a number in text form to uint16
func UInt16() parser.Parser[parser.Reader, uint16] {
	return unsignedIntParser[uint16]()
}

// UInt32 will parse a number in text form to uint32
func UInt32() parser.Parser[parser.Reader, uint32] {
	return unsignedIntParser[uint32]()
}

// UInt64 will parse a number in text form to uint64
func UInt64() parser.Parser[parser.Reader, uint64] {
	return unsignedIntParser[uint64]()
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

func signedIntParser[T signedIntConstraint]() parser.Parser[parser.Reader, T] {
	return branch.Alt(
		sequence.Preceded(bytes.Byte('-'), intParser(checkedSub[T])),
		unsignedIntParser[T](),
	)
}

func unsignedIntParser[T intConstraint]() parser.Parser[parser.Reader, T] {
	return sequence.Preceded(modifier.Optional(bytes.Byte('+')), intParser(checkedAdd[T]))
}

func intParser[T intConstraint](
	f func(a T, b uint8) (T, error),
) parser.Parser[parser.Reader, T] {
	return modifier.Map(Digit1(), func(digits []byte) (result T, err error) {
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

func checkedMul[T intConstraint](a T, b T) (T, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}

	result := a * b
	if a == result/b {
		return result, nil
	}

	return 0, ErrOverflow
}

func checkedAdd[T intConstraint](a T, b uint8) (T, error) {
	result := a + T(b)
	if (result > a) == (b > 0) {
		return result, nil
	}

	return 0, ErrOverflow
}

func checkedSub[T intConstraint](a T, b uint8) (T, error) {
	result := a - T(b)
	if (result < a) == (b > 0) {
		return result, nil
	}

	return 0, ErrOverflow
}
