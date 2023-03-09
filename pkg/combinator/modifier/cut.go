package modifier

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
)

func Cut[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		r, err := p(in)
		if err != nil {
			err = errors.NewFatalError(err)
		}
		return r, err
	}
}
