// Package ascii provides parsers for recognizing ascii bytes
package ascii

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	alphanumericParser struct {
	}
	alphanumeric0Parser struct {
	}
	alphanumeric1Parser struct {
	}
)

func (o *alphanumericParser) Parse(in parser.Reader) (byte, error) {
	b, err := in.ReadByte()
	if err != nil {
		return 0, err
	}

	if IsAlphanumeric(b) {
		return b, nil
	}

	_, _ = in.Seek(-1, io.SeekCurrent)
	return 0, errors.ErrNotMatched
}

func (o *alphanumericParser) ParseBytes(in []byte) (byte, []byte, error) {
	if len(in) == 0 {
		return 0, in, io.EOF
	}

	if IsAlphanumeric(in[0]) {
		return in[0], in[1:], nil
	}
	return 0, in, errors.ErrNotMatched
}

func (o *alphanumeric0Parser) Parse(in parser.Reader) ([]byte, error) {
	chars := make([]byte, 0)

	for {
		b, err := in.ReadByte()
		if err != nil {
			return chars, nil
		} else if !IsAlphanumeric(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return chars, nil
		}
		chars = append(chars, b)
	}
}
func (o *alphanumeric0Parser) ParseBytes(in []byte) ([]byte, []byte, error) {
	for i := 0; i < len(in); i++ {
		if !IsAlphanumeric(in[i]) {
			return in[0:i], in[i:], nil
		}
	}
	return nil, in, nil
}

func (o *alphanumeric1Parser) Parse(in parser.Reader) ([]byte, error) {
	b, err := in.ReadByte()
	if err != nil {
		return nil, err
	}

	if !IsAlphanumeric(b) {
		_, _ = in.Seek(-1, io.SeekCurrent)
		return nil, errors.ErrNotMatched
	}

	chars := []byte{b}

	for {
		b, err := in.ReadByte()
		if err != nil {
			return chars, nil
		} else if !IsAlphanumeric(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return chars, nil
		}

		chars = append(chars, b)
	}
}

func (o *alphanumeric1Parser) ParseBytes(in []byte) ([]byte, []byte, error) {
	i := 0
	for ; i < len(in); i++ {
		if !IsAlphanumeric(in[i]) {
			break
		}
	}
	if i == 0 {
		return nil, in, errors.ErrNotMatched
	}
	return in[0:i], in[i:], nil
}

var alphanumericParserInstance = &alphanumericParser{}
var alphanumeric0ParserInstance = &alphanumeric0Parser{}
var alphanumeric1ParserInstance = &alphanumeric1Parser{}

// Alphanumeric matches a single ASCII letter or digit character: [a-zA-Z0-9]
//   - If the input matches a letter or digit character, it will return the match.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match a letter character, it will return errors.ErrNotMatched
func Alphanumeric() parser.Parser[parser.Reader, byte] {
	return alphanumericParserInstance
}

// Alphanumeric0 matches zero or more ASCII letter or digit characters: [a-zA-Z0-9]
//   - If the input matches a letter or digit character, it will return a slice of all matched characters.
//   - If the input is empty, it will return an empty slice.
//   - If the input doesn't match a letter or digit character, it will return an empty slice.
func Alphanumeric0() parser.Parser[parser.Reader, []byte] {
	return alphanumeric0ParserInstance
}

// Alphanumeric1 matches one or more ASCII letter or digit characters: [a-zA-Z0-9]
//   - If the input matches a letter or digit character, it will return a slice of all matched characters.
//   - If the input is empty, it will return io.EOF.
//   - If the input doesn't match a letter or digit character, it will return errors.ErrNotMatched.
func Alphanumeric1() parser.Parser[parser.Reader, []byte] {
	return alphanumeric1ParserInstance
}
