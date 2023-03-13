package modifier

import (
	"github.com/roblovelock/gobble/pkg/parser"
)

type (
	valueParser[R parser.Reader, T, V any] struct {
		parser parser.Parser[R, T]
		value  V
	}
)

func (o *valueParser[R, T, V]) Parse(in R) (V, error) {
	if _, err := o.parser.Parse(in); err != nil {
		var v V
		return v, err
	}
	return o.value, nil
}

// Value returns the provided value if the parser succeeds.
func Value[R parser.Reader, T, V any](p parser.Parser[R, T], value V) parser.Parser[R, V] {
	return &valueParser[R, T, V]{parser: p, value: value}
}
