package bytes

import (
	"gobble/pkg/parser"
	"io"
)

// Byte matches a single byte
//
// The input data will be compared to the match argument.
//   - If the input matches the argument, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func Byte(match byte) parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, err
		}

		if b != match {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return 0, parser.ErrNotMatched
		}

		return b, nil
	}
}
