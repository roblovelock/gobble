package ascii

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"io"
)

// BlankSpace returns a single ASCII blank space character: [ \t\r\n]
func BlankSpace() parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, err
		}

		if !IsBlankSpace(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return 0, errors.ErrNotMatched
		}

		return b, nil
	}
}

// BlankSpace0 returns zero or more ASCII blank space characters: [ \t\r\n]
func BlankSpace0() parser.Parser[parser.Reader, []byte] {
	return bytes.TakeWhile(IsBlankSpace)
}

// BlankSpace1 returns one or more ASCII blank space characters: [ \t\r\n]
func BlankSpace1() parser.Parser[parser.Reader, []byte] {
	return bytes.TakeWhile1(IsBlankSpace)
}

// SkipBlankSpace skips a single ASCII blank space character: [ \t\r\n]
func SkipBlankSpace() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.Skip(IsBlankSpace)
}

// SkipBlankSpace0 skips zero or more ASCII blank space characters: [ \t\r\n]
func SkipBlankSpace0() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.SkipWhile(IsBlankSpace)
}

// SkipBlankSpace1 skips one or more ASCII blank space characters: [ \t\r\n]
func SkipBlankSpace1() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.SkipWhile1(IsBlankSpace)
}
