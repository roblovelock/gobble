package runes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/utils"
	"io"
)

type (
	oneOfParser struct {
		runes map[rune]bool
	}
)

func (o *oneOfParser) Parse(in parser.Reader) (rune, error) {
	r, i, err := in.ReadRune()
	if err != nil {
		_, _ = in.Seek(-int64(i), io.SeekCurrent)
		return 0, err
	}

	if o.runes[r] {
		return r, nil
	}

	_, _ = in.Seek(-int64(i), io.SeekCurrent)
	return 0, errors.ErrNotMatched
}

// OneOf matches one of the argument runes
//   - If the input matches the argument, it will return a single matched rune.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func OneOf(runes ...rune) parser.Parser[parser.Reader, rune] {
	return &oneOfParser{runes: utils.NewLookupMap(runes)}
}

// OneOf0 matches zero or more runes matching one of the argument runes
//   - If the input matches the argument, it will return a string of all matched runes.
//   - If the input is empty, it will return an empty string.
//   - If the input doesn't match, it will return an empty string.
func OneOf0(runes ...rune) parser.Parser[parser.Reader, string] {
	lookup := utils.NewLookupMap(runes)
	return TakeWhile(func(r rune) bool {
		return lookup[r]
	})
}

// OneOf1 matches one or more runes matching one of the argument runes
//   - If the input matches the argument, it will return a string of all matched runes.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func OneOf1(runes ...rune) parser.Parser[parser.Reader, string] {
	lookup := utils.NewLookupMap(runes)
	return TakeWhile1(func(r rune) bool {
		return lookup[r]
	})
}
