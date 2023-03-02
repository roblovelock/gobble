package character

import (
	"gobble/internal/parser"
	"strings"
)

func Take(n int) parser.Parser[parser.Reader, string] {
	return func(in parser.Reader) (string, error) {
		var builder strings.Builder
		builder.Grow(n)
		for i := 0; i < n; i++ {
			r, _, err := in.ReadRune()
			if err != nil {
				return "", err
			}
			_, _ = builder.WriteRune(r)
		}

		return builder.String(), nil
	}
}
