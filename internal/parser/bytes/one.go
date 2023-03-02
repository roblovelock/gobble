package bytes

import (
	"gobble/internal/parser"
	"io"
)

func OneOf(bytes ...byte) parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, io.EOF
		}

		for _, v := range bytes {
			if b == v {
				return b, nil
			}
		}

		_, _ = in.Seek(-1, io.SeekCurrent)
		return 0, parser.ErrNotMatched
	}
}
