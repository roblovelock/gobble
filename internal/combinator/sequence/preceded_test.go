package sequence

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/internal/parser"
	"gobble/internal/parser/character"
	"io"
	"strings"
	"testing"
)

func TestPreceded(t *testing.T) {
	type args struct {
		first  parser.Parser[parser.Reader, rune]
		second parser.Parser[parser.Reader, rune]
		input  parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  rune
		wantRemain string
		wantErr    error
	}{
		{
			name: "empty input => EOF",
			args: args{
				first:  character.Rune('a'),
				second: character.Rune('b'),
				input:  strings.NewReader(""),
			},
			wantErr: io.EOF,
		},
		{
			name: "second rune EOF => EOF",
			args: args{
				first:  character.Rune('a'),
				second: character.Rune('b'),
				input:  strings.NewReader("a"),
			},
			wantRemain: "a",
			wantErr:    io.EOF,
		},
		{
			name: "first rune mismatch => no match",
			args: args{
				first:  character.Rune('a'),
				second: character.Rune('b'),
				input:  strings.NewReader("bb"),
			},
			wantRemain: "bb",
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "second rune mismatch => no match",
			args: args{
				first:  character.Rune('a'),
				second: character.Rune('b'),
				input:  strings.NewReader("aa"),
			},
			wantRemain: "aa",
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "second rune match => second match",
			args: args{
				first:  character.Rune('a'),
				second: character.Rune('b'),
				input:  strings.NewReader("ab"),
			},
			wantMatch: 'b',
		},
		{
			name: "second rune match unicode => second match",
			args: args{
				first:  character.Rune('a'),
				second: character.Rune('😀'),
				input:  strings.NewReader("a😀😀"),
			},
			wantMatch:  '😀',
			wantRemain: "😀",
		},
		{
			name: "first rune match unicode => second match",
			args: args{
				first:  character.Rune('😀'),
				second: character.Rune('a'),
				input:  strings.NewReader("😀a"),
			},
			wantMatch: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Preceded(tt.args.first, tt.args.second)
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
