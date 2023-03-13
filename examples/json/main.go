package main

import (
	"github.com/roblovelock/gobble/pkg/combinator/branch"
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/combinator/multi"
	"github.com/roblovelock/gobble/pkg/combinator/sequence"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"strconv"
	"strings"
)

var (
	numericCheck = [256]bool{}

	jsonValue    parser.Parser[parser.Reader, interface{}]
	jsonValuePtr = parser.Pointer(&jsonValue)

	ws         = ascii.SkipBlankSpace0()
	openSqrt   = bytes.Byte('[')
	closeSqrt  = bytes.Byte(']')
	openBrace  = bytes.Byte('{')
	closeBrace = bytes.Byte('}')
	comma      = bytes.Byte(',')
	colon      = bytes.Byte(':')
	quote      = bytes.Byte('"')

	untypedValue = modifier.Value[parser.Reader, []byte, interface{}]

	nullVal   = untypedValue(bytes.Tag([]byte("null")), nil)
	trueVal   = untypedValue(bytes.Tag([]byte("true")), true)
	falseVal  = untypedValue(bytes.Tag([]byte("false")), false)
	stringVal = sequence.Delimited(
		quote,
		bytes.Escaped(
			func(b byte) bool {
				return b != '"'
			},
			'\\',
			func(b byte) bool {
				switch b {
				case '"', 'n', '\\':
					return true
				default:
					return false
				}
			}),
		quote,
	)
	numericVal = modifier.Map(
		bytes.TakeWhile(func(b byte) bool { return numericCheck[b] }),
		func(b []byte) (interface{}, error) {
			return strconv.ParseFloat(string(b), 64)
		},
	)
	arrayVal = sequence.Delimited(
		openSqrt, multi.Separated0(jsonValuePtr, sequence.Preceded(ws, comma)), sequence.Preceded(ws, closeSqrt),
	)

	fieldName = sequence.Preceded(ws, stringVal)
	objVal    = sequence.Delimited(
		openBrace,
		multi.KeyValue(fieldName, sequence.Preceded(ws, colon), jsonValuePtr, sequence.Preceded(ws, comma)),
		sequence.Preceded(ws, closeBrace),
	)
)

func init() {
	parsers := map[byte]parser.Parser[parser.Reader, interface{}]{
		'"': parser.Untyped(stringVal),
		'[': parser.Untyped(arrayVal),
		'{': parser.Untyped(objVal),
		't': trueVal,
		'f': falseVal,
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

	jsonValue = sequence.Preceded(ws, branch.PeekCase(parsers))
}

func ParseJSON(json string) (interface{}, error) {
	reader := strings.NewReader(json)
	return jsonValue.Parse(reader)
}
