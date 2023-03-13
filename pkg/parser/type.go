package parser

import (
	"github.com/roblovelock/gobble/pkg/errors"
)

type (
	untypedParser[R Reader, T any] struct {
		parser Parser[R, T]
	}

	typedParser[R Reader, T any] struct {
		parser Parser[R, interface{}]
	}
)

func (o *untypedParser[R, T]) Parse(in R) (interface{}, error) {
	return o.parser.Parse(in)
}

func (o *typedParser[R, T]) Parse(in R) (T, error) {
	r, err := o.parser.Parse(in)
	val, ok := r.(T)
	if !ok {
		var t T
		return t, errors.ErrNotMatched
	}
	return val, err
}

func Untyped[R Reader, T any](p Parser[R, T]) Parser[R, interface{}] {
	return &untypedParser[R, T]{parser: p}
}

func Typed[R Reader, T any](p Parser[R, interface{}]) Parser[R, T] {
	return &typedParser[R, T]{parser: p}
}
