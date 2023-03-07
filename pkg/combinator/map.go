package combinator

import (
	"github.com/roblovelock/gobble/pkg/parser"
)

type MapFunc[T, V any] func(T) (V, error)

func Map[R parser.Reader, T, V any](p parser.Parser[R, T], mapFunc MapFunc[T, V]) parser.Parser[R, V] {
	return func(in R) (V, error) {
		t, err := p(in)
		if err != nil {
			var v V
			return v, err
		}
		return mapFunc(t)
	}
}
