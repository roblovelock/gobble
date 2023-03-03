package parser

import (
	"errors"
	"io"
)

var (
	ErrNotMatched = errors.New("not matched") // parser didn't match input
)

type (
	Reader interface {
		io.ReadSeeker
		io.ByteReader
		io.RuneReader
	}

	BitReader interface {
		Reader
		ReadBits(uint8) (uint64, uint8, error)
		ReadBool() (bool, error)
	}

	Empty                   struct{}
	Parser[R Reader, T any] func(in R) (T, error)
	Predicate[T any]        func(T) bool
	Accumulator[T, R any]   func(R, T) R
	Pair[A, B any]          struct {
		First  A
		Second B
	}
)

func Untyped[R Reader, T any](p Parser[R, T]) Parser[R, interface{}] {
	return func(in R) (interface{}, error) {
		return p(in)
	}
}
