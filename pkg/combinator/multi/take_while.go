package multi

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	takeWhileMinMaxParser[R parser.Reader, T any] struct {
		parser    parser.Parser[R, T]
		min       int
		max       int
		predicate parser.Predicate[T]
	}
)

func (o *takeWhileMinMaxParser[R, T]) Parse(in R) ([]T, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	result := make([]T, 0)

	for i := 0; i < o.max; i++ {
		r, err := o.parser.Parse(in)
		if err == nil && !o.predicate(r) {
			err = errors.ErrNotMatched
		}
		if err != nil {
			if len(result) < o.min {
				_, _ = in.Seek(currentOffset, io.SeekStart)
				return nil, err
			}
			break
		}
		result = append(result, r)
	}

	return result, nil
}

func (o *takeWhileMinMaxParser[R, T]) ParseBytes(in []byte) ([]T, []byte, error) {
	var r T
	var err error
	out := in
	result := make([]T, 0, o.min)
	for i := 0; i < o.max; i++ {
		r, out, err = o.parser.ParseBytes(out)
		if err != nil {
			if len(result) < o.min {
				return nil, in, err
			}
			return result, in, nil
		}

		if !o.predicate(r) {
			if len(result) < o.min {
				return nil, in, errors.ErrNotMatched
			}
			return result, in, nil
		}
		result = append(result, r)
	}

	return result, out, nil
}

func TakeWhileMinMax[R parser.Reader, T any](p parser.Parser[R, T], min, max int, predicate parser.Predicate[T]) parser.Parser[R, []T] {
	return &takeWhileMinMaxParser[R, T]{parser: p, min: min, max: max, predicate: predicate}
}
