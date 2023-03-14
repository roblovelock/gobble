package modifier

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	verifyParser[R parser.Reader, T any] struct {
		parser    parser.Parser[R, T]
		predicate parser.Predicate[T]
	}
)

func (o *verifyParser[R, T]) Parse(in R) (T, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	r, err := o.parser.Parse(in)
	if err != nil {
		return r, err
	}

	if !o.predicate(r) {
		var r T
		_, _ = in.Seek(currentOffset, io.SeekStart)
		return r, errors.ErrNotMatched
	}

	return r, nil
}

func (o *verifyParser[R, T]) ParseBytes(in []byte) (T, []byte, error) {
	r, out, err := o.parser.ParseBytes(in)
	if err != nil {
		return r, in, err
	}

	if !o.predicate(r) {
		var r T
		return r, in, errors.ErrNotMatched
	}

	return r, out, nil
}

func Verify[R parser.Reader, T any](p parser.Parser[R, T], predicate parser.Predicate[T]) parser.Parser[R, T] {
	return &verifyParser[R, T]{parser: p, predicate: predicate}
}
