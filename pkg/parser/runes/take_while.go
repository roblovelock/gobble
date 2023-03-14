package runes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"math"
	"strings"
	"unicode/utf8"
)

type (
	takeWhileMinMaxParser struct {
		min       int
		max       int
		predicate parser.Predicate[rune]
	}
	takeWhile struct {
		predicate parser.Predicate[rune]
	}
)

func (o *takeWhileMinMaxParser) Parse(in parser.Reader) (string, error) {
	builder := strings.Builder{}
	builder.Grow(o.min)

	for i := 0; i < o.max; i++ {
		r, n, err := in.ReadRune()
		if err == nil && !o.predicate(r) {
			_, _ = in.Seek(-int64(n), io.SeekCurrent)
			err = errors.ErrNotMatched
		}
		if err != nil {
			if builder.Len() < o.min {
				_, _ = in.Seek(-int64(n), io.SeekCurrent)
				return "", err
			}
			break
		}
		builder.WriteRune(r)
	}

	return builder.String(), nil
}

func (o *takeWhileMinMaxParser) ParseBytes(in []byte) (string, []byte, error) {
	size := 0
	i := 0
	for ; i < o.max; i++ {
		if len(in) < size {
			break
		}
		if c := in[size]; c < utf8.RuneSelf && o.predicate(rune(c)) {
			size++
			continue
		}
		ch, s := utf8.DecodeRune(in)
		if o.predicate(ch) {
			size += s
			continue
		}
		break
	}

	if i < o.min {
		if len(in) < size {
			return "", in, io.EOF
		}
		return "", in, errors.ErrNotMatched
	}

	return string(in[:size]), in[size:], nil
}

func (o *takeWhile) Parse(in parser.Reader) (string, error) {
	builder := strings.Builder{}
	for {
		r, i, err := in.ReadRune()
		if err != nil || !o.predicate(r) {
			_, _ = in.Seek(-int64(i), io.SeekCurrent)
			break
		}
		builder.WriteRune(r)
	}

	return builder.String(), nil
}

func (o *takeWhile) ParseBytes(in []byte) (string, []byte, error) {
	size := 0
	for {
		if len(in) < size {
			break
		}
		if c := in[size]; c < utf8.RuneSelf && o.predicate(rune(c)) {
			size++
			continue
		}
		ch, s := utf8.DecodeRune(in)
		if o.predicate(ch) {
			size += s
			continue
		}
		break
	}

	return string(in[:size]), in[size:], nil
}

// TakeWhile returns a string containing zero or more Returns a string containing that match the predicate.
//   - If the input matches the predicate, it will return the matched runes.
//   - If the input is empty, it will return an empty string
//   - If the input doesn't match the predicate, it will return an empty string
func TakeWhile(predicate parser.Predicate[rune]) parser.Parser[parser.Reader, string] {
	return &takeWhile{predicate: predicate}
}

// TakeWhile1 returns a string containing one or more runes that match the predicate.
//   - If the input matches the predicate, it will return the matched runes.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the predicate, it will return errors.ErrNotMatched
func TakeWhile1(predicate parser.Predicate[rune]) parser.Parser[parser.Reader, string] {
	return TakeWhileMinMax(1, math.MaxInt, predicate)
}

// TakeWhileMinMax returns a string of length (m <= len <= n) containing runes that match the predicate.
//   - If the input matches the predicate, it will return the matched runes.
//   - If the input is empty and m > 0, it will return io.EOF
//   - If the number of matched bytes < m, it will return errors.ErrNotMatched
func TakeWhileMinMax(min, max int, predicate parser.Predicate[rune]) parser.Parser[parser.Reader, string] {
	return &takeWhileMinMaxParser{min: min, max: max, predicate: predicate}
}
