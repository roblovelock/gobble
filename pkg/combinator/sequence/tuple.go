package sequence

import (
	"gobble/pkg/parser"
	"io"
)

// Tuple applies a number of parsers one by one and returns their results as a slice.
func Tuple[R parser.Reader, T any](parsers ...parser.Parser[R, T]) parser.Parser[R, []T] {
	return func(in R) ([]T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)

		result := make([]T, len(parsers))
		for i, p := range parsers {
			r, err := p(in)
			if err != nil {
				_, _ = in.Seek(currentOffset, io.SeekStart)
				return nil, err
			}
			result[i] = r
		}

		return result, nil
	}
}
