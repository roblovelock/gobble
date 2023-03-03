package main

import (
	"fmt"
	"gobble/pkg/combinator"
	"gobble/pkg/combinator/branch"
	"gobble/pkg/combinator/multi"
	"gobble/pkg/combinator/sequence"
	"gobble/pkg/parser"
	"gobble/pkg/parser/ascii"
	"gobble/pkg/parser/bytes"
	"gobble/pkg/parser/runes"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data := "\"abc\""
	fmt.Printf("EXAMPLE 1:\nParsing a simple input string:\n\n%s", data)

	result, _ := parseString(strings.NewReader(data))
	fmt.Printf(" => %s\n\n", result)

	data = "\"tab:\\tafter tab, newline:\\nnew line, quote: \\\", emoji: \\u{1F602}, newline:\\nescaped whitespace: \\    abc\""
	fmt.Printf("EXAMPLE 2:\nParsing a string with escape sequences, newline literal, and escaped whitespace:\n\n%s", data)

	result, _ = parseString(strings.NewReader(data))
	fmt.Printf(" => %s\n\n", result)
}

func parseEscapedChar() parser.Parser[parser.Reader, string] {
	return combinator.Map(
		sequence.Preceded(
			bytes.Byte('\\'),
			branch.Alt(
				parseUnicode(),
				combinator.Value(bytes.Byte('n'), '\n'),
				combinator.Value(bytes.Byte('r'), '\r'),
				combinator.Value(bytes.Byte('t'), '\t'),
				combinator.Value(bytes.Byte('b'), '\b'),
				combinator.Value(bytes.Byte('f'), '\f'),
				combinator.Value(bytes.Byte('\\'), '\\'),
				combinator.Value(bytes.Byte('/'), '/'),
				combinator.Value(bytes.Byte('"'), '"'),
			),
		),
		func(r rune) (string, error) { return string(r), nil },
	)
}

func parseUnicode() parser.Parser[parser.Reader, rune] {
	return combinator.Map(
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

func parseEscapedWhitespace() parser.Parser[parser.Reader, string] {
	return sequence.Preceded(
		sequence.Preceded(bytes.Byte('\\'), ascii.Whitespace1()),
		combinator.Success[parser.Reader](""),
	)
}

func parseLiteral() parser.Parser[parser.Reader, string] {
	return combinator.Verify(runes.TakeWhile(isNot('"', '\\')), func(v string) bool { return len(v) > 0 })
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

func parseFragment() parser.Parser[parser.Reader, string] {
	return branch.Alt(
		parseLiteral(),
		parseEscapedChar(),
		parseEscapedWhitespace(),
	)
}

func parseString(in parser.Reader) (string, error) {
	return combinator.Map(sequence.Delimited(
		bytes.Byte('"'),
		multi.FoldMany0(
			parseFragment(),
			strings.Builder{},
			func(builder strings.Builder, v string) strings.Builder {
				builder.WriteString(v)
				return builder
			}),
		bytes.Byte('"'),
	), func(builder strings.Builder) (string, error) {
		return builder.String(), nil
	})(in)
}
