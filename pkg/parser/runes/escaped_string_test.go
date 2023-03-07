package runes_test

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestEscapedString(t *testing.T) {
	type args struct {
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
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "no opening quote => no match",
			args:       args{input: strings.NewReader(`a"`)},
			wantErr:    parser.ErrNotMatched,
			wantRemain: `a"`,
		},
		{
			name:       "invalid escape character => no match",
			args:       args{input: strings.NewReader(`"\g"`)},
			wantErr:    parser.ErrNotMatched,
			wantRemain: `"\g"`,
		},
		{
			name:       "no closing quote => EOF",
			args:       args{input: strings.NewReader(`"\u1234`)},
			wantErr:    io.EOF,
			wantRemain: `"\u1234`,
		},
		{
			name:       "special character truncated => EOF",
			args:       args{input: strings.NewReader(`"\`)},
			wantErr:    io.EOF,
			wantRemain: `"\`,
		},
		{
			name:       "unicode truncation => EOF",
			args:       args{input: strings.NewReader(`"\u`)},
			wantErr:    io.EOF,
			wantRemain: `"\u`,
		},
		{
			name:       "unicode codepoint truncation => EOF",
			args:       args{input: strings.NewReader(`"\u123`)},
			wantErr:    io.EOF,
			wantRemain: `"\u123`,
		},
		{
			name:       "unicode surrogate truncation => eof",
			args:       args{input: strings.NewReader(`"\ud83e`)},
			wantErr:    io.EOF,
			wantRemain: `"\ud83e`,
		},
		{
			name:       "unicode surrogate missing => eof",
			args:       args{input: strings.NewReader(`"\ud83e"`)},
			wantErr:    io.EOF,
			wantRemain: `"\ud83e"`,
		},
		{
			name:       "unicode surrogate invalid => no match",
			args:       args{input: strings.NewReader(`"\ud83e Some invalid text"`)},
			wantErr:    parser.ErrNotMatched,
			wantRemain: `"\ud83e Some invalid text"`,
		},
		{
			name:       "invalid unicode characters => no match",
			args:       args{input: strings.NewReader(`"\u123X"`)},
			wantErr:    parser.ErrNotMatched,
			wantRemain: `"\u123X"`,
		},
		{
			name:       "invalid unicode surrogate characters => no match",
			args:       args{input: strings.NewReader(`"\ud83e\udd2X"`)},
			wantErr:    parser.ErrNotMatched,
			wantRemain: `"\ud83e\udd2X"`,
		},
		{
			name:      "empty string match => match",
			args:      args{input: strings.NewReader(`""`)},
			wantMatch: "",
		},
		{
			name:       "string match => match",
			args:       args{input: strings.NewReader(`"this is a string" this isn't`)},
			wantMatch:  "this is a string",
			wantRemain: " this isn't",
		},
		{
			name:       "escaped character string match => match",
			args:       args{input: strings.NewReader(`"\b\f\r\n\t\\\'\/"`)},
			wantMatch:  "\b\f\r\n\t\\'/",
			wantRemain: "",
		},
		{
			name:      "unicode characters => match",
			args:      args{input: strings.NewReader(`"😀"`)},
			wantMatch: "😀",
		},
		{
			name:      "escaped unicode characters => match",
			args:      args{input: strings.NewReader(`"\u002F\u002f\ud83e\udd2d"`)},
			wantMatch: "//🤭",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := runes.EscapedString()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))

		})
	}
}
