package bytes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
)

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
	return func(in parser.Reader) (string, error) {
		startOffset, _ := in.Seek(0, io.SeekCurrent)
		for {
			if b, err := in.ReadByte(); err != nil {
				break
			} else if normal(b) {
				continue
			} else if control != b {
				_, _ = in.Seek(-1, io.SeekCurrent)
				break
			}

			if b, err := in.ReadByte(); err == nil {
				_, _ = in.Seek(-1, io.SeekCurrent)
				break
			} else if !escapable(b) {
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
}
