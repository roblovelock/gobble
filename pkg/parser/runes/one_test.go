package runes_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gobble/pkg/parser"
	"gobble/pkg/parser/runes"
	"io"
	"strings"
	"testing"
)

func ExampleOne_match() {
	input := strings.NewReader("𒀀𒀀")
	numericParser := runes.One()

	match, err := numericParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", string(match), err, remainder)

	// Output:
	// Match: '𒀀', Error: <nil>, Remainder: '𒀀'
}

func ExampleOne_endOfFile() {
	input := bytes.NewReader([]byte{})
	numericParser := runes.One()

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
		wantMatch  rune
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: bytes.NewReader([]byte{})},
			wantErr: io.EOF,
		},
		{
			name:       "rune => match",
			args:       args{input: strings.NewReader("𒀀𒀀")},
			wantMatch:  '𒀀',
			wantRemain: []byte("𒀀"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := runes.One()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))
		})
	}
}
