package combinator

import "github.com/roblovelock/gobble/pkg/parser"

type (
	successParser[R parser.Reader, T any] struct {
		value T
	}
)

func (o *successParser[R, T]) Parse(_ R) (T, error) {
	return o.value, nil
}

func (o *successParser[R, T]) ParseBytes(in []byte) (T, []byte, error) {
	return o.value, in, nil
}

// Success always succeeds. It returns the provided value without consuming any input.
func Success[R parser.Reader, T any](value T) parser.Parser[R, T] {
	return &successParser[R, T]{value: value}
}
