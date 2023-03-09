package bytes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Skip will skip over a byte if it matches the predicate.
func Skip(p parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return func(in parser.Reader) (parser.Empty, error) {
		b, err := in.ReadByte()
		if err != nil {
			return nil, err
		}

		if !p(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return nil, parser.ErrNotMatched
		}

		return nil, nil
	}
}

// SkipWhile will skip over zero or more bytes that match the predicate.
func SkipWhile(p parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return func(in parser.Reader) (parser.Empty, error) {
		for {
			b, err := in.ReadByte()
			if err != nil {
				return nil, nil
			}
			if !p(b) {
				_, _ = in.Seek(-1, io.SeekCurrent)
				return nil, nil
			}
		}
	}
}

// SkipWhile1 will skip over one or more bytes that match the predicate.
func SkipWhile1(p parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return func(in parser.Reader) (parser.Empty, error) {
		b, err := in.ReadByte()
		if err != nil {
			return nil, err
		}

		if !p(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return nil, parser.ErrNotMatched
		}

		for {
			b, err = in.ReadByte()
			if err != nil {
				return nil, nil
			}
			if !p(b) {
				_, _ = in.Seek(-1, io.SeekCurrent)
				return nil, nil
			}
		}
	}
}
