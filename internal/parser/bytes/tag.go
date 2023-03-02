package bytes

import (
	"bytes"
	"gobble/internal/parser"
	"io"
)

func Tag(b []byte) parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)

		result := make([]byte, len(b))
		n, err := in.Read(result)
		if err != nil || n != len(b) || bytes.Compare(b, result) != 0 {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return nil, parser.ErrNotMatched
		}

		return result, nil
	}
}
