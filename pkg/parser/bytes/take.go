package bytes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	takeParser struct {
		n uint
	}
)

func (o *takeParser) Parse(in parser.Reader) ([]byte, error) {
	b := make([]byte, o.n)
	n, err := io.ReadFull(in, b)
	if err != nil {
		if err == io.ErrUnexpectedEOF {
			_, _ = in.Seek(-int64(n), io.SeekCurrent)
		}
		return nil, io.EOF
	}

	return b, nil
}

func (o *takeParser) ParseBytes(in []byte) ([]byte, []byte, error) {
	if len(in) < int(o.n) {
		return nil, in, io.EOF
	}

	return in[:o.n], in[o.n:], nil
}

// Take returns a slice of n bytes from the input
//   - If the input contains n bytes, it will return a slice of n bytes.
//   - If the input doesn't contain n bytes, it will return io.EOF
func Take(n uint) parser.Parser[parser.Reader, []byte] {
	return &takeParser{n: n}
}
