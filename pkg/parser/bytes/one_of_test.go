package bytes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func ExampleOneOf_match() {
	input := strings.NewReader("abc123")
	byteParser := bytes.OneOf('a', 'b', 'c')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'a', Error: <nil>, Remainder: 'bc123'
}

func ExampleOneOf_noMatch() {
	input := strings.NewReader("123")
	byteParser := bytes.OneOf('a', 'b', 'c')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: '123'
}

func ExampleOneOf_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.OneOf('a', 'b', 'c')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func TestOneOf(t *testing.T) {
	type args struct {
		bytes []byte
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  byte
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{bytes: []byte{'a'}, input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "empty bytes => no match",
			args:       args{bytes: []byte{}, input: strings.NewReader("a")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "byte mismatch => no match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("b")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "bytes mismatch => no match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("d")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'d'},
		},
		{
			name:       "byte match => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("a")},
			wantMatch:  'a',
			wantRemain: []byte{},
		},
		{
			name:       "bytes match => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("bb")},
			wantMatch:  'b',
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.OneOf(tt.args.bytes...)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}

func ExampleOneOf1_match() {
	input := strings.NewReader("abc123")
	byteParser := bytes.OneOf1('a', 'b', 'c')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleOneOf1_noMatch() {
	input := strings.NewReader("123")
	byteParser := bytes.OneOf1('a', 'b', 'c')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: '123'
}

func ExampleOneOf1_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.OneOf1('a', 'b', 'c')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func TestOneOf1(t *testing.T) {
	type args struct {
		bytes []byte
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []byte
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{bytes: []byte{'a'}, input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "empty bytes => no match",
			args:       args{bytes: []byte{}, input: strings.NewReader("a")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "byte mismatch => no match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("b")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "bytes mismatch => no match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("d")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'d'},
		},
		{
			name:       "byte match => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "bytes match => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("b")},
			wantMatch:  []byte{'b'},
			wantRemain: []byte{},
		},
		{
			name:       "byte match many => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("aaaa")},
			wantMatch:  []byte{'a', 'a', 'a', 'a'},
			wantRemain: []byte{},
		},
		{
			name:       "bytes match many => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("cbad")},
			wantMatch:  []byte{'c', 'b', 'a'},
			wantRemain: []byte{'d'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.OneOf1(tt.args.bytes...)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}

func ExampleOneOf0_match() {
	input := strings.NewReader("abc123")
	byteParser := bytes.OneOf0('a', 'b', 'c')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'abc', Error: <nil>, Remainder: '123'
}

func ExampleOneOf0_noMatch() {
	input := strings.NewReader("123")
	byteParser := bytes.OneOf0('a', 'b', 'c')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: '123'
}

func ExampleOneOf0_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.OneOf0('a', 'b', 'c')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}

func TestOneOf0(t *testing.T) {
	type args struct {
		bytes []byte
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []byte
		wantRemain []byte
	}{
		{
			name:       "empty input => EOF",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("")},
			wantMatch:  []byte{},
			wantRemain: []byte{},
		},
		{
			name:       "empty bytes => no match",
			args:       args{bytes: []byte{}, input: strings.NewReader("a")},
			wantMatch:  []byte{},
			wantRemain: []byte{'a'},
		},
		{
			name:       "byte mismatch => no match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("b")},
			wantMatch:  []byte{},
			wantRemain: []byte{'b'},
		},
		{
			name:       "bytes mismatch => no match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("d")},
			wantMatch:  []byte{},
			wantRemain: []byte{'d'},
		},
		{
			name:       "byte match => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "bytes match => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("b")},
			wantMatch:  []byte{'b'},
			wantRemain: []byte{},
		},
		{
			name:       "byte match many => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("aaaa")},
			wantMatch:  []byte{'a', 'a', 'a', 'a'},
			wantRemain: []byte{},
		},
		{
			name:       "bytes match many => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("cbad")},
			wantMatch:  []byte{'c', 'b', 'a'},
			wantRemain: []byte{'d'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.OneOf0(tt.args.bytes...)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}
