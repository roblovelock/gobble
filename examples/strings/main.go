package main

import (
	"github.com/roblovelock/gobble/pkg/combinator"
	"github.com/roblovelock/gobble/pkg/combinator/branch"
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/combinator/multi"
	"github.com/roblovelock/gobble/pkg/combinator/sequence"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"strconv"
	"strings"
	"unicode"
)

func escapedCharParser() parser.Parser[parser.Reader, string] {
	return modifier.Map(
		sequence.Preceded(
			bytes.Byte('\\'),
			branch.Alt(
				unicodeParser(),
				modifier.Value(bytes.Byte('n'), '\n'),
				modifier.Value(bytes.Byte('r'), '\r'),
				modifier.Value(bytes.Byte('t'), '\t'),
				modifier.Value(bytes.Byte('b'), '\b'),
				modifier.Value(bytes.Byte('f'), '\f'),
				modifier.Value(bytes.Byte('\\'), '\\'),
				modifier.Value(bytes.Byte('/'), '/'),
				modifier.Value(bytes.Byte('"'), '"'),
			),
		),
		func(r rune) (string, error) { return string(r), nil },
	)
}

func unicodeParser() parser.Parser[parser.Reader, rune] {
	return modifier.Map(
		sequence.Preceded(
			bytes.Byte('u'),
			sequence.Delimited(
				bytes.Byte('{'),
				runes.TakeWhileMinMax(1, 6, func(r rune) bool { return unicode.Is(unicode.ASCII_Hex_Digit, r) }),
				bytes.Byte('}'),
			),
		),
		func(hex string) (rune, error) {
			i, err := strconv.ParseInt(hex, 16, 32)
			return rune(i), err
		},
	)
}

func escapedWhitespaceParser() parser.Parser[parser.Reader, string] {
	return sequence.Preceded(
		sequence.Preceded(bytes.Byte('\\'), ascii.Whitespace1()),
		combinator.Success[parser.Reader](""),
	)
}

func literalParser() parser.Parser[parser.Reader, string] {
	return modifier.Verify(runes.TakeWhile(isNot('"', '\\')), func(v string) bool { return len(v) > 0 })
}

func isNot(runes ...rune) parser.Predicate[rune] {
	return func(r1 rune) bool {
		for _, r2 := range runes {
			if r1 == r2 {
				return false
			}
		}
		return true
	}
}

func fragmentParser() parser.Parser[parser.Reader, string] {
	return branch.Alt(
		literalParser(),
		escapedCharParser(),
		escapedWhitespaceParser(),
	)
}

func stringParser() parser.Parser[parser.Reader, string] {
	return modifier.Map(sequence.Delimited(
		bytes.Byte('"'),
		multi.FoldMany0(
			fragmentParser(),
			strings.Builder{},
			func(builder strings.Builder, v string) strings.Builder {
				builder.WriteString(v)
				return builder
			}),
		bytes.Byte('"'),
	), func(builder strings.Builder) (string, error) {
		return builder.String(), nil
	})
}

func parseString(in parser.Reader) (string, error) {
	return stringParser().Parse(in)
}
