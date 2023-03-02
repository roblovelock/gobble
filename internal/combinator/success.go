package combinator

import "gobble/internal/parser"

func Success[R parser.Reader, T any](val T) parser.Parser[R, T] {
	return func(in R) (T, error) {
		return val, nil
	}
}
