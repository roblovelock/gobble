package modifier_test

import (
	"github.com/roblovelock/gobble/pkg/combinator/modifier"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/runes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
	"unicode"
)

func TestNot(t *testing.T) {
	type args struct {
		parser parser.Parser[parser.Reader, string]
		input  parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantRemain string
		wantErr    error
	}{
		{
			name: "empty input => match",
			args: args{
				parser: runes.TakeWhile1(unicode.IsDigit),
				input:  strings.NewReader(""),
			},
			wantRemain: "",
		},
		{
			name: "matched parser => error",
			args: args{
				parser: runes.TakeWhile1(unicode.IsDigit),
				input:  strings.NewReader("12"),
			},
			wantRemain: "12",
			wantErr:    parser.ErrNotMatched,
		},
		{
			name: "not matched parser => match",
			args: args{
				parser: runes.TakeWhile1(unicode.IsDigit),
				input:  strings.NewReader("ab"),
			},
			wantRemain: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := modifier.Not(tt.args.parser)
			s, err := p(tt.args.input)

			assert.Nil(t, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))
		})
	}
}
