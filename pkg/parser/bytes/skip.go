package bytes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"math"
)

type (
	skipWhileMinMaxParser struct {
		min       int
		max       int
		predicate parser.Predicate[byte]
	}

	skipParser struct {
		predicate parser.Predicate[byte]
	}
)

func (o *skipWhileMinMaxParser) Parse(in parser.Reader) (parser.Empty, error) {
	startOffset, _ := in.Seek(0, io.SeekCurrent)
	n := 0
	for ; n < o.max; n++ {
		b, err := in.ReadByte()
		if err != nil {
			if n < o.min {
				return nil, err
			}
			break
		}

		if !o.predicate(b) {
			_, _ = in.Seek(-1, io.SeekCurrent)
			if n < o.min {
				_, _ = in.Seek(startOffset, io.SeekStart)
				return nil, errors.ErrNotMatched
			}
			break
		}
	}

	return nil, nil
}

func (o *skipParser) Parse(in parser.Reader) (parser.Empty, error) {
	b, err := in.ReadByte()
	if err != nil {
		return nil, err
	}

	if !o.predicate(b) {
		_, _ = in.Seek(-1, io.SeekCurrent)
		return nil, errors.ErrNotMatched
	}

	return nil, nil
}

// Skip will skip over a byte if it matches the predicate.
func Skip(predicte parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return &skipParser{predicate: predicte}
}

// SkipWhile will skip over zero or more bytes that match the predicate.
func SkipWhile(p parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return SkipWhileMinMax(0, math.MaxInt, p)
}

// SkipWhile1 will skip over one or more bytes that match the predicate.
func SkipWhile1(p parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return SkipWhileMinMax(1, math.MaxInt, p)
}

// SkipWhileMinMax skip over the longest (m <= len <= n) input slice that matches the predicate.
//   - If the input matches the predicate and (m <= len <= n), it will skip over the matched bytes.
//   - If the input is empty and m > 0, it will return io.EOF
//   - If the number of matched bytes < m, it will return errors.ErrNotMatched
func SkipWhileMinMax(min, max int, predicate parser.Predicate[byte]) parser.Parser[parser.Reader, parser.Empty] {
	return &skipWhileMinMaxParser{min: min, max: max, predicate: predicate}
}
