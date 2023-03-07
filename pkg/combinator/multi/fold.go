package multi

import "github.com/roblovelock/gobble/pkg/parser"

func FoldMany0[R parser.Reader, T, A any](p parser.Parser[R, T], acc A, f parser.Accumulator[T, A]) parser.Parser[R, A] {
	return func(in R) (A, error) {
		for r, err := p(in); err == nil; r, err = p(in) {
			acc = f(acc, r)
		}
		return acc, nil
	}
}
