package modifier

import "github.com/roblovelock/gobble/pkg/parser"

// Optional will call the parser and suppress any error returned
func Optional[R parser.Reader, T any](parser parser.Parser[R, T]) parser.Parser[R, T] {
	return func(in R) (T, error) {
		v, _ := parser(in)
		return v, nil
	}
}
