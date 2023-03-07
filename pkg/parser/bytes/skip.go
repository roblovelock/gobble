package bytes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

func Skip(p parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return func(in parser.Reader) (parser.Empty, error) {
		b, err := in.ReadByte()
		if err != nil {
			return parser.Empty{}, err
		}

		if !p(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return parser.Empty{}, parser.ErrNotMatched
		}

		return parser.Empty{}, nil
	}
}

func Skip0(p parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return func(in parser.Reader) (parser.Empty, error) {
		for {
			b, err := in.ReadByte()
			if err != nil || !p(b) {
				break
			}
		}

		_, _ = in.Seek(-1, io.SeekCurrent)
		return parser.Empty{}, nil
	}
}

func Skip1(p parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return func(in parser.Reader) (parser.Empty, error) {
		b, err := in.ReadByte()
		if err != nil {
			return parser.Empty{}, err
		}

		if !p(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return parser.Empty{}, parser.ErrNotMatched
		}

		for {
			b, err := in.ReadByte()
			if err != nil || !p(b) {
				break
			}
		}

		_, _ = in.Seek(-1, io.SeekCurrent)
		return parser.Empty{}, nil
	}
}
