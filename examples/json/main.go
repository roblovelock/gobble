package main

import (
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
)

var (
	numericCheck = [256]bool{}

	ws         = ascii.SkipWhitespace0()
	openSqrt   = bytes.Byte('[')
	closeSqrt  = bytes.Byte(']')
	openBrace  = bytes.Byte('{')
	closeBrace = bytes.Byte('}')
	comma      = bytes.Byte(',')
	colon      = bytes.Byte(':')

	stringVal  = runes.EscapedString()
	numericVal = modifier.Map(
		bytes.TakeWhile(func(b byte) bool { return numericCheck[b] }),
		func(b []byte) (interface{}, error) {
			return strconv.ParseFloat(string(b), 64)
		},
	)

	nullVal  = modifier.Value[parser.Reader, []byte, interface{}](bytes.Tag([]byte("null")), nil)
	trueVal  = modifier.Value(bytes.Tag([]byte("true")), true)
	falseVal = modifier.Value(bytes.Tag([]byte("false")), false)

	fieldName = sequence.Terminated(stringVal, ws)

	val parser.Parser[parser.Reader, interface{}]
)

func init() {
	var arrayVal parser.Parser[parser.Reader, []interface{}]
	var objVal parser.Parser[parser.Reader, map[string]interface{}]
	parsers := map[byte]parser.Parser[parser.Reader, interface{}]{
		'"': parser.Untyped(stringVal),
		't': parser.Untyped(trueVal),
		'f': parser.Untyped(falseVal),
		'[': parser.Untyped(parser.Ptr(&arrayVal)),
		'{': parser.Untyped(parser.Ptr(&objVal)),
		'n': nullVal,
		'-': numericVal,
		'+': numericVal,
		'.': numericVal,
	}

	for i := byte('0'); i <= '9'; i++ {
		numericCheck[i] = true
		parsers[i] = numericVal
	}
	numericCheck['-'] = true
	numericCheck['+'] = true
	numericCheck['.'] = true
	numericCheck['e'] = true
	numericCheck['E'] = true

	val = sequence.Delimited(ws, branch.PeekCase(parsers), ws)

	arrayVal = sequence.Delimited(
		openSqrt, multi.Separated0(val, comma), closeSqrt,
	)

	field := sequence.SeparatedPair(fieldName, colon, val)

	objVal = modifier.Map(sequence.Delimited(
		sequence.Terminated(openBrace, ws),
		multi.Separated0(field, sequence.Delimited(ws, comma, ws)),
		sequence.Preceded(ws, closeBrace),
	), func(pairs []parser.Pair[string, interface{}]) (map[string]interface{}, error) {
		m := make(map[string]interface{}, len(pairs))
		for _, p := range pairs {
			m[p.First] = p.Second
		}
		return m, nil
	},
	)
}

func ParseJSON(json string) (interface{}, error) {
	reader := strings.NewReader(json)
	return val(reader)
}
