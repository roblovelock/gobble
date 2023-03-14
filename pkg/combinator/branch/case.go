package branch

import (
	"github.com/roblovelock/gobble/pkg/combinator"
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	peekCaseParser[R parser.Reader, T any] struct {
		parsers map[byte]parser.Parser[R, T]
	}

	caseParser[R parser.Reader, C comparable, T any] struct {
		parser        parser.Parser[R, C]
		parsers       map[C]parser.Parser[R, T]
		defaultParser parser.Parser[R, T]
	}
)

func (o *caseParser[R, C, T]) Parse(in R) (T, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	c, err := o.parser.Parse(in)
	if err != nil {
		var t T
		return t, err
	}
	choice, ok := o.parsers[c]
	if !ok {
		choice = o.defaultParser
	}

	result, err := choice.Parse(in)
	if err != nil {
		_, _ = in.Seek(currentOffset, io.SeekStart)
	}

	return result, err
}

func (o *caseParser[R, C, T]) ParseBytes(in []byte) (T, []byte, error) {
	c, out, err := o.parser.ParseBytes(in)
	if err != nil {
		var t T
		return t, in, err
	}
	choice, ok := o.parsers[c]
	if !ok {
		choice = o.defaultParser
	}

	result, out, err := choice.ParseBytes(out)
	if err != nil {
		return result, in, err
	}

	return result, out, nil
}

func (o *peekCaseParser[R, T]) Parse(in R) (T, error) {
	b, err := in.ReadByte()
	if err != nil {
		var t T
		return t, err
	}
	_, _ = in.Seek(-1, io.SeekCurrent)
	p, ok := o.parsers[b]
	if !ok {
		var t T
		return t, errors.ErrNotMatched
	}

	return p.Parse(in)
}

func (o *peekCaseParser[R, T]) ParseBytes(in []byte) (T, []byte, error) {
	if len(in) == 0 {
		var t T
		return t, in, io.EOF
	}
	p, ok := o.parsers[in[0]]
	if !ok {
		var t T
		return t, in, errors.ErrNotMatched
	}

	return p.ParseBytes(in)
}

// Case will choose which parser should process the input stream, from the provided map of parsers, based on the result
// of the initial parser.
func Case[R parser.Reader, C comparable, T any](
	p parser.Parser[R, C], parsers map[C]parser.Parser[R, T],
) parser.Parser[R, T] {
	return CaseOrDefault(p, parsers, combinator.Fail[R, T](errors.ErrNotMatched))
}

// CaseOrDefault will choose which parser should process the input stream, from the provided map of parsers, based on
// the result of the initial parser. If the provided map doesn't contain a value the default parser will be used.
func CaseOrDefault[R parser.Reader, C comparable, T any](
	p parser.Parser[R, C], parsers map[C]parser.Parser[R, T], d parser.Parser[R, T],
) parser.Parser[R, T] {
	return &caseParser[R, C, T]{parser: p, parsers: parsers, defaultParser: d}
}

// PeekCase will look ahead one byte and choose which parser should process the input stream, from the provided map of
// parsers, based on the next byte value.
func PeekCase[R parser.Reader, T any](parsers map[byte]parser.Parser[R, T]) parser.Parser[R, T] {
	return &peekCaseParser[R, T]{parsers: parsers}
}
