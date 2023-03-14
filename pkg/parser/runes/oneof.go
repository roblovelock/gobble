package runes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/utils"
	"io"
	"unicode/utf8"
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

func (o *oneOfParser) ParseBytes(in []byte) (ch rune, out []byte, err error) {
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

	if o.runes[ch] {
		return
	}

	return 0, in, errors.ErrNotMatched
}

// OneOf matches one of the argument runes
//   - If the input matches the argument, it will return a single matched rune.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return errors.ErrNotMatched
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
//   - If the input doesn't match the argument, it will return errors.ErrNotMatched
func OneOf1(runes ...rune) parser.Parser[parser.Reader, string] {
	lookup := utils.NewLookupMap(runes)
	return TakeWhile1(func(r rune) bool {
		return lookup[r]
	})
}
