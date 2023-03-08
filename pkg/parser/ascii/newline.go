// Package ascii provides parsers for recognizing ascii bytes
package ascii

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
)

// CRLF matches the sequence `\r\n`
//   - If the input matches `\r\n`, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match, it will return parser.ErrNotMatched
func CRLF() parser.Parser[parser.Reader, []byte] {
	return bytes.Tag([]byte{'\r', '\n'})
}

// Newline matches a new line: `\n`
//   - If the input matches, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match, it will return parser.ErrNotMatched.
func Newline() parser.Parser[parser.Reader, byte] {
	return bytes.Byte('\n')
}

// LineEnding matches the end of a line: `\n` and `\r\n`
//   - If the input matches, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match, it will return parser.ErrNotMatched.
func LineEnding() parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return nil, err
		}
		if b == '\n' {
			return []byte{'\n'}, nil
		}
		if b != '\r' {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return nil, parser.ErrNotMatched
		}
		b, err = in.ReadByte()
		if err != nil {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return nil, err
		}
		if b != '\n' {
			_, _ = in.Seek(-2, io.SeekCurrent)
			return nil, parser.ErrNotMatched
		}

		return []byte{'\r', '\n'}, nil
	}
}
