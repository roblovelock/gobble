package ascii

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
	"math"
)

var whitespace = [math.MaxUint8]bool{}

func init() {
	whitespace[' '] = true
	whitespace['\r'] = true
	whitespace['\n'] = true
	whitespace['\t'] = true
	whitespace['\v'] = true
	whitespace['\f'] = true
}

func Whitespace() parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, err
		}

		if !whitespace[b] {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return 0, parser.ErrNotMatched
		}

		return b, nil
	}
}

func Whitespace0() parser.Parser[parser.Reader, []byte] {
	return bytes.OneOf0(' ', '\r', '\n', '\t')
}

func Whitespace1() parser.Parser[parser.Reader, []byte] {
	return bytes.OneOf1(' ', '\r', '\n', '\t')
}

func SkipWhitespace() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.Skip(func(b byte) bool { return whitespace[b] })
}

func SkipWhitespace0() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.Skip0(func(b byte) bool { return whitespace[b] })
}

func SkipWhitespace1() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.Skip1(func(b byte) bool { return whitespace[b] })
}
