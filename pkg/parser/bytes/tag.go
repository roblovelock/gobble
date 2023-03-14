package bytes

import (
	"bytes"
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	tagParser struct {
		tag []byte
	}
)

func (o *tagParser) Parse(in parser.Reader) ([]byte, error) {
	result := make([]byte, len(o.tag))
	n, err := in.Read(result)
	if err != nil || n != len(o.tag) {
		_, _ = in.Seek(-int64(n), io.SeekCurrent)
		return nil, io.EOF
	}

	if bytes.Compare(o.tag, result) != 0 {
		_, _ = in.Seek(-int64(n), io.SeekCurrent)
		return nil, errors.ErrNotMatched
	}

	return result, nil
}

func (o *tagParser) ParseBytes(in []byte) ([]byte, []byte, error) {
	if len(in) < len(o.tag) {
		return nil, in, io.EOF
	}
	if bytes.Compare(o.tag, in[:len(o.tag)]) != 0 {
		return nil, in, errors.ErrNotMatched
	}

	return in[:len(o.tag)], in[len(o.tag):], nil
}

// Tag matches the argument
//   - If the input matches the argument, it will return the tag.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return errors.ErrNotMatched
func Tag(tag []byte) parser.Parser[parser.Reader, []byte] {
	return &tagParser{tag: tag}
}
