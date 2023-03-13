// Package runes provides parsers for recognizing runes
package runes

import (
	"github.com/roblovelock/gobble/pkg/parser"
)

type (
	oneParser struct{}
)

var oneParserInstance = &oneParser{}

func (o *oneParser) Parse(in parser.Reader) (rune, error) {
	r, _, err := in.ReadRune()
	return r, err
}

// One reads a single rune
//
//   - If the input isn't empty, it will return a single rune.
//   - If the input is empty, it will return io.EOF
func One() parser.Parser[parser.Reader, rune] {
	return oneParserInstance
}
