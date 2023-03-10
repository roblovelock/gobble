package bytes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

type (
	escapedParser struct {
		normal    parser.Predicate[byte]
		control   byte
		escapable parser.Predicate[byte]
	}
)

func (o *escapedParser) Parse(in parser.Reader) (string, error) {
	startOffset, _ := in.Seek(0, io.SeekCurrent)
	for {
		if b, err := in.ReadByte(); err != nil {
			break
		} else if o.normal(b) {
			continue
		} else if o.control != b {
			_, _ = in.Seek(-1, io.SeekCurrent)
			break
		}

		if b, err := in.ReadByte(); err == nil {
			_, _ = in.Seek(-1, io.SeekCurrent)
			break
		} else if !o.escapable(b) {
			_, _ = in.Seek(-2, io.SeekCurrent)
			break
		}
	}
	endOffset, _ := in.Seek(0, io.SeekCurrent)
	_, _ = in.Seek(startOffset, io.SeekStart)
	result := make([]byte, endOffset-startOffset)
	_, _ = in.Read(result)
	return string(result), nil
}

func (o *escapedParser) ParseBytes(in []byte) (string, []byte, error) {
	n := 0
	for ; n < len(in); n++ {
		if o.normal(in[n]) {
			continue
		}
		if o.control != in[n] {
			break
		}

		if len(in) < n+1 {
			n--
			break
		}
		if !o.escapable(in[n+1]) {
			n--
			break
		}
		n++
	}
	if n == 0 {
		if len(in) == 0 {
			return "", in, io.EOF
		}
		return "", in, errors.ErrNotMatched
	}
	return string(in[:n]), in[n:], nil
}

// Escaped matches a byte stream with escape characters
//
//   - The first argument matches the normal characters (it must not accept the control character)
//   - The second argument is the control character (like \ in most languages)
//   - The third argument matches the escaped characters
func Escaped(
	normal parser.Predicate[byte],
	control byte,
	escapable parser.Predicate[byte],
) parser.Parser[parser.Reader, string] {
	return &escapedParser{normal: normal, control: control, escapable: escapable}
}
