package combinator

import "github.com/roblovelock/gobble/pkg/parser"

type (
	failParser[R parser.Reader, T any] struct {
		err error
	}
)

func (o *failParser[R, T]) Parse(_ R) (T, error) {
	var t T
	return t, o.err
}

// Fail always fails. It returns the provided err without consuming any input.
func Fail[R parser.Reader, T any](err error) parser.Parser[R, T] {
	return &failParser[R, T]{err: err}
}
