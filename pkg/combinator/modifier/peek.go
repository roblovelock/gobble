package modifier

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	peekParser[R parser.Reader, T any] struct {
		parser parser.Parser[R, T]
	}
)

func (o *peekParser[R, T]) Parse(in R) (T, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	t, err := o.parser.Parse(in)
	_, _ = in.Seek(currentOffset, io.SeekStart)
	return t, err
}

// Peek returns the result of the parser without consuming the input.
func Peek[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, T] {
	return &peekParser[R, T]{parser: p}
}
