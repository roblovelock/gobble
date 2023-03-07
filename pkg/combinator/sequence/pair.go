package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Pair Gets an object from the first parser, then gets another object from the second parser.
func Pair[R parser.Reader, F, S any](first parser.Parser[R, F], second parser.Parser[R, S]) parser.Parser[R, parser.Pair[F, S]] {
	return func(in R) (parser.Pair[F, S], error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		f, err := first(in)
		if err != nil {
			var r parser.Pair[F, S]
			return r, err
		}
		s, err := second(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			var r parser.Pair[F, S]
			return r, err
		}
		return parser.Pair[F, S]{First: f, Second: s}, err
	}
}

func SeparatedPair[R parser.Reader, F, S, T any](
	first parser.Parser[R, F], separator parser.Parser[R, T], second parser.Parser[R, S],
) parser.Parser[R, parser.Pair[F, S]] {
	return func(in R) (parser.Pair[F, S], error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		f, err := first(in)
		if err != nil {
			var r parser.Pair[F, S]
			return r, err
		}

		if _, err := separator(in); err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			var r parser.Pair[F, S]
			return r, err
		}

		s, err := second(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			var r parser.Pair[F, S]
			return r, err
		}
		return parser.Pair[F, S]{First: f, Second: s}, err
	}
}
