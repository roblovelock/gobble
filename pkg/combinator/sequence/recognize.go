package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

// Recognize If the child parser was successful, return the consumed input as produced value.
func Recognize[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, []byte] {
	return func(in R) ([]byte, error) {
		startOffset, _ := in.Seek(0, io.SeekCurrent)
		_, err := p(in)
		if err != nil {
			_, _ = in.Seek(startOffset, io.SeekStart)
			return nil, err
		}
		endOffset, _ := in.Seek(0, io.SeekCurrent)
		_, _ = in.Seek(startOffset, io.SeekStart)
		result := make([]byte, endOffset-startOffset)
		_, _ = in.Read(result)
		return result, nil
	}
}
