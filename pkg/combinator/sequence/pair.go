package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	pairParser[R parser.Reader, F, S any] struct {
		first  parser.Parser[R, F]
		second parser.Parser[R, S]
	}

	separatedPairParser[R parser.Reader, F, S, T any] struct {
		first     parser.Parser[R, F]
		separator parser.Parser[R, T]
		second    parser.Parser[R, S]
	}
)

func (o *pairParser[R, F, S]) Parse(in R) (parser.Pair[F, S], error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	f, err := o.first.Parse(in)
	if err != nil {
		var r parser.Pair[F, S]
		return r, err
	}
	s, err := o.second.Parse(in)
	if err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
		var r parser.Pair[F, S]
		return r, err
	}
	return parser.Pair[F, S]{First: f, Second: s}, err
}

func (o *separatedPairParser[R, F, S, T]) Parse(in R) (parser.Pair[F, S], error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	f, err := o.first.Parse(in)
	if err != nil {
		var r parser.Pair[F, S]
		return r, err
	}

	if _, err := o.separator.Parse(in); err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
		var r parser.Pair[F, S]
		return r, err
	}

	s, err := o.second.Parse(in)
	if err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
		var r parser.Pair[F, S]
		return r, err
	}
	return parser.Pair[F, S]{First: f, Second: s}, err
}

// Pair Gets an object from the first parser, then gets another object from the second parser.
func Pair[R parser.Reader, F, S any](
	first parser.Parser[R, F], second parser.Parser[R, S],
) parser.Parser[R, parser.Pair[F, S]] {
	return &pairParser[R, F, S]{first: first, second: second}
}

// SeparatedPair Gets an object from the first parser, then checks the separator parser, before getting another object
// from the final parser.
func SeparatedPair[R parser.Reader, F, S, T any](
	first parser.Parser[R, F], separator parser.Parser[R, T], second parser.Parser[R, S],
) parser.Parser[R, parser.Pair[F, S]] {
	return &separatedPairParser[R, F, S, T]{first: first, separator: separator, second: second}
}
