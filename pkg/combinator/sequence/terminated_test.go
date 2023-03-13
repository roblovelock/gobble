package sequence

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestTerminated(t *testing.T) {
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
				first:  runes.Rune('a'),
				second: runes.Rune('b'),
				input:  strings.NewReader(""),
			},
			wantErr: io.EOF,
		},
		{
			name: "second rune EOF => EOF",
			args: args{
				first:  runes.Rune('a'),
				second: runes.Rune('b'),
				input:  strings.NewReader("a"),
			},
			wantRemain: "a",
			wantErr:    io.EOF,
		},
		{
			name: "first rune mismatch => no match",
			args: args{
				first:  runes.Rune('a'),
				second: runes.Rune('b'),
				input:  strings.NewReader("bb"),
			},
			wantRemain: "bb",
			wantErr:    errors.ErrNotMatched,
		},
		{
			name: "second rune mismatch => no match",
			args: args{
				first:  runes.Rune('a'),
				second: runes.Rune('b'),
				input:  strings.NewReader("aa"),
			},
			wantRemain: "aa",
			wantErr:    errors.ErrNotMatched,
		},
		{
			name: "second rune match => first match",
			args: args{
				first:  runes.Rune('a'),
				second: runes.Rune('b'),
				input:  strings.NewReader("ab"),
			},
			wantMatch: 'a',
		},
		{
			name: "second rune match unicode => first match",
			args: args{
				first:  runes.Rune('a'),
				second: runes.Rune('😀'),
				input:  strings.NewReader("a😀😀"),
			},
			wantMatch:  'a',
			wantRemain: "😀",
		},
		{
			name: "first rune match unicode => first match",
			args: args{
				first:  runes.Rune('😀'),
				second: runes.Rune('a'),
				input:  strings.NewReader("😀a"),
			},
			wantMatch: '😀',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Terminated(tt.args.first, tt.args.second)
			s, err := p.Parse(tt.args.input)

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
