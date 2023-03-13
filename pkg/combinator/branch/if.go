package branch

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	ifParser[R parser.Reader, C, T any] struct {
		condition parser.Parser[R, C]
		success   parser.Parser[R, T]
		err       parser.Parser[R, T]
	}
)

func (o *ifParser[R, C, T]) Parse(in R) (result T, err error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	if _, err = o.condition.Parse(in); err != nil {
		result, err = o.err.Parse(in)
	} else {
		result, err = o.success.Parse(in)
	}

	if err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
	}

	return
}

// If runs the conditional parser and chooses whether to run the success or error parser based on the outcome.
func If[R parser.Reader, C, T any](
	condition parser.Parser[R, C], success parser.Parser[R, T], err parser.Parser[R, T],
) parser.Parser[R, T] {
	return &ifParser[R, C, T]{condition: condition, success: success, err: err}
}
