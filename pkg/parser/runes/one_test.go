package runes_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/pkg/parser"
	"gobble/pkg/parser/runes"
	"io"
	"strings"
	"testing"
)

func ExampleOneOf_match() {
	input := strings.NewReader("𒀀a𒀀")
	runeParser := runes.OneOf('𒀀', 'a')

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '𒀀', Error: <nil>, Remainder: 'a𒀀'
}

func ExampleOneOf_noMatch() {
	input := strings.NewReader("123")
	runeParser := runes.OneOf('𒀀', 'a')

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: '123'
}

func ExampleOneOf_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.OneOf('𒀀', 'a')

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func TestOneOf(t *testing.T) {
	type args struct {
		runes []rune
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
			wantMatch: 'a',
		},
		{
			name:      "runes match => match",
			args:      args{runes: []rune{'a', 'b', 'c'}, input: strings.NewReader("b")},
			wantMatch: 'b',
		},
		{
			name:      "rune match unicode => match",
			args:      args{runes: []rune{'😀'}, input: strings.NewReader("😀")},
			wantMatch: '😀',
		},
		{
			name:       "runes match unicode => match",
			args:       args{runes: []rune{'😀', 'a'}, input: strings.NewReader("😀ab")},
			wantMatch:  '😀',
			wantRemain: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := runes.OneOf(tt.args.runes...)
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

func ExampleOneOf1_match() {
	input := strings.NewReader("𒀀a𒀀123")
	runeParser := runes.OneOf1('𒀀', 'a')

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '𒀀a𒀀', Error: <nil>, Remainder: '123'
}

func ExampleOneOf1_noMatch() {
	input := strings.NewReader("123")
	runeParser := runes.OneOf1('𒀀', 'a')

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: '123'
}

func ExampleOneOf1_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.OneOf1('𒀀', 'a')

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func TestOneOf1(t *testing.T) {
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
			p := runes.OneOf1(tt.args.runes...)
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

func ExampleOneOf0_match() {
	input := strings.NewReader("𒀀a𒀀123")
	runeParser := runes.OneOf0('𒀀', 'a')

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '𒀀a𒀀', Error: <nil>, Remainder: '123'
}

func ExampleOneOf0_noMatch() {
	input := strings.NewReader("123")
	runeParser := runes.OneOf0('𒀀', 'a')

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: '123'
}

func ExampleOneOf0_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.OneOf0('𒀀', 'a')

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}

func TestOneOf0(t *testing.T) {
	type args struct {
		runes []rune
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  string
		wantRemain string
	}{
		{
			name: "empty input => EOF",
			args: args{runes: []rune{'a'}, input: strings.NewReader("")},
		},
		{
			name:       "empty runes => no match",
			args:       args{runes: []rune{}, input: strings.NewReader("a")},
			wantRemain: "a",
		},
		{
			name:       "rune mismatch => no match",
			args:       args{runes: []rune{'a'}, input: strings.NewReader("b")},
			wantRemain: "b",
		},
		{
			name:       "runes mismatch => no match",
			args:       args{runes: []rune{'a', 'b', 'c'}, input: strings.NewReader("d")},
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
			p := runes.OneOf0(tt.args.runes...)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)
			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))
		})
	}
}
