package ascii

import (
	"gobble/pkg/parser"
	"gobble/pkg/parser/bytes"
)

func Whitespace() parser.Parser[parser.Reader, byte] {
	return bytes.OneOf(' ', '\r', '\n', '\t')
}

func Whitespace0() parser.Parser[parser.Reader, []byte] {
	return bytes.OneOf0(' ', '\r', '\n', '\t')
}

func Whitespace1() parser.Parser[parser.Reader, []byte] {
	return bytes.OneOf1(' ', '\r', '\n', '\t')
}
