package combinator

import (
	"gobble/internal/parser"
	"io"
)

func Not[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, parser.Empty] {
	return func(in R) (parser.Empty, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		_, err := p(in)
		_, _ = in.Seek(currentOffset, io.SeekStart)
		if err != nil {
			return parser.Empty{}, nil
		}
		return parser.Empty{}, parser.ErrNotMatched
	}
}
