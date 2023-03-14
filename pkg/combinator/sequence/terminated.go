package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	terminatedParser[R parser.Reader, F, S any] struct {
		first  parser.Parser[R, F]
		second parser.Parser[R, S]
	}
)

func (o *terminatedParser[R, F, S]) Parse(in R) (F, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	f, err := o.first.Parse(in)
	if err != nil {
		return f, err
	}

	if _, err := o.second.Parse(in); err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
		var r F
		return r, err
	}
	return f, err
}

func (o *terminatedParser[R, F, S]) ParseBytes(in []byte) (F, []byte, error) {
	f, out, err := o.first.ParseBytes(in)
	if err != nil {
		return f, in, err
	}
	_, out, err = o.second.ParseBytes(out)
	if err != nil {
		var r F
		return r, in, err
	}
	return f, out, err
}

// Terminated Gets an object from the first parser, then matches an object from the second parser and discards it.
func Terminated[R parser.Reader, F, S any](first parser.Parser[R, F], second parser.Parser[R, S]) parser.Parser[R, F] {
	return &terminatedParser[R, F, S]{first: first, second: second}
}
