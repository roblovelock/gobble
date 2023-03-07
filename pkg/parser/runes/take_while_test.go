package runes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
	"unicode"
)

func ExampleTakeWhileMinMax_match() {
	input := strings.NewReader("abc")
	runeParser := runes.TakeWhileMinMax(1, 2, unicode.IsLetter)

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 'ab', Error: <nil>, Remainder: 'c'
}

func ExampleTakeWhileMinMax_noMatch() {
	input := strings.NewReader("abc")
	runeParser := runes.TakeWhileMinMax(1, 2, unicode.IsDigit)

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleTakeWhileMinMax_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.TakeWhileMinMax(1, 2, unicode.IsDigit)

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func TestTakeWhileMinMax(t *testing.T) {
	type args struct {
		input     parser.Reader
		predicate parser.Predicate[rune]
		max       int
		min       int
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  string
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input min 1 => EOF",
			args:       args{min: 1, max: 10, predicate: unicode.IsDigit, input: strings.NewReader("")},
			wantRemain: []byte{},
			wantErr:    io.EOF,
		},
		{
			name:       "empty input min 0 => empty",
			args:       args{min: 0, max: 10, predicate: unicode.IsDigit, input: strings.NewReader("")},
			wantMatch:  "",
			wantRemain: []byte{},
		},
		{
			name:       "take min 1 no match => empty",
			args:       args{min: 1, max: 10, predicate: unicode.IsDigit, input: strings.NewReader("b")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "take min 0 no match => empty",
			args:       args{min: 0, max: 10, predicate: unicode.IsDigit, input: strings.NewReader("b")},
			wantMatch:  "",
			wantRemain: []byte{'b'},
		},
		{
			name:       "take max 1 match => match 1",
			args:       args{min: 0, max: 1, predicate: unicode.IsDigit, input: strings.NewReader("12b")},
			wantMatch:  "1",
			wantRemain: []byte{'2', 'b'},
		},
		{
			name:       "take max 10 match => match all",
			args:       args{min: 0, max: 10, predicate: unicode.IsDigit, input: strings.NewReader("12b")},
			wantMatch:  "12",
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := runes.TakeWhileMinMax(tt.args.min, tt.args.max, tt.args.predicate)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleTakeWhile1_match() {
	input := strings.NewReader("abc123")
	runeParser := runes.TakeWhile1(unicode.IsLetter)

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleTakeWhile1_noMatch() {
	input := strings.NewReader("abc")
	runeParser := runes.TakeWhile1(unicode.IsDigit)

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleTakeWhile1_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.TakeWhile1(unicode.IsDigit)

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func TestTakeWhile1(t *testing.T) {
	type args struct {
		input     parser.Reader
		predicate parser.Predicate[rune]
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  string
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{predicate: unicode.IsDigit, input: strings.NewReader("")},
			wantRemain: []byte{},
			wantErr:    io.EOF,
		},
		{
			name:       "take no match => empty",
			args:       args{predicate: unicode.IsDigit, input: strings.NewReader("b")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "take match => match all",
			args:       args{predicate: unicode.IsDigit, input: strings.NewReader("12b")},
			wantMatch:  "12",
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := runes.TakeWhile1(tt.args.predicate)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func ExampleTakeWhile_match() {
	input := strings.NewReader("abc123")
	runeParser := runes.TakeWhile(unicode.IsLetter)

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleTakeWhile_noMatch() {
	input := strings.NewReader("abc")
	runeParser := runes.TakeWhile(unicode.IsDigit)

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: 'abc'
}

func ExampleTakeWhile_endOfFile() {
	input := strings.NewReader("")
	runeParser := runes.TakeWhile(unicode.IsDigit)

	match, err := runeParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}

func TestTakeWhile(t *testing.T) {
	type args struct {
		input     parser.Reader
		predicate parser.Predicate[rune]
		max       int
		min       int
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  string
		wantRemain []byte
	}{
		{
			name:       "empty input => empty",
			args:       args{min: 1, max: 10, predicate: unicode.IsDigit, input: strings.NewReader("")},
			wantRemain: []byte{},
			wantMatch:  "",
		},
		{
			name:       "take no match => empty",
			args:       args{min: 0, max: 10, predicate: unicode.IsDigit, input: strings.NewReader("b")},
			wantMatch:  "",
			wantRemain: []byte{'b'},
		},
		{
			name:       "take match => match all",
			args:       args{min: 0, max: 10, predicate: unicode.IsDigit, input: strings.NewReader("12b")},
			wantMatch:  "12",
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := runes.TakeWhile(tt.args.predicate)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}
