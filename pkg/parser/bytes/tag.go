package bytes

import (
	"bytes"
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Tag matches the argument
//   - If the input matches the argument, it will return the tag.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the argument, it will return parser.ErrNotMatched
func Tag(b []byte) parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		result := make([]byte, len(b))
		n, err := in.Read(result)
		if err != nil || n != len(b) {
			_, _ = in.Seek(-int64(n), io.SeekCurrent)
			return nil, io.EOF
		}

		if bytes.Compare(b, result) != 0 {
			_, _ = in.Seek(-int64(n), io.SeekCurrent)
			return nil, errors.ErrNotMatched
		}

		return result, nil
	}
}
