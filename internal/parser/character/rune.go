package character

import (
	"gobble/internal/parser"
	"io"
)

func Rune(c rune) parser.Parser[parser.Reader, rune] {
	return func(in parser.Reader) (rune, error) {
		b, i, err := in.ReadRune()
		if err != nil {
			_, _ = in.Seek(-int64(i), io.SeekCurrent)
			return 0, err
		}

		if b != c {
			_, _ = in.Seek(-int64(i), io.SeekCurrent)
			return 0, parser.ErrNotMatched
		}

		return b, nil
	}
}
