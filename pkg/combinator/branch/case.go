package branch

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Case will choose which parser should process the input stream, from the provided map of parsers, based on the result
// of the initial parser.
func Case[R parser.Reader, C comparable, T any](
	p parser.Parser[R, C], parsers map[C]parser.Parser[R, T],
) parser.Parser[R, T] {
	return func(in R) (T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		c, err := p(in)
		if err != nil {
			var t T
			return t, err
		}
		choice, ok := parsers[c]
		if !ok {
			var t T
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return t, errors.ErrNotMatched
		}

		result, err := choice(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
		}

		return result, err
	}
}

// CaseOrDefault will choose which parser should process the input stream, from the provided map of parsers, based on
// the result of the initial parser. If the provided map doesn't contain a value the default parser will be used.
func CaseOrDefault[R parser.Reader, C comparable, T any](
	p parser.Parser[R, C], parsers map[C]parser.Parser[R, T], d parser.Parser[R, T],
) parser.Parser[R, T] {
	return func(in R) (T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		c, err := p(in)
		if err != nil {
			var t T
			return t, err
		}

		choice, ok := parsers[c]
		if !ok {
			choice = d
		}

		result, err := choice(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
		}

		return result, err
	}
}

// PeekCase will look ahead one byte and choose which parser should process the input stream, from the provided map of
// parsers, based on the next byte value.
func PeekCase[R parser.Reader, T any](parsers map[byte]parser.Parser[R, T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		b, err := in.ReadByte()
		if err != nil {
			var t T
			return t, err
		}
		_, _ = in.Seek(-1, io.SeekCurrent)
		p, ok := parsers[b]
		if !ok {
			var t T
			return t, errors.ErrNotMatched
		}

		return p(in)
	}
}
