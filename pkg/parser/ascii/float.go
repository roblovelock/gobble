package ascii

import (
	"github.com/roblovelock/gobble/pkg/combinator/branch"
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/combinator/sequence"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"strconv"
)

type (
	floatConstraint interface {
		float32 | float64
	}
)

// Float32 will parse a number in text form to float32
func Float32() parser.Parser[parser.Reader, float32] {
	return floatParser[float32](func(b []byte) (float32, error) {
		r, err := strconv.ParseFloat(string(b), 64)
		return float32(r), err
	})
}

// Float64 will parse a number in text form to float64
func Float64() parser.Parser[parser.Reader, float64] {
	return floatParser[float64](func(b []byte) (float64, error) {
		return strconv.ParseFloat(string(b), 64)
	})
}

func floatParser[T floatConstraint](
	f func(b []byte) (T, error),
) parser.Parser[parser.Reader, T] {
	var sign = bytes.OneOf('+', '-')
	var point = bytes.Byte('.')

	var float = sequence.Recognize(sequence.Delimited(
		modifier.Optional(sign),
		branch.Alt(
			sequence.Terminated(Digit1(), modifier.Optional(sequence.Preceded(point, Digit1()))),
			modifier.Cut(sequence.Preceded(point, Digit1())),
		),
		modifier.Optional(sequence.Delimited(
			bytes.OneOf('e', 'E'),
			modifier.Optional(sign),
			Digit1(),
		)),
	))

	return modifier.Map(float, f)
}
