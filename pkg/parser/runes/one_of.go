package runes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"strings"
)

// OneOf matches one of the argument runes
//   - If the input matches the argument, it will return a single matched rune.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func OneOf(runes ...rune) parser.Parser[parser.Reader, rune] {
	return func(in parser.Reader) (rune, error) {
		r, i, err := in.ReadRune()
		if err != nil {
			_, _ = in.Seek(-int64(i), io.SeekCurrent)
			return 0, err
		}

		for _, v := range runes {
			if r == v {
				return r, nil
			}
		}

		_, _ = in.Seek(-int64(i), io.SeekCurrent)
		return 0, errors.ErrNotMatched
	}
}

// OneOf0 matches zero or more runes matching one of the argument runes
//   - If the input matches the argument, it will return a string of all matched runes.
//   - If the input is empty, it will return an empty string.
//   - If the input doesn't match, it will return an empty string.
func OneOf0(runes ...rune) parser.Parser[parser.Reader, string] {
	return func(in parser.Reader) (string, error) {
		var builder strings.Builder
		for {
			r, err := OneOf(runes...)(in)
			if err != nil {
				return builder.String(), nil
			}
			_, _ = builder.WriteRune(r)
		}

	}
}

// OneOf1 matches one or more runes matching one of the argument runes
//   - If the input matches the argument, it will return a string of all matched runes.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func OneOf1(runes ...rune) parser.Parser[parser.Reader, string] {
	return func(in parser.Reader) (string, error) {
		var builder strings.Builder
		for {
			r, err := OneOf(runes...)(in)
			if err != nil {
				if builder.Len() == 0 {
					return "", err
				}
				break
			}
			_, _ = builder.WriteRune(r)
		}
		return builder.String(), nil
	}
}
