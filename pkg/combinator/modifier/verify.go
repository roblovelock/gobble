package modifier

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

func Verify[R parser.Reader, T any](p parser.Parser[R, T], predicate parser.Predicate[T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		r, err := p(in)
		if err != nil {
			return r, err
		}

		if !predicate(r) {
			var r T
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return r, errors.ErrNotMatched
		}

		return r, nil
	}
}
