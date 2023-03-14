package runes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"unicode/utf8"
)

type (
	runeParser struct {
		r rune
	}
)

func (o *runeParser) Parse(in parser.Reader) (rune, error) {
	b, i, err := in.ReadRune()
	if err != nil {
		_, _ = in.Seek(-int64(i), io.SeekCurrent)
		return 0, err
	}

	if b != o.r {
		_, _ = in.Seek(-int64(i), io.SeekCurrent)
		return 0, errors.ErrNotMatched
	}

	return b, nil
}

func (o *runeParser) ParseBytes(in []byte) (ch rune, out []byte, err error) {
	if len(in) == 0 {
		return 0, in, io.EOF
	}

	if c := in[0]; c < utf8.RuneSelf {
		ch = rune(c)
		out = in[1:]
	} else {
		var size int
		ch, size = utf8.DecodeRune(in)
		out = in[size:]
	}

	if o.r == ch {
		return
	}

	return 0, in, errors.ErrNotMatched
}

// Rune matches a single rune
//
// The input data will be compared to the match argument.
//   - If the input matches the argument, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return errors.ErrNotMatched
func Rune(r rune) parser.Parser[parser.Reader, rune] {
	return &runeParser{r: r}
}
