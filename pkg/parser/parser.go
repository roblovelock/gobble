package parser

import (
	"io"
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

	Parser[R Reader, T any] interface {
		Parse(in R) (T, error)
		ParseBytes(in []byte) (T, []byte, error)
	}

	Empty                 *struct{}
	Predicate[T any]      func(T) bool
	MapFunc[T, V any]     func(T) (V, error)
	Accumulator[T, R any] func(R, T) R
	Pair[A, B any]        struct {
		First  A
		Second B
	}
)
