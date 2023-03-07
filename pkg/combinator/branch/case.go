package branch

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

func Case[R parser.Reader, C comparable, T any](
	p parser.Parser[R, C], parsers map[C]parser.Parser[R, T],
) parser.Parser[R, T] {
	return func(in R) (T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		c, err := p(in)
		if err != nil {
			var t T
			return t, err
		}
		_, _ = in.Seek(currentOffset, io.SeekStart)
		p, ok := parsers[c]
		if !ok {
			var t T
			return t, parser.ErrNotMatched
		}

		return p(in)
	}
}

func PeekCase[R parser.Reader, T any](parsers map[byte]parser.Parser[R, T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		b, err := in.ReadByte()
		if err != nil {
			var t T
			return t, err
		}
		_, _ = in.Seek(-1, io.SeekCurrent)
		p, ok := parsers[b]
		if !ok {
			var t T
			return t, parser.ErrNotMatched
		}

		return p(in)
	}
}
