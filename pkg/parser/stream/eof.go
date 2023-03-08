package stream

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// EOF Returns successfully if it is at the end of input data
func EOF() parser.Parser[parser.Reader, parser.Empty] {
	return func(in parser.Reader) (parser.Empty, error) {
		_, err := in.ReadByte()
		if err == io.EOF {
			return nil, nil
		}
		_, _ = in.Seek(-1, io.SeekCurrent)
		return nil, parser.ErrNotMatched
	}
}
