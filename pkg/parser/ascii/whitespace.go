package ascii

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
)

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

func Whitespace0() parser.Parser[parser.Reader, []byte] {
	return bytes.TakeWhile(IsWhitespace)
}

func Whitespace1() parser.Parser[parser.Reader, []byte] {
	return bytes.TakeWhile1(IsWhitespace)
}

func SkipWhitespace() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.Skip(IsWhitespace)
}

func SkipWhitespace0() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.SkipWhile(IsWhitespace)
}

func SkipWhitespace1() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.SkipWhile1(IsWhitespace)
}
