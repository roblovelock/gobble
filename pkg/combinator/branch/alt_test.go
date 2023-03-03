package branch

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/pkg/parser"
	"gobble/pkg/parser/bytes"
	"gobble/pkg/parser/runes"
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
			name: "empty input => EOF",
			args: args{
				first:  runes.Rune('a'),
				second: bytes.Byte('b'),
				input:  strings.NewReader(""),
			},
			wantErr: io.EOF,
		},
		{
			name: "mismatch => no match",
			args: args{
				first:  runes.Rune('a'),
				second: bytes.Byte('b'),
				input:  strings.NewReader("c"),
			},
			wantRemain: "c",
			wantErr:    parser.ErrNotMatched,
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