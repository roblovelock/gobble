package parser

func Untyped[R Reader, T any](p Parser[R, T]) Parser[R, interface{}] {
	return func(in R) (interface{}, error) {
		return p(in)
	}
}

func Typed[R Reader, T any](p Parser[R, interface{}]) Parser[R, T] {
	return func(in R) (T, error) {
		r, err := p(in)
		val, ok := r.(T)
		if !ok {
			var t T
			return t, ErrNotMatched
		}
		return val, err
	}
}
