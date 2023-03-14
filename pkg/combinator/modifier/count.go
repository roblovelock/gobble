package modifier

import "github.com/roblovelock/gobble/pkg/parser"

type (
	countParserConstraint[T any] interface {
		[]T | map[any]T
	}
	countParser[R parser.Reader, V any, T countParserConstraint[V]] struct {
		parser parser.Parser[R, T]
	}
)

func (o *countParser[R, V, T]) Parse(in R) (int, error) {
	v, err := o.parser.Parse(in)
	if err != nil {
		return 0, err
	}
	return len(v), nil
}

func (o *countParser[R, V, T]) ParseBytes(in []byte) (int, []byte, error) {
	v, out, err := o.parser.ParseBytes(in)
	if err != nil {
		return 0, in, err
	}
	return len(v), out, nil
}

// Count will return the length of the value returned from the parser
func Count[R parser.Reader, V any, T countParserConstraint[V]](p parser.Parser[R, T]) parser.Parser[R, int] {
	return &countParser[R, V, T]{parser: p}
}
