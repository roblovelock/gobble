package bytes

import (
	"gobble/internal/parser"
)

func IsA(bytes ...byte) parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		result := make([]byte, 0)
		for {
			b, err := OneOf(bytes...)(in)
			if err != nil {
				if len(result) == 0 {
					return nil, err
				}
				break
			}
			result = append(result, b)
		}
		return result, nil
	}
}
