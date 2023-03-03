package bytes_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/pkg/parser"
	"gobble/pkg/parser/bytes"
	"io"
	"strings"
	"testing"
)

func ExampleTake_match() {
	input := strings.NewReader("abc")
	byteParser := bytes.Take(2)

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: 'ab', Error: <nil>, Remainder: 'c'
}

func ExampleTake_unexpectedEndOfFile() {
	input := strings.NewReader("abc")
	byteParser := bytes.Take(4)

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: 'abc'
}

func ExampleTake_endOfFile() {
	input := strings.NewReader("")
	byteParser := bytes.Take(4)

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", string(match), err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func TestTake(t *testing.T) {
	type args struct {
		take  uint
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
			args:    args{take: 1, input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "take 0 => empty",
			args:       args{take: 0, input: strings.NewReader("b")},
			wantMatch:  []byte{},
			wantRemain: []byte{'b'},
		},
		{
			name:       "take 1 => match",
			args:       args{take: 1, input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "take many => match many",
			args:       args{take: 3, input: strings.NewReader("1234")},
			wantMatch:  []byte{'1', '2', '3'},
			wantRemain: []byte{'4'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.Take(tt.args.take)
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
