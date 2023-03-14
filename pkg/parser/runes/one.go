// Package runes provides parsers for recognizing runes
package runes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"unicode/utf8"
)

type (
	oneParser struct{}
)

var oneParserInstance = &oneParser{}

func (o *oneParser) Parse(in parser.Reader) (rune, error) {
	r, _, err := in.ReadRune()
	return r, err
}

func (o *oneParser) ParseBytes(in []byte) (rune, []byte, error) {
	if len(in) == 0 {
		return 0, in, io.EOF
	}

	if c := in[0]; c < utf8.RuneSelf {
		return rune(c), in[1:], nil
	}

	ch, size := utf8.DecodeRune(in)
	return ch, in[size:], nil
}

// One reads a single rune
//
//   - If the input isn't empty, it will return a single rune.
//   - If the input is empty, it will return io.EOF
func One() parser.Parser[parser.Reader, rune] {
	return oneParserInstance
}
