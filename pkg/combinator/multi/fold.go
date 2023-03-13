package multi

import "github.com/roblovelock/gobble/pkg/parser"

type (
	foldMany0Parser[R parser.Reader, T, A any] struct {
		parser      parser.Parser[R, T]
		accumulator A
		fn          parser.Accumulator[T, A]
	}
)

func (o *foldMany0Parser[R, T, A]) Parse(in R) (A, error) {
	for r, err := o.parser.Parse(in); err == nil; r, err = o.parser.Parse(in) {
		o.accumulator = o.fn(o.accumulator, r)
	}
	return o.accumulator, nil
}

func FoldMany0[R parser.Reader, T, A any](p parser.Parser[R, T], acc A, f parser.Accumulator[T, A]) parser.Parser[R, A] {
	return &foldMany0Parser[R, T, A]{parser: p, accumulator: acc, fn: f}
}
