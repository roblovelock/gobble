package bytes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	escapedTransformParser struct {
		normal    parser.Predicate[byte]
		control   byte
		transform parser.MapFunc[byte, byte]
	}
)

func (o *escapedTransformParser) Parse(in parser.Reader) ([]byte, error) {
	var result []byte
	startOffset, _ := in.Seek(0, io.SeekCurrent)
	for {
		b, err := in.ReadByte()
		if err != nil {
			break
		}

		if o.normal(b) {
			continue
		}

		if o.control != b {
			_, _ = in.Seek(-1, io.SeekCurrent)
			break
		}

		b, err = in.ReadByte()
		if err != nil {
			_, _ = in.Seek(-1, io.SeekCurrent)
			break
		}

		bytes, err := o.transform(b)
		if err != nil {
			_, _ = in.Seek(-2, io.SeekCurrent)
			break
		}

		endOffset, _ := in.Seek(-2, io.SeekCurrent)
		_, _ = in.Seek(startOffset, io.SeekStart)
		length := endOffset - startOffset
		chunk := make([]byte, length, length+1)
		_, _ = in.Read(chunk)
		result = append(result, append(chunk, bytes)...)
		startOffset, _ = in.Seek(2, io.SeekCurrent)
	}
	endOffset, _ := in.Seek(0, io.SeekCurrent)
	_, _ = in.Seek(startOffset, io.SeekStart)
	chunk := make([]byte, endOffset-startOffset)
	_, _ = in.Read(chunk)
	result = append(result, chunk...)
	return result, nil
}

func (o *escapedTransformParser) ParseBytes(in []byte) ([]byte, []byte, error) {
	var result []byte
	n := 0
	out := in
	for ; n < len(out); n++ {
		if o.normal(out[n]) {
			continue
		}
		if o.control != out[n] {
			break
		}

		if len(in) < n+1 {
			n--
			break
		}
		bytes, err := o.transform(out[n+1])
		if err != nil {
			n--
			break
		}
		result = append(result, append(out[:n-1], bytes)...)
		n = 0
		out = in[n+1:]
	}
	if result == nil {
		return out[:n], out[n:], nil
	}
	result = append(result, out[:n]...)
	return result, out[n:], nil
}

// EscapedTransform matches a byte stream with escape characters and transforms them using the transform function
//
//   - The first argument matches the normal characters (it must not accept the control character)
//   - The second argument is the control character (like \ in most languages)
//   - The third argument matches the escaped characters
func EscapedTransform(
	normal parser.Predicate[byte],
	control byte,
	transform parser.MapFunc[byte, byte],
) parser.Parser[parser.Reader, []byte] {
	return &escapedTransformParser{normal: normal, control: control, transform: transform}
}
