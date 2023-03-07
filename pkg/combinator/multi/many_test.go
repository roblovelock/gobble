package multi

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestMany0(t *testing.T) {
	type args struct {
		p     parser.Parser[parser.Reader, rune]
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []rune
		wantRemain string
		wantErr    error
	}{
		{
			name:      "empty input => EOF",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("")},
			wantMatch: []rune{},
		},
		{
			name:       "rune mismatch => no match",
			args:       args{p: runes.Rune('a'), input: strings.NewReader("b")},
			wantMatch:  []rune{},
			wantRemain: "b",
		},
		{
			name:      "rune match => match",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("a")},
			wantMatch: []rune{'a'},
		},
		{
			name:      "rune match many => match",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("aaaa")},
			wantMatch: []rune{'a', 'a', 'a', 'a'},
		},
		{
			name:       "runes match unicode many => match",
			args:       args{p: runes.Rune('😀'), input: strings.NewReader("😀😀a")},
			wantMatch:  []rune{'😀', '😀'},
			wantRemain: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Many0(tt.args.p)
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

func TestMany1(t *testing.T) {
	type args struct {
		p     parser.Parser[parser.Reader, rune]
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []rune
		wantRemain string
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{p: runes.Rune('a'), input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "rune mismatch => no match",
			args:       args{p: runes.Rune('a'), input: strings.NewReader("b")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: "b",
		},
		{
			name:      "rune match => match",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("a")},
			wantMatch: []rune{'a'},
		},
		{
			name:      "rune match many => match",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("aaaa")},
			wantMatch: []rune{'a', 'a', 'a', 'a'},
		},
		{
			name:       "runes match unicode many => match",
			args:       args{p: runes.Rune('😀'), input: strings.NewReader("😀😀a")},
			wantMatch:  []rune{'😀', '😀'},
			wantRemain: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Many1(tt.args.p)
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

func TestMany0Count(t *testing.T) {
	type args struct {
		p     parser.Parser[parser.Reader, rune]
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint
		wantRemain string
		wantErr    error
	}{
		{
			name:      "empty input => zero",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("")},
			wantMatch: 0,
		},
		{
			name:       "rune mismatch => zero",
			args:       args{p: runes.Rune('a'), input: strings.NewReader("b")},
			wantMatch:  0,
			wantRemain: "b",
		},
		{
			name:      "rune match => count",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("a")},
			wantMatch: 1,
		},
		{
			name:      "rune match many => count",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("aaaa")},
			wantMatch: 4,
		},
		{
			name:       "runes match unicode many => count",
			args:       args{p: runes.Rune('😀'), input: strings.NewReader("😀😀a")},
			wantMatch:  2,
			wantRemain: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Many0Count(tt.args.p)
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

func TestMany1Count(t *testing.T) {
	type args struct {
		p     parser.Parser[parser.Reader, rune]
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint
		wantRemain string
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{p: runes.Rune('a'), input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "rune mismatch => no match",
			args:       args{p: runes.Rune('a'), input: strings.NewReader("b")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: "b",
		},
		{
			name:      "rune match => count",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("a")},
			wantMatch: 1,
		},
		{
			name:      "rune match many => count",
			args:      args{p: runes.Rune('a'), input: strings.NewReader("aaaa")},
			wantMatch: 4,
		},
		{
			name:       "runes match unicode many => count",
			args:       args{p: runes.Rune('😀'), input: strings.NewReader("😀😀a")},
			wantMatch:  2,
			wantRemain: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Many1Count(tt.args.p)
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
