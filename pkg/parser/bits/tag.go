package bits

import (
	"github.com/roblovelock/gobble/pkg/combinator"
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"math/bits"
)

type (
	tagParserConstraint interface {
		uint8 | uint16 | uint32 | uint64 | uint
	}

	tagParser[T tagParserConstraint] struct {
		n   uint8
		tag T
	}
)

func (o *tagParser[T]) Parse(in parser.BitReader) (T, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	b, _, err := in.ReadBits(o.n)
	if err != nil {
		var t T
		_, _ = in.Seek(currentOffset, io.SeekStart)
		return t, err
	}
	if T(b) != o.tag {
		var t T
		_, _ = in.Seek(currentOffset, io.SeekStart)
		return t, errors.ErrNotMatched
	}

	return o.tag, nil

}

func Tag[T tagParserConstraint](n uint8, tag T) parser.Parser[parser.BitReader, T] {
	var t T
	l := bits.Len64(uint64(t) - 1)

	if int(n) > l {
		return combinator.Fail[parser.BitReader, T](ErrBitsOverflow)
	}

	return &tagParser[T]{n: n, tag: tag}
}
