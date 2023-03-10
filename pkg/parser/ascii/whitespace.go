package ascii

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
)

type (
	whitespaceParser struct {
	}
)

var whitespaceParserInstance = whitespaceParser{}

func (*whitespaceParser) Parse(in parser.Reader) (byte, error) {
	b, err := in.ReadByte()
	if err != nil {
		return 0, err
	}

	if !IsWhitespace(b) {
		_, _ = in.Seek(-1, io.SeekCurrent)
		return 0, errors.ErrNotMatched
	}

	return b, nil
}

func (o *whitespaceParser) ParseBytes(in []byte) (byte, []byte, error) {
	if len(in) == 0 {
		return 0, in, io.EOF
	}

	if !IsWhitespace(in[0]) {
		return 0, in, errors.ErrNotMatched
	}

	return in[0], in[1:], nil
}

// Whitespace returns a single ASCII whitespace character: [ \t\r\n\v\f]
func Whitespace() parser.Parser[parser.Reader, byte] {
	return &whitespaceParserInstance
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
