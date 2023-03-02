package sequence

import (
	"gobble/internal/parser"
	"io"
)

// Delimited Matches an object from the first parser and discards it, then gets an object from the second parser,
// and finally matches an object from the third parser and discards it.
func Delimited[R parser.Reader, F, S, T any](
	first parser.Parser[R, F], second parser.Parser[R, S], third parser.Parser[R, T],
) parser.Parser[R, S] {
	return func(in R) (S, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		if _, err := first(in); err != nil {
			var r S
			return r, err
		}

		s, err := second(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return s, err
		}

		if _, err := third(in); err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			var r S
			return r, err
		}

		return s, nil
	}
}
