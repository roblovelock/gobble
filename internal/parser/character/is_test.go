package character

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/internal/parser"
	"io"
	"strings"
	"testing"
)

func TestIsA(t *testing.T) {
	type args struct {
		runes []rune
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  string
		wantRemain string
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{runes: []rune{'a'}, input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "empty runes => no match",
			args:       args{runes: []rune{}, input: strings.NewReader("a")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: "a",
		},
		{
			name:       "rune mismatch => no match",
			args:       args{runes: []rune{'a'}, input: strings.NewReader("b")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: "b",
		},
		{
			name:       "runes mismatch => no match",
			args:       args{runes: []rune{'a', 'b', 'c'}, input: strings.NewReader("d")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: "d",
		},
		{
			name:      "rune match => match",
			args:      args{runes: []rune{'a'}, input: strings.NewReader("a")},
			wantMatch: "a",
		},
		{
			name:      "runes match => match",
			args:      args{runes: []rune{'a', 'b', 'c'}, input: strings.NewReader("b")},
			wantMatch: "b",
		},
		{
			name:      "runes match unicode => match",
			args:      args{runes: []rune{'😀'}, input: strings.NewReader("😀")},
			wantMatch: "😀",
		},
		{
			name:      "rune match many => match",
			args:      args{runes: []rune{'a'}, input: strings.NewReader("aaaa")},
			wantMatch: "aaaa",
		},
		{
			name:       "runes match many => match",
			args:       args{runes: []rune{'a', 'b', 'c'}, input: strings.NewReader("cbad")},
			wantMatch:  "cba",
			wantRemain: "d",
		},
		{
			name:      "runes match unicode many => match",
			args:      args{runes: []rune{'😀'}, input: strings.NewReader("😀😀")},
			wantMatch: "😀😀",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := IsA(tt.args.runes...)
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
