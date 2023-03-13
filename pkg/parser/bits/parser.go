package bits

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	bitsParser[T any] struct {
		parser parser.Parser[parser.BitReader, T]
	}
)

func (o *bitsParser[T]) Parse(in parser.Reader) (T, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	reader := &bitReader{Reader: in}
	result, err := o.parser.Parse(reader)
	if err != nil {
		var t T
		return t, err
	}

	if !reader.isAligned() {
		_, _ = in.Seek(currentOffset, io.SeekStart)
		return result, ErrRemainingBits
	}

	return result, nil
}

func Bits[T any](p parser.Parser[parser.BitReader, T]) parser.Parser[parser.Reader, T] {
	return &bitsParser[T]{parser: p}
}
