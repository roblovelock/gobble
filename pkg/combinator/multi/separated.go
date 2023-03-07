package multi

import (
	"github.com/roblovelock/gobble/pkg/parser"
)

func Separated0[R parser.Reader, T any, S any](
	p parser.Parser[R, T], separator parser.Parser[R, S],
) parser.Parser[R, []T] {
	return func(in R) ([]T, error) {
		result := make([]T, 0)
		for r, err := p(in); err == nil; r, err = p(in) {
			result = append(result, r)
			if _, err := separator(in); err != nil {
				break
			}
		}

		return result, nil
	}
}

func Separated0Count[R parser.Reader, T any, S any](
	p parser.Parser[R, T], separator parser.Parser[R, S],
) parser.Parser[R, uint] {
	return func(in R) (uint, error) {
		var count uint = 0
		for _, err := p(in); err == nil; _, err = p(in) {
			count++
			if _, err := separator(in); err != nil {
				break
			}
		}

		return count, nil
	}
}

func Separated1[R parser.Reader, T any, S any](
	p parser.Parser[R, T], separator parser.Parser[R, S],
) parser.Parser[R, []T] {
	return func(in R) ([]T, error) {
		r, err := p(in)
		if err != nil {
			return nil, err
		}

		result := []T{r}
		for r, err := p(in); err == nil; r, err = p(in) {
			result = append(result, r)
			if _, err := separator(in); err != nil {
				break
			}
		}

		return result, nil
	}
}

func Separated1Count[R parser.Reader, T any, S any](
	p parser.Parser[R, T], separator parser.Parser[R, S],
) parser.Parser[R, uint] {
	return func(in R) (uint, error) {
		if _, err := p(in); err != nil {
			return 0, err
		}

		var count uint = 1
		for _, err := p(in); err == nil; _, err = p(in) {
			count++
			if _, err := separator(in); err != nil {
				break
			}
		}

		return count, nil
	}
}
