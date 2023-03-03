package combinator

import (
	"gobble/pkg/parser"
)

func Value[R parser.Reader, T, V any](p parser.Parser[R, T], val V) parser.Parser[R, V] {
	return func(in R) (V, error) {
		if _, err := p(in); err != nil {
			var v V
			return v, err
		}
		return val, nil
	}
}
