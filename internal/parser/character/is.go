package character

import (
	"gobble/internal/parser"
	"strings"
)

func IsA(runes ...rune) parser.Parser[parser.Reader, string] {
	return func(in parser.Reader) (string, error) {
		var builder strings.Builder
		for {
			r, err := OneOf(runes...)(in)
			if err != nil {
				if builder.Len() == 0 {
					return "", err
				}
				break
			}
			_, _ = builder.WriteRune(r)
		}
		return builder.String(), nil
	}
}

func IsNot(runes ...rune) parser.Parser[parser.Reader, string] {
	return TakeWhile(isNot(runes...))
}

func isNot(runes ...rune) parser.Predicate[rune] {
	return func(r rune) bool {
		for _, v := range runes {
			if r == v {
				return false
			}
		}
		return true
	}
}
