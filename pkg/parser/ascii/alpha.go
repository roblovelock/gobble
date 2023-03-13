// Package ascii provides parsers for recognizing ascii bytes
package ascii

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	alphaParser struct {
	}
	alpha0Parser struct {
	}
	alpha1Parser struct {
	}
)

func (o *alphaParser) Parse(in parser.Reader) (byte, error) {
	b, err := in.ReadByte()
	if err != nil {
		return 0, err
	}

	if IsLetter(b) {
		return b, nil
	}

	_, _ = in.Seek(-1, io.SeekCurrent)
	return 0, errors.ErrNotMatched
}

func (o *alpha0Parser) Parse(in parser.Reader) ([]byte, error) {
	chars := make([]byte, 0)

	for {
		b, err := in.ReadByte()
		if err != nil {
			return chars, nil
		} else if !IsLetter(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return chars, nil
		}
		chars = append(chars, b)
	}
}

func (o *alpha1Parser) Parse(in parser.Reader) ([]byte, error) {
	b, err := in.ReadByte()
	if err != nil {
		return nil, err
	}

	if !IsLetter(b) {
		_, _ = in.Seek(-1, io.SeekCurrent)
		return nil, errors.ErrNotMatched
	}

	chars := []byte{b}

	for {
		b, err := in.ReadByte()
		if err != nil {
			return chars, nil
		} else if !IsLetter(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return chars, nil
		}

		chars = append(chars, b)
	}
}

var alphaParserInstance = &alphaParser{}
var alpha0ParserInstance = &alpha0Parser{}
var alpha1ParserInstance = &alpha1Parser{}

// Alpha matches a single ASCII letter character: [a-zA-Z]
//   - If the input matches a letter character, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match a letter character, it will return parser.ErrNotMatched
func Alpha() parser.Parser[parser.Reader, byte] {
	return alphaParserInstance
}

// Alpha0 matches zero or more ASCII letter characters: [a-zA-Z]
//   - If the input matches a letter character, it will return a slice of all matched letters.
//   - If the input is empty, it will return an empty slice.
//   - If the input doesn't match a letter character, it will return an empty slice.
func Alpha0() parser.Parser[parser.Reader, []byte] {
	return alpha0ParserInstance
}

// Alpha1 matches one or more ASCII letter characters: [a-zA-Z]
//   - If the input matches a letter character, it will return a slice of all matched letters.
//   - If the input is empty, it will return io.EOF.
//   - If the input doesn't match a letter character, it will return parser.ErrNotMatched.
func Alpha1() parser.Parser[parser.Reader, []byte] {
	return alpha1ParserInstance
}
