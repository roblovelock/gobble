package stream_test

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/stream"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestEOF(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantRemain string
		wantErr    error
	}{
		{
			name: "empty input => EOF",
			args: args{input: strings.NewReader("")},
		},
		{
			name:       "not at end => no match",
			args:       args{input: strings.NewReader("a")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := stream.EOF()
			s, err := p.Parse(tt.args.input)

			assert.Nil(t, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, string(remain))

		})
	}
}
