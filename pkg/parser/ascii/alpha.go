// Package ascii provides parsers for recognizing ascii bytes
package ascii

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Alpha matches a single ASCII letter character: [a-zA-Z]
//   - If the input matches a letter character, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match a letter character, it will return parser.ErrNotMatched
func Alpha() parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, err
		}

		if IsLetter(b) {
			return b, nil
		}

		_, _ = in.Seek(-1, io.SeekCurrent)
		return 0, errors.ErrNotMatched
	}
}

// Alpha0 matches zero or more ASCII letter characters: [a-zA-Z]
//   - If the input matches a letter character, it will return a slice of all matched letters.
//   - If the input is empty, it will return an empty slice.
//   - If the input doesn't match a letter character, it will return an empty slice.
func Alpha0() parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		digits := make([]byte, 0)

		for {
			b, err := in.ReadByte()
			if err != nil {
				return digits, nil
			}

			if !IsLetter(b) {
				_, _ = in.Seek(-1, io.SeekCurrent)
				return digits, nil
			}
			digits = append(digits, b)
		}
	}
}

// Alpha1 matches one or more ASCII letter characters: [a-zA-Z]
//   - If the input matches a letter character, it will return a slice of all matched letters.
//   - If the input is empty, it will return io.EOF.
//   - If the input doesn't match a letter character, it will return parser.ErrNotMatched.
func Alpha1() parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return nil, err
		}

		if !IsLetter(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return nil, errors.ErrNotMatched
		}

		digits := []byte{b}

		for {
			b, err := in.ReadByte()
			if err != nil {
				return digits, nil
			}

			if !IsLetter(b) {
				_, _ = in.Seek(-1, io.SeekCurrent)
				return digits, nil
			}

			digits = append(digits, b)
		}
	}
}
