// Package runes provides parsers for recognizing runes
package runes

import (
	"github.com/roblovelock/gobble/pkg/parser"
)

// One reads a single rune
//
//   - If the input isn't empty, it will return a single rune.
//   - If the input is empty, it will return io.EOF
func One() parser.Parser[parser.Reader, rune] {
	return func(in parser.Reader) (rune, error) {
		r, _, err := in.ReadRune()
		return r, err
	}
}
