package multi

import (
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	keyValueParser[R parser.Reader, F comparable, S1, S2, T any] struct {
		key   parser.Parser[R, F]
		s1    parser.Parser[R, S1]
		value parser.Parser[R, T]
		s2    parser.Parser[R, S2]
	}
)

func (o *keyValueParser[R, F, S1, S2, T]) Parse(in R) (map[F]T, error) {
	currentOffset, _ := in.Seek(0, io.SeekCurrent)
	result := make(map[F]T, 7)
	for {
		f, err := o.key.Parse(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			break
		}

		if _, err := o.s1.Parse(in); err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			break
		}

		s, err := o.value.Parse(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			break
		}
		result[f] = s

		currentOffset, _ = in.Seek(0, io.SeekCurrent)
		if _, err := o.s2.Parse(in); err != nil {
			break
		}
	}

	return result, nil
}

func (o *keyValueParser[R, F, S1, S2, T]) ParseBytes(in []byte) (map[F]T, []byte, error) {
	result := make(map[F]T, 7)
	for {
		f, out, err := o.key.ParseBytes(in)
		if err != nil {
			return result, in, nil
		}

		_, out, err = o.s1.ParseBytes(out)
		if err != nil {
			return result, in, nil
		}

		s, out, err := o.value.ParseBytes(out)
		if err != nil {
			return result, in, nil
		}
		result[f] = s

		_, out, err = o.s2.ParseBytes(out)
		if err != nil {
			return result, out, nil
		}
		in = out
	}
}

// KeyValue returns a map of key value pairs.
//
//   - The first parser is the key.
//   - The second parser is the separator between the key and value.
//   - The third parser is the value.
//   - The fourth parser is the separator between each key value pair.
//
// When a valid key value pair isn't found after the first element, the fourth parsers value isn't consumed from the
// input
func KeyValue[R parser.Reader, F comparable, S1, S2, T any](
	key parser.Parser[R, F], s1 parser.Parser[R, S1], value parser.Parser[R, T], s2 parser.Parser[R, S2],
) parser.Parser[R, map[F]T] {
	return &keyValueParser[R, F, S1, S2, T]{
		key: key, s1: s1, value: value, s2: s2,
	}
}

// KeyValue1 returns a map of key value pairs containing at least one element.
//
//   - The first parser is the key.
//   - The second parser is the separator between the key and value.
//   - The third parser is the value.
//   - The fourth parser is the separator between each key value pair.
//
// When a valid key value pair isn't found after the first element, the fourth parsers value isn't consumed from the
// input
func KeyValue1[R parser.Reader, F comparable, S1, S2, T any](
	key parser.Parser[R, F], s1 parser.Parser[R, S1], value parser.Parser[R, T], s2 parser.Parser[R, S2],
) parser.Parser[R, map[F]T] {
	return modifier.Verify(
		KeyValue(key, s1, value, s2),
		func(m map[F]T) bool {
			return len(m) > 0
		},
	)
}
