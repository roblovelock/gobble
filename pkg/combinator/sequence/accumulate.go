package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Accumulate Matches a sequence of parsers and accumulates the results.
func Accumulate[R parser.Reader, A, T any](
	acc A, f parser.Accumulator[T, A], parsers ...parser.Parser[R, T],
) parser.Parser[R, A] {
	return func(in R) (A, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		a := acc
		for _, p := range parsers {
			r, err := p(in)
			if err != nil {
				var a A
				_, _ = in.Seek(currentOffset, io.SeekStart)
				return a, err
			}
			a = f(a, r)
		}

		return a, nil
	}
}

// AccumulateBytes Matches a sequence of byte slice parsers and accumulates the results.
func AccumulateBytes[R parser.Reader](parsers ...parser.Parser[R, []byte]) parser.Parser[R, []byte] {
	return Accumulate(
		[]byte{},
		func(a []byte, t []byte) []byte {
			return append(a, t...)
		},
		parsers...,
	)
}
