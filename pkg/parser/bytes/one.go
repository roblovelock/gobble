// Package bytes provides parsers for recognizing bytes
package bytes

import (
	"gobble/pkg/parser"
)

// One reads a single byte
//
//   - If the input isn't empty, it will return a single byte.
//   - If the input is empty, it will return io.EOF
func One() parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		return in.ReadByte()
	}
}
