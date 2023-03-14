package bytes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/utils"
	"io"
)

type (
	oneOfParser struct {
		bytes [256]bool
	}

	oneOf0Parser struct {
		bytes [256]bool
	}

	oneOf1Parser struct {
		bytes [256]bool
	}
)

func (o *oneOfParser) Parse(in parser.Reader) (byte, error) {
	b, err := in.ReadByte()
	if err != nil {
		return 0, err
	}
	if o.bytes[b] {
		return b, nil
	}
	_, _ = in.Seek(-1, io.SeekCurrent)
	return 0, errors.ErrNotMatched
}

func (o *oneOfParser) ParseBytes(in []byte) (byte, []byte, error) {
	if len(in) == 0 {
		return 0, in, io.EOF
	}
	if !o.bytes[in[0]] {
		return 0, in, errors.ErrNotMatched
	}

	return in[0], in[1:], nil
}

func (o *oneOf0Parser) Parse(in parser.Reader) ([]byte, error) {
	result := make([]byte, 0)
	for {
		b, err := in.ReadByte()
		if err != nil {
			break
		}
		if !o.bytes[b] {
			_, _ = in.Seek(-1, io.SeekCurrent)
			break
		}
		result = append(result, b)
	}
	return result, nil
}

func (o *oneOf0Parser) ParseBytes(in []byte) ([]byte, []byte, error) {
	n := 0
	for ; n < len(in); n++ {
		if !o.bytes[in[n]] {
			break
		}
	}

	return in[:n], in[n:], nil
}

func (o *oneOf1Parser) Parse(in parser.Reader) ([]byte, error) {
	result := make([]byte, 0, 1)
	for {
		b, err := in.ReadByte()
		if err != nil {
			if len(result) == 0 {
				return nil, err
			}
			break
		}
		if !o.bytes[b] {
			_, _ = in.Seek(-1, io.SeekCurrent)
			if len(result) == 0 {
				return nil, errors.ErrNotMatched
			}
			break
		}
		result = append(result, b)
	}

	return result, nil
}

func (o *oneOf1Parser) ParseBytes(in []byte) ([]byte, []byte, error) {
	if len(in) == 0 {
		return nil, in, io.EOF
	}
	n := 0
	for ; n < len(in); n++ {
		if !o.bytes[in[n]] {
			break
		}
	}
	if n < 1 {
		return nil, in, errors.ErrNotMatched
	}

	return in[:n], in[n:], nil
}

// OneOf matches one of the argument bytes
//   - If the input matches the argument, it will return a single matched byte.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return errors.ErrNotMatched
func OneOf(bytes ...byte) parser.Parser[parser.Reader, byte] {
	return &oneOfParser{bytes: utils.NewByteLookupArray(bytes)}
}

// NotOneOf matches any byte not matching the argument bytes
//   - If the input doesn't match the argument, it will return a single byte.
//   - If the input is empty, it will return io.EOF
//   - If the input matches the argument, it will return errors.ErrNotMatched
func NotOneOf(bytes ...byte) parser.Parser[parser.Reader, byte] {
	return &oneOfParser{bytes: utils.NewInverseByteLookupArray(bytes)}
}

// OneOf0 matches zero or more bytes matching one of the argument bytes
//   - If the input matches the argument, it will return a slice of all matched bytes.
//   - If the input is empty, it will return an empty slice.
//   - If the input doesn't match, it will return an empty slice.
func OneOf0(bytes ...byte) parser.Parser[parser.Reader, []byte] {
	return &oneOf0Parser{bytes: utils.NewByteLookupArray(bytes)}
}

// OneOf1 matches one or more bytes matching one of the argument bytes
//   - If the input matches the argument, it will return a slice of all matched bytes.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return errors.ErrNotMatched
func OneOf1(bytes ...byte) parser.Parser[parser.Reader, []byte] {
	return &oneOf1Parser{bytes: utils.NewByteLookupArray(bytes)}
}
