package branch

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestAlt(t *testing.T) {
	type args struct {
		first  parser.Parser[parser.Reader, rune]
		second parser.Parser[parser.Reader, byte]
		input  parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  interface{}
		wantRemain string
		wantErr    error
	}{
		{
			name: "empty input => no match",
			args: args{
				first:  runes.Rune('a'),
				second: bytes.Byte('b'),
				input:  strings.NewReader(""),
			},
			wantErr: errors.ErrNotMatched,
		},
		{
			name: "mismatch => no match",
			args: args{
				first:  runes.Rune('a'),
				second: bytes.Byte('b'),
				input:  strings.NewReader("c"),
			},
			wantRemain: "c",
			wantErr:    errors.ErrNotMatched,
		},
		{
			name: "first match => match",
			args: args{
				first:  runes.Rune('😀'),
				second: bytes.Byte('b'),
				input:  strings.NewReader("😀"),
			},
			wantMatch: '😀',
		},
		{
			name: "second match => match",
			args: args{
				first:  runes.Rune('a'),
				second: bytes.Byte('b'),
				input:  strings.NewReader("b😀"),
			},
			wantMatch:  byte('b'),
			wantRemain: "😀",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Alt(parser.Untyped(tt.args.first), parser.Untyped(tt.args.second))
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, string(remain))
			}
		})
	}
}
