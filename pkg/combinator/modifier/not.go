package modifier

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Not returns a result only if the parser returns an error. It doesn't consume any input
func Not[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, parser.Empty] {
	return func(in R) (parser.Empty, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		_, err := p(in)
		_, _ = in.Seek(currentOffset, io.SeekStart)
		if err != nil {
			return nil, nil
		}
		return nil, parser.ErrNotMatched
	}
}
