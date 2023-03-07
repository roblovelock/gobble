package sequence

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestDelimited(t *testing.T) {
	type args struct {
		first  parser.Parser[parser.Reader, rune]
		second parser.Parser[parser.Reader, rune]
		third  parser.Parser[parser.Reader, rune]
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
				first:  runes.Rune('"'),
				second: runes.Rune('b'),
				third:  runes.Rune('"'),
				input:  strings.NewReader(""),
			},
			wantErr: io.EOF,
		},
		{
			name: "second rune EOF => EOF",
			args: args{
				first:  runes.Rune('"'),
				second: runes.Rune('b'),
				third:  runes.Rune('"'),
				input:  strings.NewReader(`"`),
			},
			wantRemain: "a",
			wantErr:    io.EOF,
		},
		{
			name: "third rune EOF => EOF",
			args: args{
				first:  runes.Rune('"'),
				second: runes.Rune('a'),
				third:  runes.Rune('"'),
				input:  strings.NewReader(`"a`),
			},
			wantRemain: `"a`,
			wantErr:    io.EOF,
		},
		{
			name: "first rune mismatch => no match",
			args: args{
				first:  runes.Rune('"'),
				second: runes.Rune('b'),
				third:  runes.Rune('"'),
				input:  strings.NewReader(`bb"`),
			},
			wantRemain: `bb"`,
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "second rune mismatch => no match",
			args: args{
				first:  runes.Rune('"'),
				second: runes.Rune('b'),
				third:  runes.Rune('"'),
				input:  strings.NewReader(`"a"`),
			},
			wantRemain: `"a"`,
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "second rune match => second match",
			args: args{
				first:  runes.Rune('"'),
				second: runes.Rune('a'),
				third:  runes.Rune('"'),
				input:  strings.NewReader(`"a"`),
			},
			wantMatch: 'a',
		},
		{
			name: "second rune match unicode => second match",
			args: args{
				first:  runes.Rune('a'),
				second: runes.Rune('😀'),
				third:  runes.Rune('😀'),
				input:  strings.NewReader("a😀😀😀"),
			},
			wantMatch:  '😀',
			wantRemain: "😀",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Delimited(tt.args.first, tt.args.second, tt.args.third)
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
