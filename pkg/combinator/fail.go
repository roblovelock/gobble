package combinator

import "github.com/roblovelock/gobble/pkg/parser"

// Fail always fails. It returns the provided err without consuming any input.
func Fail[R parser.Reader, T any](err error) parser.Parser[R, T] {
	return func(in R) (T, error) {
		var t T
		return t, err
	}
}
