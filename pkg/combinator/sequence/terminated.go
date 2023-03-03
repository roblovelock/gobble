package sequence

import (
	"gobble/pkg/parser"
	"io"
)

// Terminated Gets an object from the first parser, then matches an object from the second parser and discards it.
func Terminated[R parser.Reader, F, S any](first parser.Parser[R, F], second parser.Parser[R, S]) parser.Parser[R, F] {
	return func(in R) (F, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		f, err := first(in)
		if err != nil {
			return f, err
		}

		if _, err := second(in); err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			var r F
			return r, err
		}
		return f, err
	}
}
