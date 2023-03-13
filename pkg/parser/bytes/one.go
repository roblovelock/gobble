// Package bytes provides parsers for recognizing bytes
package bytes

import (
	"github.com/roblovelock/gobble/pkg/parser"
)

type (
	oneParser struct {
	}
)

var oneParserInstance = &oneParser{}

func (o *oneParser) Parse(in parser.Reader) (byte, error) {
	return in.ReadByte()
}

// One reads a single byte
//
//   - If the input isn't empty, it will return a single byte.
//   - If the input is empty, it will return io.EOF
func One() parser.Parser[parser.Reader, byte] {
	return oneParserInstance
}
