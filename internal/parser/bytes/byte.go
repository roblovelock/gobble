package bytes

import (
	"gobble/internal/parser"
	"io"
)

func Byte(c byte) parser.Parser[parser.Reader, byte] {
	return func(in parser.Reader) (byte, error) {
		b, err := in.ReadByte()
		if err != nil {
			return 0, err
		}

		if b != c {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return 0, parser.ErrNotMatched
		}

		return b, nil
	}
}
