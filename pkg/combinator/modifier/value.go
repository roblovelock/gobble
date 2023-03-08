package modifier

import (
	"github.com/roblovelock/gobble/pkg/parser"
)

// Value returns the provided value if the parser succeeds.
func Value[R parser.Reader, T, V any](p parser.Parser[R, T], value V) parser.Parser[R, V] {
	return func(in R) (V, error) {
		if _, err := p(in); err != nil {
			var v V
			return v, err
		}
		return value, nil
	}
}
