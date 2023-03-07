// Package ascii provides parsers for recognizing ascii bytes
package ascii

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Digit matches a single ASCII numerical characters: 0-9
//   - If the input matches a numerical character, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func Digit() parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, err
		}

		if IsDigit(b) {
			return b, nil
		}

		_, _ = in.Seek(-1, io.SeekCurrent)
		return 0, parser.ErrNotMatched
	}
}

// Digit0 matches zero or more ASCII numerical character: 0-9
//   - If the input matches a numerical character, it will return a slice of all matched digits.
//   - If the input is empty, it will return an empty slice.
//   - If the input doesn't match a numerical character, it will return an empty slice.
func Digit0() parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		digits := make([]byte, 0)

		for {
			b, err := in.ReadByte()
			if err != nil {
				return digits, nil
			} else if !IsDigit(b) {
				_, _ = in.Seek(-1, io.SeekCurrent)
				return digits, nil
			}
			digits = append(digits, b)
		}
	}
}

// Digit1 matches one or more ASCII numerical characters: 0-9
//   - If the input matches a numerical character, it will return a slice of all matched digits.
//   - If the input is empty, it will return io.EOF.
//   - If the input doesn't match a numerical character, it will return parser.ErrNotMatched.
func Digit1() parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return nil, err
		}

		if !IsDigit(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return nil, parser.ErrNotMatched
		}

		digits := []byte{b}

		for {
			b, err := in.ReadByte()
			if err != nil {
				return digits, nil
			} else if !IsDigit(b) {
				_, _ = in.Seek(-1, io.SeekCurrent)
				return digits, nil
			}

			digits = append(digits, b)
		}
	}
}
