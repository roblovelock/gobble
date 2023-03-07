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
)

func ExampleTake_match() {
	input := strings.NewReader("𒀀a𒀀")
	byteParser := runes.Take(2)

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: %v, Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '𒀀a', Error: <nil>, Remainder: '𒀀'
}

func ExampleTake_unexpectedEndOfFile() {
	input := strings.NewReader("𒀀a𒀀")
	byteParser := runes.Take(4)

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: '𒀀a𒀀'
}

func ExampleTake_endOfFile() {
	input := strings.NewReader("")
	byteParser := runes.Take(4)

	match, err := byteParser(input)
	remainder, _ := io.ReadAll(input)
	fmt.Printf("Match: '%s', Error: '%v', Remainder: '%s'", match, err, string(remainder))

	// Output:
	// Match: '', Error: 'EOF', Remainder: ''
}

func TestTake(t *testing.T) {
	type args struct {
		take  int
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
			args:    args{take: 1, input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "take 0 => empty",
			args:       args{take: 0, input: strings.NewReader("b")},
			wantMatch:  "",
			wantRemain: "b",
		},
		{
			name:      "take 1 => match",
			args:      args{take: 1, input: strings.NewReader("a")},
			wantMatch: "a",
		},
		{
			name:      "take 1 unicode => match",
			args:      args{take: 1, input: strings.NewReader("😀")},
			wantMatch: "😀",
		},
		{
			name:       "take many unicode => match many",
			args:       args{take: 5, input: strings.NewReader("1234😀😀")},
			wantMatch:  "1234😀",
			wantRemain: "😀",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := runes.Take(tt.args.take)
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
