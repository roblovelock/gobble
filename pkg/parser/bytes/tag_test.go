package bytes_test

import (
	"fmt"
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func ExampleTag_match() {
	input := strings.NewReader("abc")
	byteParser := bytes.Tag([]byte("ab"))

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'ab', Error: <nil>, Remainder: 'c'
}

func ExampleTag_noMatch() {
	input := strings.NewReader("abc")
	byteParser := bytes.Tag([]byte("bc"))

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'not matched', Remainder: 'abc'
}

func ExampleTag_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.Tag([]byte("ab"))

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func TestTag(t *testing.T) {
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
			name:       "empty input => EOF",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("")},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "short input => EOF",
			args:       args{bytes: []byte{'a', 'b'}, input: strings.NewReader("a")},
			wantErr:    io.EOF,
			wantRemain: []byte{'a'},
		},
		{
			name:       "byte mismatch => no match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("b")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'b'},
		},
		{
			name:       "bytes mismatch => no match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("cba")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'c', 'b', 'a'},
		},
		{
			name:       "byte match => match",
			args:       args{bytes: []byte{'a'}, input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "bytes match => match",
			args:       args{bytes: []byte{'a', 'b', 'c'}, input: strings.NewReader("abcd")},
			wantMatch:  []byte{'a', 'b', 'c'},
			wantRemain: []byte{'d'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.Tag(tt.args.bytes)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}
