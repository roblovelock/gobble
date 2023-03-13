package stream

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	offsetParser struct{}
)

var offsetParserInstance = &offsetParser{}

func (o *offsetParser) Parse(in parser.Reader) (int64, error) {
	return in.Seek(0, io.SeekCurrent)
}

// Offset Returns the current stream offset. It doesn't consume any input.
func Offset() parser.Parser[parser.Reader, int64] {
	return offsetParserInstance
}
