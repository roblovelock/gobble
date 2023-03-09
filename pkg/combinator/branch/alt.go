package branch

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
)

// Alt Trys a list of parsers and returns the result of the first successful one.
func Alt[R parser.Reader, T any](parsers ...parser.Parser[R, T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		for _, p := range parsers {
			if r, err := p(in); err == nil {
				return r, nil
			} else if errors.IsFatal(err) {
				var t T
				return t, err
			}
		}
		var r T
		return r, errors.ErrNotMatched
	}
}
