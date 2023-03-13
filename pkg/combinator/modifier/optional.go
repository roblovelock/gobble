package modifier

import "github.com/roblovelock/gobble/pkg/parser"

type (
	optionalParser[R parser.Reader, T any] struct {
		parser parser.Parser[R, T]
	}
)

func (o *optionalParser[R, T]) Parse(in R) (T, error) {
	v, _ := o.parser.Parse(in)
	return v, nil
}

// Optional will call the parser and suppress any error returned
func Optional[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, T] {
	return &optionalParser[R, T]{parser: p}
}
