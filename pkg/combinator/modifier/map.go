package modifier

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	MapFunc[T, V any] func(T) (V, error)

	mapParser[R parser.Reader, T, V any] struct {
		parser parser.Parser[R, T]
		fn     MapFunc[T, V]
	}
)

func (o *mapParser[R, T, V]) Parse(in R) (V, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	t, err := o.parser.Parse(in)
	if err != nil {
		var v V
		return v, err
	}
	r, err := o.fn(t)
	if err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
	}
	return r, err
}

// Map passes the output from the parser to the map function, before returning the mapped result.
func Map[R parser.Reader, T, V any](p parser.Parser[R, T], mapFunc MapFunc[T, V]) parser.Parser[R, V] {
	return &mapParser[R, T, V]{parser: p, fn: mapFunc}
}
