package parser

type (
	pointerParser[R Reader, T any] struct {
		parser *Parser[R, T]
	}
)

func (o *pointerParser[R, T]) Parse(in R) (T, error) {
	return (*o.parser).Parse(in)
}

func (o *pointerParser[R, T]) ParseBytes(in []byte) (T, []byte, error) {
	return (*o.parser).ParseBytes(in)
}

func Pointer[R Reader, T any](p *Parser[R, T]) Parser[R, T] {
	return &pointerParser[R, T]{parser: p}
}
