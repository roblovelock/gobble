package combinator

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

func EOF() parser.Parser[parser.Reader, parser.Empty] {
	return func(in parser.Reader) (parser.Empty, error) {
		_, err := in.ReadByte()
		if err == io.EOF {
			return parser.Empty{}, nil
		}
		_, _ = in.Seek(-1, io.SeekCurrent)
		return parser.Empty{}, parser.ErrNotMatched
	}
}
