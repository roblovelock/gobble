package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	recognizeParser[R parser.Reader, T any] struct {
		parser parser.Parser[R, T]
	}
)

func (o *recognizeParser[R, T]) Parse(in R) ([]byte, error) {
	startOffset, _ := in.Seek(0, io.SeekCurrent)
	_, err := o.parser.Parse(in)
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

// Recognize If the child parser was successful, return the consumed input as produced value.
func Recognize[R parser.Reader, T any](p parser.Parser[R, T]) parser.Parser[R, []byte] {
	return &recognizeParser[R, T]{parser: p}
}
