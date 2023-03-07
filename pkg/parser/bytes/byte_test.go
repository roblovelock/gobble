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

func ExampleByte_match() {
	input := strings.NewReader("abc")
	byteParser := bytes.Byte('a')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'a', Error: <nil>, Remainder: 'bc'
}

func ExampleByte_noMatch() {
	input := strings.NewReader("abc")
	byteParser := bytes.Byte('b')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'not matched', Remainder: 'abc'
}

func ExampleByte_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.Byte('a')

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: 0, Error: 'EOF', Remainder: ''
}

func TestByte(t *testing.T) {
	type args struct {
		byte  byte
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
			args:    args{byte: 'a', input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "byte mismatch => no match",
			args:       args{byte: 'a', input: strings.NewReader("b")},
			wantRemain: []byte{'b'},
			wantErr:    parser.ErrNotMatched,
		},
		{
			name:       "byte match => match",
			args:       args{byte: 'a', input: strings.NewReader("a")},
			wantMatch:  'a',
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.Byte(tt.args.byte)
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
