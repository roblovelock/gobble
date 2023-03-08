package combinator

import "github.com/roblovelock/gobble/pkg/parser"

// Success always succeeds. It returns the provided value without consuming any input.
func Success[R parser.Reader, T any](value T) parser.Parser[R, T] {
	return func(in R) (T, error) {
		return value, nil
	}
}
