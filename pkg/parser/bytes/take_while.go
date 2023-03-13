package bytes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"math"
)

type (
	takeWhileMinMaxParser struct {
		min       int
		max       int
		predicate parser.Predicate[byte]
	}
)

func (o *takeWhileMinMaxParser) Parse(in parser.Reader) ([]byte, error) {
	startOffset, _ := in.Seek(0, io.SeekCurrent)
	n := 0
	for ; n < o.max; n++ {
		b, err := in.ReadByte()
		if err != nil {
			if n < o.min {
				_, _ = in.Seek(startOffset, io.SeekStart)
				return nil, err
			}
			break
		}

		if !o.predicate(b) {
			break
		}
	}

	_, _ = in.Seek(startOffset, io.SeekStart)
	if n < o.min {
		return nil, errors.ErrNotMatched
	}
	result := make([]byte, n)
	_, _ = in.Read(result)

	return result, nil
}

// TakeWhile Returns zero or more bytes that match the predicate.
//   - If the input matches the predicate, it will return the matched bytes.
//   - If the input is empty, it will return an empty slice
//   - If the input doesn't match the predicate, it will return an empty slice
func TakeWhile(predicate parser.Predicate[byte]) parser.Parser[parser.Reader, []byte] {
	return TakeWhileMinMax(0, math.MaxInt, predicate)
}

// TakeWhile1 Returns one or more bytes that match the predicate.
//   - If the input matches the predicate, it will return the matched bytes.
//   - If the input is empty, it will return io.EOF
//   - If the input doesn't match the predicate, it will return parser.ErrNotMatched
func TakeWhile1(predicate parser.Predicate[byte]) parser.Parser[parser.Reader, []byte] {
	return TakeWhileMinMax(1, math.MaxInt, predicate)
}

// TakeWhileMinMax Returns the longest (m <= len <= n) input slice that matches the predicate.
//   - If the input matches the predicate and (m <= len <= n), it will return the matched bytes.
//   - If the input is empty and m > 0, it will return io.EOF
//   - If the number of matched bytes < m, it will return parser.ErrNotMatched
func TakeWhileMinMax(min, max int, predicate parser.Predicate[byte]) parser.Parser[parser.Reader, []byte] {
	return &takeWhileMinMaxParser{min: min, max: max, predicate: predicate}
}
