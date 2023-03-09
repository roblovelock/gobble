package branch

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// If runs the conditional parser and chooses whether to run the success or error parser based on the outcome.
func If[R parser.Reader, C, T any](
	condition parser.Parser[R, C], success parser.Parser[R, T], errParser parser.Parser[R, T],
) parser.Parser[R, T] {
	return func(in R) (result T, err error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		if _, err = condition(in); err != nil {
			result, err = errParser(in)
		} else {
			result, err = success(in)
		}

		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
		}

		return
	}
}
