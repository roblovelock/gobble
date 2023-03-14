package modifier

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	mapParser[R parser.Reader, T, V any] struct {
		parser parser.Parser[R, T]
		fn     parser.MapFunc[T, V]
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

func (o *mapParser[R, T, V]) ParseBytes(in []byte) (V, []byte, error) {
	t, out, err := o.parser.ParseBytes(in)
	if err != nil {
		var v V
		return v, in, err
	}
	r, err := o.fn(t)
	if err != nil {
		return r, in, err
	}
	return r, out, err
}

// Map passes the output from the parser to the map function, before returning the mapped result.
func Map[R parser.Reader, T, V any](p parser.Parser[R, T], mapFunc parser.MapFunc[T, V]) parser.Parser[R, V] {
	return &mapParser[R, T, V]{parser: p, fn: mapFunc}
}
