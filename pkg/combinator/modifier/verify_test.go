package modifier_test

import (
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestVerify(t *testing.T) {
	type args struct {
		parser parser.Parser[parser.Reader, byte]
		input  parser.Reader
		verify parser.Predicate[byte]
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  byte
		wantRemain []byte
		wantErr    error
	}{
		{
			name: "empty input => EOF",
			args: args{
				parser: ascii.Digit(),
				verify: ascii.IsDigit,
				input:  strings.NewReader(""),
			},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name: "paser fail => no match",
			args: args{
				parser: ascii.Digit(),
				verify: ascii.IsDigit,
				input:  strings.NewReader("a"),
			},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name: "verify fail => no match",
			args: args{
				parser: bytes.One(),
				verify: ascii.IsDigit,
				input:  strings.NewReader("a"),
			},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'a'},
		},
		{
			name: "digit => match",
			args: args{
				parser: ascii.Digit(),
				verify: ascii.IsDigit,
				input:  strings.NewReader("9"),
			},
			wantMatch:  '9',
			wantRemain: []byte{},
		},
		{
			name: "digit => match with remainder",
			args: args{
				parser: ascii.Digit(),
				verify: ascii.IsDigit,
				input:  strings.NewReader("12"),
			},
			wantMatch:  '1',
			wantRemain: []byte{'2'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := modifier.Verify(tt.args.parser, tt.args.verify)
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)

		})
	}
}
