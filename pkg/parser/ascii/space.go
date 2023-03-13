package ascii

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
)

// Space returns a single ASCII space character: [ \t]
func Space() parser.Parser[parser.Reader, byte] {
	return bytes.OneOf(' ', '\t')
}

// Space0 returns zero or more ASCII space characters: [ \t]
func Space0() parser.Parser[parser.Reader, []byte] {
	return bytes.TakeWhile(IsSpace)
}

// Space1 returns one or more ASCII space characters: [ \t]
func Space1() parser.Parser[parser.Reader, []byte] {
	return bytes.TakeWhile1(IsSpace)
}

// SkipSpace skips a single ASCII space character: [ \t]
func SkipSpace() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.Skip(IsSpace)
}

// SkipSpace0 skips zero or more ASCII space characters: [ \t]
func SkipSpace0() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.SkipWhile(IsSpace)
}

// SkipSpace1 skips one or more ASCII space characters: [ \t]
func SkipSpace1() parser.Parser[parser.Reader, parser.Empty] {
	return bytes.SkipWhile1(IsSpace)
}
