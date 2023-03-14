package multi

import (
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/parser"
)

type (
	separated0Parser[R parser.Reader, T any, S any] struct {
		parser    parser.Parser[R, T]
		separator parser.Parser[R, S]
	}
)

func (o *separated0Parser[R, T, S]) Parse(in R) ([]T, error) {
	result := make([]T, 0)
	for r, err := o.parser.Parse(in); err == nil; r, err = o.parser.Parse(in) {
		result = append(result, r)
		if _, err := o.separator.Parse(in); err != nil {
			break
		}
	}

	return result, nil
}

func (o *separated0Parser[R, T, S]) ParseBytes(in []byte) ([]T, []byte, error) {
	result := make([]T, 0)
	for {
		r, out, err := o.parser.ParseBytes(in)
		if err != nil {
			return result, in, nil
		}
		result = append(result, r)

		_, out, err = o.separator.ParseBytes(out)
		if err != nil {
			return result, out, nil
		}
		in = out
	}
}

func Separated0[R parser.Reader, T any, S any](
	p parser.Parser[R, T], separator parser.Parser[R, S],
) parser.Parser[R, []T] {
	return &separated0Parser[R, T, S]{parser: p, separator: separator}
}

func Separated0Count[R parser.Reader, T any, S any](
	p parser.Parser[R, T], separator parser.Parser[R, S],
) parser.Parser[R, int] {
	return modifier.Count[R, T, []T](Separated0(p, separator))
}

func Separated1[R parser.Reader, T any, S any](
	p parser.Parser[R, T], separator parser.Parser[R, S],
) parser.Parser[R, []T] {
	return modifier.Verify(Separated0(p, separator), func(ts []T) bool {
		return len(ts) > 0
	})
}

func Separated1Count[R parser.Reader, T any, S any](
	p parser.Parser[R, T], separator parser.Parser[R, S],
) parser.Parser[R, int] {
	return modifier.Count[R, T, []T](Separated1(p, separator))
}
