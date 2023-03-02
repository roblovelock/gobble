package combinator

import (
	"gobble/internal/parser"
	"io"
)

func Peek[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		t, err := p(in)
		_, _ = in.Seek(currentOffset, io.SeekStart)
		return t, err
	}
}
