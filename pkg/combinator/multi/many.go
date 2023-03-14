package multi

import (
	"github.com/roblovelock/gobble/pkg/parser"
)

type (
	many0Parser[R parser.Reader, T any] struct {
		parser parser.Parser[R, T]
	}

	many0CountParser[R parser.Reader, T any] struct {
		parser parser.Parser[R, T]
	}

	many1Parser[R parser.Reader, T any] struct {
		parser parser.Parser[R, T]
	}

	many1CountParser[R parser.Reader, T any] struct {
		parser parser.Parser[R, T]
	}
)

func (o *many0Parser[R, T]) Parse(in R) ([]T, error) {
	result := make([]T, 0)
	for {
		r, err := o.parser.Parse(in)
		if err != nil {
			break
		}
		result = append(result, r)
	}

	return result, nil
}

func (o *many0Parser[R, T]) ParseBytes(in []byte) (result []T, out []byte, err error) {
	var r T
	out = in
	for {
		r, out, err = o.parser.ParseBytes(out)
		if err != nil {
			return result, out, nil
		}
		result = append(result, r)
	}
}

func (o *many0CountParser[R, T]) Parse(in R) (uint, error) {
	var count uint = 0
	for {
		if _, err := o.parser.Parse(in); err != nil {
			break
		}
		count++
	}

	return count, nil
}

func (o *many0CountParser[R, T]) ParseBytes(in []byte) (count uint, out []byte, err error) {
	out = in
	for {
		_, out, err = o.parser.ParseBytes(out)
		if err != nil {
			return count, out, nil
		}
		count++
	}
}

func (o *many1Parser[R, T]) Parse(in R) ([]T, error) {
	r, err := o.parser.Parse(in)
	if err != nil {
		return nil, err
	}

	result := []T{r}
	for r, err := o.parser.Parse(in); err == nil; r, err = o.parser.Parse(in) {
		result = append(result, r)
	}

	return result, nil
}

func (o *many1Parser[R, T]) ParseBytes(in []byte) (result []T, out []byte, err error) {
	var r T
	out = in
	for {
		r, out, err = o.parser.ParseBytes(out)
		if err != nil {
			if len(result) == 0 {
				return result, in, err
			}
			return result, out, nil
		}
		result = append(result, r)
	}
}

func (o *many1CountParser[R, T]) Parse(in R) (uint, error) {
	if _, err := o.parser.Parse(in); err != nil {
		return 0, err
	}

	var count uint = 1
	for _, err := o.parser.Parse(in); err == nil; _, err = o.parser.Parse(in) {
		count++
	}

	return count, nil
}

func (o *many1CountParser[R, T]) ParseBytes(in []byte) (count uint, out []byte, err error) {
	out = in
	for {
		_, out, err = o.parser.ParseBytes(out)
		if err != nil {
			if count == 0 {
				return count, in, err
			}
			return count, out, nil
		}
		count++
	}
}

func Many0[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, []T] {
	return &many0Parser[R, T]{parser: p}
}

func Many0Count[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, uint] {
	return &many0CountParser[R, T]{parser: p}
}

func Many1[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, []T] {
	return &many1Parser[R, T]{parser: p}
}

func Many1Count[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, uint] {
	return &many1CountParser[R, T]{parser: p}
}
