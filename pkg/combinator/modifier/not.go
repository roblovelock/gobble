package modifier

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	notParser[R parser.Reader, T any] struct {
		parser parser.Parser[R, T]
	}
)

func (o *notParser[R, T]) Parse(in R) (parser.Empty, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	_, err := o.parser.Parse(in)
	_, _ = in.Seek(currentOffset, io.SeekStart)
	if err != nil {
		return nil, nil
	}
	return nil, errors.ErrNotMatched
}

func (o *notParser[R, T]) ParseBytes(in []byte) (parser.Empty, []byte, error) {
	_, _, err := o.parser.ParseBytes(in)
	if err != nil {
		return nil, in, nil
	}
	return nil, in, errors.ErrNotMatched
}

// Not returns a result only if the parser returns an error. It doesn't consume any input
func Not[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, parser.Empty] {
	return &notParser[R, T]{parser: p}
}
