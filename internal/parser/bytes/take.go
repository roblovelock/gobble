package bytes

import (
	"gobble/internal/parser"
	"io"
)

func Take[R parser.Reader](n uint) parser.Parser[R, []byte] {
	return func(in R) ([]byte, error) {
		b := make([]byte, n)
		_, err := io.ReadFull(in, b)
		if err != nil {
			return nil, io.EOF
		}

		return b, nil
	}
}
