package modifier_test

import (
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestValue(t *testing.T) {
	type args struct {
		parser parser.Parser[parser.Reader, byte]
		value  string
		input  parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  string
		wantErr    error
		wantRemain []byte
	}{
		{
			name: "empty input => EOF",
			args: args{
				parser: ascii.Digit(),
				value:  "match",
				input:  strings.NewReader(""),
			},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name: "non digit => no match",
			args: args{
				parser: ascii.Digit(),
				value:  "match",
				input:  strings.NewReader("a"),
			},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name: "digit => match",
			args: args{
				parser: ascii.Digit(),
				value:  "match",
				input:  strings.NewReader("9"),
			},
			wantMatch:  "match",
			wantRemain: []byte{},
		},
		{
			name: "digit => match with remainder",
			args: args{
				parser: ascii.Digit(),
				value:  "match",
				input:  strings.NewReader("12"),
			},
			wantMatch:  "match",
			wantRemain: []byte{'2'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := modifier.Value(tt.args.parser, tt.args.value)
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}
