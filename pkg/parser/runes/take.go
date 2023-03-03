package runes

import (
	"gobble/pkg/parser"
	"io"
	"strings"
)

// Take returns a string containing n runes from the input
//   - If the input contains n runes, it will return a string.
//   - If the input doesn't contain n runes, it will return io.EOF.
func Take(n int) parser.Parser[parser.Reader, string] {
	return func(in parser.Reader) (string, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		var builder strings.Builder
		builder.Grow(n)
		for i := 0; i < n; i++ {
			r, _, err := in.ReadRune()
			if err != nil {
				_, _ = in.Seek(currentOffset, io.SeekStart)
				return "", err
			}
			_, _ = builder.WriteRune(r)
		}

		return builder.String(), nil
	}
}