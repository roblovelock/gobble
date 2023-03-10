package bytes_test

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestByte(t *testing.T) {
	type args struct {
		byte  byte
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  byte
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{byte: 'a', input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "byte mismatch => no match",
			args:       args{byte: 'a', input: strings.NewReader("b")},
			wantRemain: []byte{'b'},
			wantErr:    errors.ErrNotMatched,
		},
		{
			name:       "byte match => match",
			args:       args{byte: 'a', input: strings.NewReader("a")},
			wantMatch:  'a',
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bytes.Byte(tt.args.byte)
			s, err := p.Parse(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			if err != io.EOF {
				remain, err := io.ReadAll(tt.args.input)
				require.NoError(t, err)
				assert.Equal(t, tt.wantRemain, remain)
			}
		})
	}
}
