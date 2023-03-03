package combinator

import "gobble/pkg/parser"

func Optional[R parser.Reader, T any](parser parser.Parser[R, T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		v, _ := parser(in)
		return v, nil
	}
}
