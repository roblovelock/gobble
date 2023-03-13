// Package expr provides a parser to parse basic arithmetic expression based on the
// following rule.
//
//	expr  -> sum
//	prod  -> value (mulop value)*
//	mulop -> "*"
//	      |  "/"
//	sum   -> prod (addop prod)*
//	addop -> "+"
//	      |  "-"
//	value -> num
//	      | "(" expr ")"
package main

import (
	"github.com/roblovelock/gobble/pkg/combinator/branch"
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/combinator/multi"
	"github.com/roblovelock/gobble/pkg/combinator/sequence"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"strings"
)

var (
	openParen  = sequence.Terminated(bytes.Byte('('), ascii.SkipWhitespace0())
	closeParen = sequence.Preceded(ascii.SkipWhitespace0(), bytes.Byte(')'))
	add        = bytes.Byte('+')
	subtract   = bytes.Byte('-')
	multiply   = bytes.Byte('*')
	divide     = bytes.Byte('/')

	// addop -> "+" |  "-"
	addOp = branch.Alt(add, subtract)
	// mulop -> "*" |  "/"
	mulOp = branch.Alt(multiply, divide)

	foldOp = func(p parser.Pair[int64, []func(int64) int64]) (int64, error) {
		val := p.First
		for _, op := range p.Second {
			val = op(val)
		}
		return val, nil
	}

	sum parser.Parser[parser.Reader, int64]
)

func init() {

	// num | "(" expr ")"
	value := sequence.Delimited(
		ascii.SkipWhitespace0(),
		branch.Alt(ascii.Int64(), sequence.Delimited(openParen, parser.Pointer(&sum), closeParen)),
		ascii.SkipWhitespace0(),
	)

	// (mulop value)*
	mulOpValue := multi.Many0(modifier.Map(
		sequence.Pair(mulOp, value),
		func(p parser.Pair[byte, int64]) (func(int64) int64, error) {
			if p.First == '*' {
				return func(v int64) int64 {
					return v * p.Second
				}, nil
			}
			return func(v int64) int64 {
				return v / p.Second
			}, nil
		},
	))

	prod := modifier.Map(sequence.Pair(value, mulOpValue), foldOp)

	// (addop value)*
	addOpValue := multi.Many0(modifier.Map(
		sequence.Pair(addOp, prod),
		func(p parser.Pair[byte, int64]) (func(int64) int64, error) {
			if p.First == '+' {
				return func(v int64) int64 {
					return v + p.Second
				}, nil
			}
			return func(v int64) int64 {
				return v - p.Second
			}, nil
		},
	))

	sum = modifier.Map(sequence.Pair(prod, addOpValue), foldOp)
}

func ParseExpr(expr string) (int64, error) {
	reader := strings.NewReader(expr)
	return sum.Parse(reader)
}
