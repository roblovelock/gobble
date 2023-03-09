package bytes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// OneOf matches one of the argument bytes
//   - If the input matches the argument, it will return a single matched byte.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func OneOf(bytes ...byte) parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, err
		}

		for _, v := range bytes {
			if b == v {
				return b, nil
			}
		}

		_, _ = in.Seek(-1, io.SeekCurrent)
		return 0, errors.ErrNotMatched
	}
}

// OneOf0 matches zero or more bytes matching one of the argument bytes
//   - If the input matches the argument, it will return a slice of all matched bytes.
//   - If the input is empty, it will return an empty slice.
//   - If the input doesn't match, it will return an empty slice.
func OneOf0(bytes ...byte) parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		result := make([]byte, 0)
		for {
			b, err := OneOf(bytes...)(in)
			if err != nil {
				return result, nil
			}
			result = append(result, b)
		}
	}
}

// OneOf1 matches one or more bytes matching one of the argument bytes
//   - If the input matches the argument, it will return a slice of all matched bytes.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func OneOf1(bytes ...byte) parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		result := make([]byte, 0)
		for {
			b, err := OneOf(bytes...)(in)
			if err != nil {
				if len(result) == 0 {
					return nil, err
				}
				break
			}
			result = append(result, b)
		}
		return result, nil
	}
}
