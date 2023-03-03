package combinator

import (
	"gobble/pkg/parser"
)

func Verify[R parser.Reader, T any](p parser.Parser[R, T], predicate parser.Predicate[T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		r, err := p(in)
		if err != nil {
			return r, err
		}

		if !predicate(r) {
			var r T
			return r, parser.ErrNotMatched
		}

		return r, nil
	}
}
