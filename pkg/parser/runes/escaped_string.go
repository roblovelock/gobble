package runes

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"strings"
	"unicode/utf16"
)

func EscapedString() parser.Parser[parser.Reader, string] {
	return func(in parser.Reader) (string, error) {
		if b, err := in.ReadByte(); err != nil {
			return "", err
		} else if b != '"' {
			_, _ = in.Seek(-1, io.SeekCurrent)
			return "", parser.ErrNotMatched
		}

		currentOffset, _ := in.Seek(0, io.SeekCurrent)
		var builder strings.Builder
		for {
			r, _, err := in.ReadRune()
			if err != nil {
				_, _ = in.Seek(currentOffset-1, io.SeekStart)
				return "", err
			}

			if r == '"' {
				break
			}

			if r == '\\' {
				r, err = readSpecial(in)
				if err != nil {
					_, _ = in.Seek(currentOffset-1, io.SeekStart)
					return "", err
				}
			}

			builder.WriteRune(r)
		}

		return builder.String(), nil
	}
}

func readSpecial(in parser.Reader) (rune, error) {
	b, err := in.ReadByte()
	if err != nil {
		return 0, err
	}

	switch b {
	case '"', '\\', '/', '\'':
		return rune(b), nil
	case 'b':
		return '\b', nil
	case 'f':
		return '\f', nil
	case 'n':
		return '\n', nil
	case 'r':
		return '\r', nil
	case 't':
		return '\t', nil
	case 'u':
		return decodeUnicode(in)
	default:
		return 0, parser.ErrNotMatched
	}
}

func decodeUnicode(in parser.Reader) (rune, error) {
	unicodeBuffer := make([]byte, 6)
	if n, err := in.Read(unicodeBuffer[:4]); err != nil {
		return 0, err
	} else if n != 4 {
		return 0, io.EOF
	}
	r, err := unicodeToRune(unicodeBuffer[:4])
	if err != nil {
		return 0, err
	}
	if utf16.IsSurrogate(r) {
		if n, err := in.Read(unicodeBuffer); err != nil {
			return 0, err
		} else if n != 6 {
			return 0, io.EOF
		}
		if unicodeBuffer[0] != '\\' || unicodeBuffer[1] != 'u' {
			return 0, parser.ErrNotMatched
		}
		r2, err := unicodeToRune(unicodeBuffer[2:])
		if err != nil {
			return 0, err
		}
		r = utf16.DecodeRune(r, r2)
	}

	return r, nil
}

func unicodeToRune(code []byte) (rune, error) {
	var r rune
	for i := 0; i < len(code); i++ {
		h, err := hexToRune(code[i])
		if err != nil {
			return 0, err
		}
		r = r*16 + h
	}
	return r, nil
}

func hexToRune(b byte) (rune, error) {
	switch b {
	case '0':
		return 0, nil
	case '1':
		return 1, nil
	case '2':
		return 2, nil
	case '3':
		return 3, nil
	case '4':
		return 4, nil
	case '5':
		return 5, nil
	case '6':
		return 6, nil
	case '7':
		return 7, nil
	case '8':
		return 8, nil
	case '9':
		return 9, nil
	case 'A':
		return 10, nil
	case 'B':
		return 11, nil
	case 'C':
		return 12, nil
	case 'D':
		return 13, nil
	case 'E':
		return 14, nil
	case 'F':
		return 15, nil
	case 'a':
		return 10, nil
	case 'b':
		return 11, nil
	case 'c':
		return 12, nil
	case 'd':
		return 13, nil
	case 'e':
		return 14, nil
	case 'f':
		return 15, nil
	default:
		return 0, parser.ErrNotMatched
	}
}
