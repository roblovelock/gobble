package branch

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
)

type (
	altParser[R parser.Reader, T any] struct {
		parsers []parser.Parser[R, T]
	}
)

func (o *altParser[R, T]) Parse(in R) (T, error) {
	for _, p := range o.parsers {
		if r, err := p.Parse(in); err == nil {
			return r, nil
		} else if errors.IsFatal(err) {
			var t T
			return t, err
		}
	}
	var r T
	return r, errors.ErrNotMatched
}

// Alt Trys a list of parsers and returns the result of the first successful one.
func Alt[R parser.Reader, T any](parsers ...parser.Parser[R, T]) parser.Parser[R, T] {
	return &altParser[R, T]{parsers: parsers}
}
