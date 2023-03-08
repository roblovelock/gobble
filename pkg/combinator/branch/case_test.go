package branch

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
	"unicode"
)

func TestCase(t *testing.T) {
	type args struct {
		switchParser parser.Parser[parser.Reader, rune]
		caseParsers  map[rune]parser.Parser[parser.Reader, rune]
		input        parser.Reader
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
				switchParser: runes.One(),
				caseParsers:  map[rune]parser.Parser[parser.Reader, rune]{},
				input:        strings.NewReader(""),
			},
			wantRemain: "",
			wantErr:    io.EOF,
		},
		{
			name: "empty map => no match",
			args: args{
				switchParser: runes.One(),
				caseParsers:  map[rune]parser.Parser[parser.Reader, rune]{},
				input:        strings.NewReader("a"),
			},
			wantRemain: "a",
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "matched parser EOF => EOF",
			args: args{
				switchParser: runes.One(),
				caseParsers: map[rune]parser.Parser[parser.Reader, rune]{
					'a': runes.Rune('b'),
				},
				input: strings.NewReader("a"),
			},
			wantRemain: "a",
			wantErr:    io.EOF,
		},
		{
			name: "matched parser no match => no match",
			args: args{
				switchParser: runes.One(),
				caseParsers: map[rune]parser.Parser[parser.Reader, rune]{
					'a': runes.Rune('b'),
				},
				input: strings.NewReader("ac"),
			},
			wantRemain: "ac",
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "matched parser match => match",
			args: args{
				switchParser: runes.One(),
				caseParsers: map[rune]parser.Parser[parser.Reader, rune]{
					'a': runes.Rune('c'),
				},
				input: strings.NewReader("ac"),
			},
			wantMatch: 'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Case(tt.args.switchParser, tt.args.caseParsers)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))
		})
	}
}

func TestCaseOrDefault(t *testing.T) {
	type args struct {
		switchParser  parser.Parser[parser.Reader, rune]
		caseParsers   map[rune]parser.Parser[parser.Reader, rune]
		defaultParser parser.Parser[parser.Reader, rune]
		input         parser.Reader
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
				switchParser:  runes.One(),
				caseParsers:   map[rune]parser.Parser[parser.Reader, rune]{},
				defaultParser: runes.One(),
				input:         strings.NewReader(""),
			},
			wantRemain: "",
			wantErr:    io.EOF,
		},
		{
			name: "matched parser EOF => EOF",
			args: args{
				switchParser: runes.One(),
				caseParsers: map[rune]parser.Parser[parser.Reader, rune]{
					'a': runes.Rune('b'),
				},
				defaultParser: runes.One(),
				input:         strings.NewReader("a"),
			},
			wantRemain: "a",
			wantErr:    io.EOF,
		},
		{
			name: "matched parser no match => no match",
			args: args{
				switchParser: runes.One(),
				caseParsers: map[rune]parser.Parser[parser.Reader, rune]{
					'a': runes.Rune('b'),
				},
				defaultParser: runes.One(),
				input:         strings.NewReader("ac"),
			},
			wantRemain: "ac",
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "default parser no match => no match",
			args: args{
				switchParser: runes.One(),
				caseParsers: map[rune]parser.Parser[parser.Reader, rune]{
					'b': runes.Rune('c'),
				},
				defaultParser: runes.Rune('b'),
				input:         strings.NewReader("ac"),
			},
			wantRemain: "ac",
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "matched parser match => match",
			args: args{
				switchParser: runes.One(),
				caseParsers: map[rune]parser.Parser[parser.Reader, rune]{
					'a': runes.Rune('c'),
				},
				defaultParser: runes.Rune('b'),
				input:         strings.NewReader("ac"),
			},
			wantMatch: 'c',
		},
		{
			name: "default parser match => match",
			args: args{
				switchParser:  runes.One(),
				caseParsers:   map[rune]parser.Parser[parser.Reader, rune]{},
				defaultParser: runes.One(),
				input:         strings.NewReader("ab"),
			},
			wantMatch: 'b',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CaseOrDefault(tt.args.switchParser, tt.args.caseParsers, tt.args.defaultParser)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))
		})
	}
}

func TestPeekCase(t *testing.T) {
	type args struct {
		caseParsers map[byte]parser.Parser[parser.Reader, string]
		input       parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  string
		wantRemain string
		wantErr    error
	}{
		{
			name: "empty input => EOF",
			args: args{
				caseParsers: map[byte]parser.Parser[parser.Reader, string]{},
				input:       strings.NewReader(""),
			},
			wantRemain: "",
			wantErr:    io.EOF,
		},
		{
			name: "empty map => no match",
			args: args{
				caseParsers: map[byte]parser.Parser[parser.Reader, string]{},
				input:       strings.NewReader("a"),
			},
			wantRemain: "a",
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "matched parser no match => no match",
			args: args{
				caseParsers: map[byte]parser.Parser[parser.Reader, string]{
					'a': runes.TakeWhile1(unicode.IsDigit),
				},
				input: strings.NewReader("ab"),
			},
			wantRemain: "ab",
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "matched parser match => match",
			args: args{
				caseParsers: map[byte]parser.Parser[parser.Reader, string]{
					'a': runes.Take(2),
				},
				input: strings.NewReader("ab"),
			},
			wantMatch: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PeekCase(tt.args.caseParsers)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))
		})
	}
}
