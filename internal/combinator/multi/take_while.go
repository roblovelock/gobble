package multi

import (
	"gobble/internal/parser"
	"io"
)

func TakeWhileMN[R parser.Reader, T any](p parser.Parser[R, T], m, n int, predicate parser.Predicate[T]) parser.Parser[R, []T] {
	return func(in R) ([]T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		result := make([]T, 0)

		for i := 0; i < n; i++ {
			r, err := p(in)
			if err == nil && !predicate(r) {
				err = parser.ErrNotMatched
			}
			if err != nil {
				if len(result) < m {
					_, _ = in.Seek(currentOffset, io.SeekStart)
					return nil, err
				}
				break
			}
			result = append(result, r)
		}

		return result, nil
	}
}
