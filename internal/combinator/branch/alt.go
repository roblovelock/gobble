package branch

import "gobble/internal/parser"

func Alt[R parser.Reader, T any](parsers ...parser.Parser[R, T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		for _, p := range parsers {
			if r, err := p(in); err == nil {
				return r, nil
			}
		}
		var r T
		return r, parser.ErrNotMatched
	}
}
