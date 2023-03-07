package bytes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"math"
)

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
	return func(in parser.Reader) ([]byte, error) {
		result, err := readMin(min, predicate, in)
		if err != nil {
			return nil, err
		}

		for i := min; i < max; i++ {
			b, err := in.ReadByte()
			if err != nil {
				return result, nil
			}

			if !predicate(b) {
				_, _ = in.Seek(-1, io.SeekCurrent)
				return result, nil
			}

			result = append(result, b)
		}

		return result, nil
	}
}

func readMin(m int, p parser.Predicate[byte], in parser.Reader) ([]byte, error) {
	result := make([]byte, m)
	if m > 0 {
		n, err := in.Read(result)
		if err != nil {
			_, _ = in.Seek(-int64(n), io.SeekCurrent)
			return nil, err
		}

		for _, r := range result {
			if !p(r) {
				_, _ = in.Seek(-int64(n), io.SeekCurrent)
				return nil, parser.ErrNotMatched
			}
		}
	}
	return result, nil
}
