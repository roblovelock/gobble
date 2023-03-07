package ascii_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func ExampleDigit_match() {
	input := strings.NewReader("123")
	byteParser := ascii.Digit()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '1', Error: <nil>, Remainder: '23'
}

func ExampleDigit_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Digit()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleDigit_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Digit()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func TestDigit(t *testing.T) {
	type args struct {
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
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  '9',
			wantRemain: []byte{},
		},
		{
			name:       "digit => match with remainder",
			args:       args{input: strings.NewReader("12")},
			wantMatch:  '1',
			wantRemain: []byte{'2'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Digit()
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

func ExampleDigit1_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.Digit1()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '123', Error: <nil>, Remainder: 'abc'
}

func ExampleDigit1_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Digit1()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleDigit1_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Digit1()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func TestDigit1(t *testing.T) {
	type args struct {
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
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    parser.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  []byte{'9'},
			wantRemain: []byte{},
		},
		{
			name:       "digit => match with remainder",
			args:       args{input: strings.NewReader("12a")},
			wantMatch:  []byte{'1', '2'},
			wantRemain: []byte{'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Digit1()
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

func ExampleDigit0_match() {
	input := strings.NewReader("123abc")
	byteParser := ascii.Digit0()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '123', Error: <nil>, Remainder: 'abc'
}

func ExampleDigit0_noMatch() {
	input := strings.NewReader("abc")
	byteParser := ascii.Digit0()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: 'abc'
}

func ExampleDigit0_endOfFile() {
	input := strings.NewReader("")
	byteParser := ascii.Digit0()

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: <nil>, Remainder: ''
}

func TestDigit0(t *testing.T) {
	type args struct {
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
			args:       args{input: strings.NewReader("")},
			wantMatch:  []byte{},
			wantRemain: []byte{},
		},
		{
			name:       "non digit => no match",
			args:       args{input: strings.NewReader("a")},
			wantMatch:  []byte{},
			wantRemain: []byte{'a'},
		},
		{
			name:       "digit => match",
			args:       args{input: strings.NewReader("9")},
			wantMatch:  []byte{'9'},
			wantRemain: []byte{},
		},
		{
			name:       "digit => match with remainder",
			args:       args{input: strings.NewReader("12a")},
			wantMatch:  []byte{'1', '2'},
			wantRemain: []byte{'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Digit0()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}
