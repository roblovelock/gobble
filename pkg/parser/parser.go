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

	Parser[R Reader, T any] func(in R) (T, error)

	Empty                 *struct{}
	Predicate[T any]      func(T) bool
	Accumulator[T, R any] func(R, T) R
	Pair[A, B any]        struct {
		First  A
		Second B
	}
)

func Ptr[R Reader, T any](p *Parser[R, T]) Parser[R, T] {
	return func(in R) (T, error) {
		return (*p)(in)
	}
}
