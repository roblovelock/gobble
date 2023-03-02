package bytes

import (
	"gobble/internal/parser"
	"gobble/internal/predicate"
	"io"
)

func Digit1() parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, io.EOF
		}

		if predicate.IsDigit(b) {
			return b, nil
		}

		return 0, parser.ErrNotMatched
	}
}
