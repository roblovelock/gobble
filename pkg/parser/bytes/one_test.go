package bytes_test

import (
	"bytes"
	"github.com/roblovelock/gobble/pkg/parser"
	gobble "github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

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
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}
