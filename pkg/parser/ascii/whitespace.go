package ascii

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
)

// Whitespace returns a single ASCII whitespace character: [ \t\r\n\v\f]
func Whitespace() parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, err
		}

		if !IsWhitespace(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return 0, parser.ErrNotMatched
		}

		return b, nil
	}
}

// Whitespace0 returns zero or more ASCII whitespace characters: [ \t\r\n\v\f]
func Whitespace0() parser.Parser[parser.Reader, []byte] {
	return bytes.TakeWhile(IsWhitespace)
}

// Whitespace1 returns one or more ASCII whitespace characters: [ \t\r\n\v\f]
func Whitespace1() parser.Parser[parser.Reader, []byte] {
	return bytes.TakeWhile1(IsWhitespace)
}

// SkipWhitespace skips a single ASCII whitespace character: [ \t\r\n\v\f]
func SkipWhitespace() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.Skip(IsWhitespace)
}

// SkipWhitespace0 skips zero or more ASCII whitespace characters: [ \t\r\n\v\f]
func SkipWhitespace0() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.SkipWhile(IsWhitespace)
}

// SkipWhitespace1 skips one or more ASCII whitespace characters: [ \t\r\n\v\f]
func SkipWhitespace1() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.SkipWhile1(IsWhitespace)
}
