package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// KeyValue returns a single element map, where the key is the result from the first parser and the value is
// the result from the final parser. The key and value parsers are separated by the value of the second parser.
func KeyValue[R parser.Reader, F comparable, S, T any](
	key parser.Parser[R, F], separator parser.Parser[R, S], value parser.Parser[R, T],
) parser.Parser[R, map[F]T] {
	return func(in R) (map[F]T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		f, err := key(in)
		if err != nil {
			return nil, err
		}

		if _, err := separator(in); err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return nil, err
		}

		s, err := value(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return nil, err
		}
		return map[F]T{f: s}, err
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
	return func(in R) (map[F]T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		f, err := key(in)
		if err != nil {
			return nil, err
		}

		if _, err := s1(in); err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return nil, err
		}

		s, err := value(in)
		if err != nil {
			_, _ = in.Seek(currentOffset, io.SeekStart)
			return nil, err
		}

		result := map[F]T{f: s}

		for {
			currentOffset, _ = in.Seek(0, io.SeekCurrent)
			if _, err := s2(in); err != nil {
				break
			}

			f, err = key(in)
			if err != nil {
				_, _ = in.Seek(currentOffset, io.SeekStart)
				break
			}

			if _, err := s1(in); err != nil {
				_, _ = in.Seek(currentOffset, io.SeekStart)
				break
			}

			s, err = value(in)
			if err != nil {
				_, _ = in.Seek(currentOffset, io.SeekStart)
				break
			}
			result[f] = s
		}

		return result, nil
	}
}

// KeyValue0 returns a map of key value pairs.
//
//   - The first parser is the key.
//   - The second parser is the separator between the key and value.
//   - The third parser is the value.
//   - The fourth parser is the separator between each key value pair.
//
// When a valid key value pair isn't found after the first element, the fourth parsers value isn't consumed from the
// input
func KeyValue0[R parser.Reader, F comparable, S1, S2, T any](
	key parser.Parser[R, F], s1 parser.Parser[R, S1], value parser.Parser[R, T], s2 parser.Parser[R, S2],
) parser.Parser[R, map[F]T] {
	return func(in R) (map[F]T, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		result := make(map[F]T)
		for {
			f, err := key(in)
			if err != nil {
				_, _ = in.Seek(currentOffset, io.SeekStart)
				break
			}

			if _, err := s1(in); err != nil {
				_, _ = in.Seek(currentOffset, io.SeekStart)
				break
			}

			s, err := value(in)
			if err != nil {
				_, _ = in.Seek(currentOffset, io.SeekStart)
				break
			}
			result[f] = s

			currentOffset, _ = in.Seek(0, io.SeekCurrent)
			if _, err := s2(in); err != nil {
				break
			}
		}

		return result, nil
	}
}
