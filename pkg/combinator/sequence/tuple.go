package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	tupleParser[R parser.Reader, T any] struct {
		parsers []parser.Parser[R, T]
	}
)

func (o *tupleParser[R, T]) Parse(in R) ([]T, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)

	result := make([]T, len(o.parsers))
	for i, p := range o.parsers {
		r, err := p.Parse(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return nil, err
		}
		result[i] = r
	}

	return result, nil
}

// Tuple applies a number of parsers one by one and returns their results as a slice.
func Tuple[R parser.Reader, T any](parsers ...parser.Parser[R, T]) parser.Parser[R, []T] {
	return &tupleParser[R, T]{parsers: parsers}
}
