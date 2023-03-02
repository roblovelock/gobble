package sequence

import (
	"gobble/internal/parser"
	"io"
)

// Preceded Matches an object from the first parser and discards it, then gets an object from the second parser.
func Preceded[R parser.Reader, F, S any](first parser.Parser[R, F], second parser.Parser[R, S]) parser.Parser[R, S] {
	return func(in R) (S, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		if _, err := first(in); err != nil {
			var r S
			return r, err
		}
		s, err := second(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
		}
		return s, err
	}
}
