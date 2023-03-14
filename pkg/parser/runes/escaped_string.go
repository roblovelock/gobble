package runes

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"io"
	"unicode/utf16"
)

type (
	escapedStringParser struct{}
)

var escapedStringParserInstance = &escapedStringParser{}

func (o *escapedStringParser) Parse(in parser.Reader) (string, error) {
	if b, err := in.ReadByte(); err != nil {
		return "", err
	} else if b != '"' {
		_, _ = in.Seek(-1, io.SeekCurrent)
		return "", errors.ErrNotMatched
	}

	var result []rune
	startOffset, _ := in.Seek(0, io.SeekCurrent)
	currentOffset := startOffset
	for {
		r, _, err := in.ReadRune()
		if err != nil {
			_, _ = in.Seek(startOffset-1, io.SeekStart)
			return "", err
		}

		if r == '"' {
			break
		}

		if r == '\\' {
			result = append(result, readRunes(in, currentOffset)...)
			r, err = readSpecial(in)
			if err != nil {
				_, _ = in.Seek(startOffset-1, io.SeekStart)
				return "", err
			}
			result = append(result, r)
			currentOffset, _ = in.Seek(0, io.SeekCurrent)
		}
	}

	result = append(result, readRunes(in, currentOffset)...)
	return string(result), nil
}

func (*escapedStringParser) ParseBytes(in []byte) (string, []byte, error) {
	//TODO
	return "", in, errors.ErrNotSupported
}

func EscapedString() parser.Parser[parser.Reader, string] {
	return escapedStringParserInstance
}

func readRunes(in parser.Reader, start int64) []rune {
	end, _ := in.Seek(0, io.SeekCurrent)
	_, _ = in.Seek(start, io.SeekStart)
	result := make([]byte, end-start-1)
	_, _ = in.Read(result)
	_, _ = in.Seek(end, io.SeekStart)

	return []rune(string(result))
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
		return 0, errors.ErrNotMatched
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
			return 0, errors.ErrNotMatched
		}
		r2, err := unicodeToRune(unicodeBuffer[2:])
		if err != nil {
			return 0, err
		}
		r = utf16.DecodeRune(r, r2)
	}

	return r, nil
}
