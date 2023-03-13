package runes_test

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

func TestRune(t *testing.T) {
	type args struct {
		rune  rune
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  rune
		wantRemain string
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{rune: 'a', input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "rune mismatch => no match",
			args:       args{rune: 'a', input: strings.NewReader("b")},
			wantRemain: "b",
			wantErr:    errors.ErrNotMatched,
		},
		{
			name:      "rune match => match",
			args:      args{rune: 'a', input: strings.NewReader("a")},
			wantMatch: 'a',
		},
		{
			name:      "rune match unicode => match",
			args:      args{rune: '😀', input: strings.NewReader("😀")},
			wantMatch: '😀',
		},
		{
			name:       "rune match unicode => match one",
			args:       args{rune: '😀', input: strings.NewReader("😀😀")},
			wantMatch:  '😀',
			wantRemain: "😀",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := runes.Rune(tt.args.rune)
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
