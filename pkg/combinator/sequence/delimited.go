package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	delimitedParser[R parser.Reader, F, S, T any] struct {
		first  parser.Parser[R, F]
		second parser.Parser[R, S]
		third  parser.Parser[R, T]
	}
)

func (o *delimitedParser[R, F, S, T]) Parse(in R) (S, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	if _, err := o.first.Parse(in); err != nil {
		var r S
		return r, err
	}

	s, err := o.second.Parse(in)
	if err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
		return s, err
	}

	if _, err := o.third.Parse(in); err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
		var r S
		return r, err
	}

	return s, nil
}

func (o *delimitedParser[R, F, S, T]) ParseBytes(in []byte) (S, []byte, error) {
	_, out, err := o.first.ParseBytes(in)
	if err != nil {
		var r S
		return r, in, err
	}

	s, out, err := o.second.ParseBytes(out)
	if err != nil {
		return s, in, err
	}
	_, out, err = o.third.ParseBytes(out)
	if err != nil {
		var r S
		return r, in, err
	}

	return s, out, nil
}

// Delimited Matches an object from the first parser and discards it, then gets an object from the second parser,
// and finally matches an object from the third parser and discards it.
func Delimited[R parser.Reader, F, S, T any](
	first parser.Parser[R, F], second parser.Parser[R, S], third parser.Parser[R, T],
) parser.Parser[R, S] {
	return &delimitedParser[R, F, S, T]{first: first, second: second, third: third}
}
