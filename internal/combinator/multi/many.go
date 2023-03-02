package multi

import (
	"gobble/internal/parser"
)

func Many0[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, []T] {
	return func(in R) ([]T, error) {
		result := make([]T, 0)
		for r, err := p(in); err == nil; r, err = p(in) {
			result = append(result, r)
		}

		return result, nil
	}
}

func Many0Count[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, uint] {
	return func(in R) (uint, error) {
		var count uint = 0
		for _, err := p(in); err == nil; _, err = p(in) {
			count++
		}

		return count, nil
	}
}

func Many1[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, []T] {
	return func(in R) ([]T, error) {
		r, err := p(in)
		if err != nil {
			return nil, err
		}

		result := []T{r}
		for r, err := p(in); err == nil; r, err = p(in) {
			result = append(result, r)
		}

		return result, nil
	}
}

func Many1Count[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, uint] {
	return func(in R) (uint, error) {
		if _, err := p(in); err != nil {
			return 0, err
		}

		var count uint = 1
		for _, err := p(in); err == nil; _, err = p(in) {
			count++
		}

		return count, nil
	}
}
