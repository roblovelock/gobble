package modifier_test

import (
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestOptional(t *testing.T) {
	type args struct {
		parser parser.Parser[parser.Reader, byte]
		input  parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  byte
		wantRemain []byte
	}{
		{
			name: "empty input => EOF",
			args: args{
				parser: ascii.Digit(),
				input:  strings.NewReader(""),
			},
			wantRemain: []byte{},
		},
		{
			name: "non digit => no match",
			args: args{
				parser: ascii.Digit(),
				input:  strings.NewReader("a"),
			},
			wantRemain: []byte{'a'},
		},
		{
			name: "digit => match",
			args: args{
				parser: ascii.Digit(),
				input:  strings.NewReader("9"),
			},
			wantMatch:  '9',
			wantRemain: []byte{},
		},
		{
			name: "digit => match with remainder",
			args: args{
				parser: ascii.Digit(),
				input:  strings.NewReader("12"),
			},
			wantMatch:  '1',
			wantRemain: []byte{'2'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := modifier.Optional(tt.args.parser)
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}
