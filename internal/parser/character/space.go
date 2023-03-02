package character

import (
	"gobble/internal/combinator"
	"gobble/internal/parser"
)

func Space0() parser.Parser[parser.Reader, string] {
	return combinator.Map(IsA(' ', '\r', '\n', '\t'), func(s string) (string, error) { return s, nil })
}

func Space1() parser.Parser[parser.Reader, string] {
	return IsA(' ', '\r', '\n', '\t')
}
