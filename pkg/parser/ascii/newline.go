// Package ascii provides parsers for recognizing ascii bytes
package ascii

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
)

type (
	lineEndingParser struct {
	}
)

var lineEndingParserInstance = &lineEndingParser{}

func (o *lineEndingParser) Parse(in parser.Reader) ([]byte, error) {
	b, err := in.ReadByte()
	if err != nil {
		return nil, err
	}
	if b == '\n' {
		return []byte{'\n'}, nil
	}
	if b != '\r' {
		_, _ = in.Seek(-1, io.SeekCurrent)
		return nil, errors.ErrNotMatched
	}
	b, err = in.ReadByte()
	if err != nil {
		_, _ = in.Seek(-1, io.SeekCurrent)
		return nil, err
	}
	if b != '\n' {
		_, _ = in.Seek(-2, io.SeekCurrent)
		return nil, errors.ErrNotMatched
	}

	return []byte{'\r', '\n'}, nil
}

func (o *lineEndingParser) ParseBytes(in []byte) ([]byte, []byte, error) {
	if len(in) == 0 {
		return nil, in, io.EOF
	}

	if in[0] == '\n' {
		return in[:1], in[1:], nil
	}

	if in[0] == '\r' {
		if len(in) == 1 {
			return nil, in, io.EOF
		}

		if in[1] == '\n' {
			return in[:2], in[2:], nil
		}
	}

	return nil, in, errors.ErrNotMatched

}

// CRLF matches the sequence `\r\n`
//   - If the input matches `\r\n`, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match, it will return errors.ErrNotMatched
func CRLF() parser.Parser[parser.Reader, []byte] {
	return bytes.Tag([]byte{'\r', '\n'})
}

// Newline matches a new line: `\n`
//   - If the input matches, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match, it will return errors.ErrNotMatched.
func Newline() parser.Parser[parser.Reader, byte] {
	return bytes.Byte('\n')
}

// LineEnding matches the end of a line: `\n` and `\r\n`
//   - If the input matches, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match, it will return errors.ErrNotMatched.
func LineEnding() parser.Parser[parser.Reader, []byte] {
	return lineEndingParserInstance
}
