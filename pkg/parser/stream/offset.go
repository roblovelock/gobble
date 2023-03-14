package stream

import (
	"github.com/roblovelock/gobble/pkg/errors"
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

func (o *offsetParser) ParseBytes(in []byte) (int64, []byte, error) {
	return 0, in, errors.ErrNotSupported
}

// Offset Returns the current stream offset. It doesn't consume any input.
func Offset() parser.Parser[parser.Reader, int64] {
	return offsetParserInstance
}
