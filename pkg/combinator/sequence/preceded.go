package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	precededParser[R parser.Reader, F, S any] struct {
		first  parser.Parser[R, F]
		second parser.Parser[R, S]
	}
)

func (o *precededParser[R, F, S]) Parse(in R) (S, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	if _, err := o.first.Parse(in); err != nil {
		var r S
		return r, err
	}
	s, err := o.second.Parse(in)
	if err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
	}
	return s, err
}

// Preceded Matches an object from the first parser and discards it, then gets an object from the second parser.
func Preceded[R parser.Reader, F, S any](first parser.Parser[R, F], second parser.Parser[R, S]) parser.Parser[R, S] {
	return &precededParser[R, F, S]{first: first, second: second}
}
