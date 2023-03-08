package bytes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

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

func SkipN(n uint) parser.Parser[parser.Reader, parser.Empty] {
	return func(in parser.Reader) (parser.Empty, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		_, err := in.Seek(int64(n), io.SeekCurrent)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
		}
		return nil, err
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
		return nil, nil
	}
}

func Skip1(p parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
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
			b, err := in.ReadByte()
			if err != nil || !p(b) {
				break
			}
		}

		_, _ = in.Seek(-1, io.SeekCurrent)
		return nil, nil
	}
}
