package character

import (
	"gobble/internal/parser"
	"io"
	"strings"
)

func TakeWhile(p parser.Predicate[rune]) parser.Parser[parser.Reader, string] {
	return func(in parser.Reader) (string, error) {
		builder := strings.Builder{}
		for {
			r, i, err := in.ReadRune()
			if err != nil || !p(r) {
				_, _ = in.Seek(-int64(i), io.SeekCurrent)
				break
			}
			builder.WriteRune(r)
		}

		return builder.String(), nil
	}
}

func TakeWhileMN(m, n int, p parser.Predicate[rune]) parser.Parser[parser.Reader, string] {
	return func(in parser.Reader) (string, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		builder := strings.Builder{}
		builder.Grow(m)

		for i := 0; i < n; i++ {
			r, x, err := in.ReadRune()
			if err == nil && !p(r) {
				_, _ = in.Seek(-int64(x), io.SeekCurrent)
				err = parser.ErrNotMatched
			}
			if err != nil {
				if builder.Len() < m {
					_, _ = in.Seek(currentOffset, io.SeekStart)
					return "", err
				}
				break
			}
			builder.WriteRune(r)
		}

		return builder.String(), nil
	}
}
