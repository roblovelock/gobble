package stream

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Offset Returns the current stream offset. It doesn't consume any input.
func Offset() parser.Parser[parser.Reader, int64] {
	return func(in parser.Reader) (int64, error) {
		return in.Seek(0, io.SeekCurrent)
	}
}
