package stream

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	eofParser struct{}
)

var eofParserInstance = &eofParser{}

func (o *eofParser) Parse(in parser.Reader) (parser.Empty, error) {
	_, err := in.ReadByte()
	if err == io.EOF {
		return nil, nil
	}
	_, _ = in.Seek(-1, io.SeekCurrent)
	return nil, errors.ErrNotMatched
}

func (o *eofParser) ParseBytes(in []byte) (parser.Empty, []byte, error) {
	if len(in) == 0 {
		return nil, in, nil
	}
	return nil, in, errors.ErrNotMatched
}

// EOF Returns successfully if it is at the end of input data
func EOF() parser.Parser[parser.Reader, parser.Empty] {
	return eofParserInstance
}
