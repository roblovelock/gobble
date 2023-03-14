// Package bytes provides parsers for recognizing bytes
package bytes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	oneParser struct {
	}
)

var oneParserInstance = &oneParser{}

func (o *oneParser) Parse(in parser.Reader) (byte, error) {
	return in.ReadByte()
}

func (o *oneParser) ParseBytes(in []byte) (byte, []byte, error) {
	if len(in) == 0 {
		return 0, in, io.EOF
	}
	return in[0], in[1:], nil
}

// One reads a single byte
//
//   - If the input isn't empty, it will return a single byte.
//   - If the input is empty, it will return io.EOF
func One() parser.Parser[parser.Reader, byte] {
	return oneParserInstance
}
