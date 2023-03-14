package modifier

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
)

type (
	cutParser[R parser.Reader, T any] struct {
		parser parser.Parser[R, T]
	}
)

func (o *cutParser[R, T]) Parse(in R) (T, error) {
	r, err := o.parser.Parse(in)
	if err != nil {
		err = errors.NewFatalError(err)
	}
	return r, err
}

func (o *cutParser[R, T]) ParseBytes(in []byte) (T, []byte, error) {
	t, out, err := o.parser.ParseBytes(in)
	if err != nil {
		return t, in, errors.NewFatalError(err)
	}
	return t, out, nil
}

func Cut[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, T] {
	return &cutParser[R, T]{parser: p}
}
