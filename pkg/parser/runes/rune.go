package runes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
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

// Rune matches a single rune
//
// The input data will be compared to the match argument.
//   - If the input matches the argument, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func Rune(r rune) parser.Parser[parser.Reader, rune] {
	return &runeParser{r: r}
}
