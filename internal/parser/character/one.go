package character

import (
	"gobble/internal/parser"
	"io"
)

func OneOf(runes ...rune) parser.Parser[parser.Reader, rune] {
	return func(in parser.Reader) (rune, error) {
		r, i, err := in.ReadRune()
		if err != nil {
			if err == io.EOF {
				return 0, io.EOF
			}
			_, _ = in.Seek(-int64(i), io.SeekCurrent)
			return 0, err
		}

		for _, v := range runes {
			if r == v {
				return r, nil
			}
		}

		_, _ = in.Seek(-int64(i), io.SeekCurrent)
		return 0, parser.ErrNotMatched
	}
}
