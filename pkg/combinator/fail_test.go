package combinator

import (
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestFail(t *testing.T) {
	type args struct {
		err   error
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantErr    error
		wantRemain string
	}{
		{
			name: "empty input => success",
			args: args{
				err:   parser.ErrNotMatched,
				input: strings.NewReader(""),
			},
			wantErr:    parser.ErrNotMatched,
			wantRemain: "",
		},
		{
			name: "input => success",
			args: args{
				err:   parser.ErrNotMatched,
				input: strings.NewReader("ab"),
			},
			wantErr:    parser.ErrNotMatched,
			wantRemain: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Fail[parser.Reader, parser.Empty](tt.args.err)
			s, err := p(tt.args.input)

			assert.Nil(t, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))
		})
	}
}
