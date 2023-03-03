package bytes

import (
	"gobble/pkg/parser"
	"io"
)

// Take returns a slice of n bytes from the input
//   - If the input contains n bytes, it will return a slice of n bytes.
//   - If the input doesn't contain n bytes, it will return io.EOF
func Take(n uint) parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		b := make([]byte, n)
		n, err := io.ReadFull(in, b)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				_, _ = in.Seek(-int64(n), io.SeekCurrent)
			}
			return nil, io.EOF
		}

		return b, nil
	}
}
