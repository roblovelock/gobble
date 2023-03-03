package bytes_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/pkg/parser"
	gobble "gobble/pkg/parser/bytes"
	"io"
	"testing"
)

func ExampleOne_match() {
	input := bytes.NewReader([]byte{1, 2, 3})
	numericParser := gobble.One()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: %v, Remainder: %v", match, err, remainder)

	// Output:
	// Match: 1, Error: <nil>, Remainder: [2 3]
}

func ExampleOne_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := gobble.One()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: %d, Error: '%v', Remainder: %v", match, err, remainder)

	// Output:
	// Match: 0, Error: 'EOF', Remainder: []
}

func TestOne(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  uint8
		wantRemain []byte
		wantErr    error
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: bytes.NewReader([]byte{})},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "byte => match",
			args:       args{input: bytes.NewReader([]byte{1, 2, 3})},
			wantMatch:  1,
			wantRemain: []byte{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := gobble.One()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}
