package bits

import (
	"github.com/roblovelock/gobble/pkg/combinator"
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"math/bits"
)

type (
	takeParserConstraint interface {
		uint8 | uint16 | uint32 | uint64 | uint
	}

	takeParser[T takeParserConstraint] struct {
		n uint8
	}
)

func (o *takeParser[T]) Parse(in parser.BitReader) (T, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	b, _, err := in.ReadBits(o.n)
	if err != nil {
		var t T
		_, _ = in.Seek(currentOffset, io.SeekStart)
		return t, err
	}

	return T(b), nil
}

func (*takeParser[T]) ParseBytes(in []byte) (T, []byte, error) {
	return 0, in, errors.ErrNotSupported
}

func Take[T takeParserConstraint](n uint8) parser.Parser[parser.BitReader, T] {
	var t T
	l := bits.Len64(uint64(t) - 1)

	if int(n) > l {
		return combinator.Fail[parser.BitReader, T](ErrBitsOverflow)
	}

	return &takeParser[T]{n: n}
}
