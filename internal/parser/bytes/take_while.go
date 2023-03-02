package bytes

import (
	"gobble/internal/parser"
	"io"
)

func TakeWhileMN(m, n int, p parser.Predicate[byte]) parser.Parser[parser.Reader, []byte] {
	return func(in parser.Reader) ([]byte, error) {
		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		result := make([]byte, 0)

		for i := 0; i < n; i++ {
			r, err := in.ReadByte()
			if err == nil && !p(r) {
				err = parser.ErrNotMatched
			}
			if err != nil {
				if len(result) < m {
					_, _ = in.Seek(currentOffset, io.SeekStart)
					return nil, err
				}
				break
			}
			result = append(result, r)
		}

		return result, nil
	}
}
