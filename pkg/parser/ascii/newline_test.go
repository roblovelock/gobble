package ascii_test

import (
	"github.com/roblovelock/gobble/pkg/errors"
	"github.com/roblovelock/gobble/pkg/parser"
	"github.com/roblovelock/gobble/pkg/parser/ascii"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestCRLF(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []byte
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "missing linefeed => EOF",
			args:       args{input: strings.NewReader("\r")},
			wantErr:    io.EOF,
			wantRemain: []byte{'\r'},
		},
		{
			name:       "no match => not matched",
			args:       args{input: strings.NewReader("\r\r")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'\r', '\r'},
		},
		{
			name:       "match => match",
			args:       args{input: strings.NewReader("\r\n")},
			wantMatch:  []byte{'\r', '\n'},
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.CRLF()
			s, err := p(tt.args.input)

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

func TestNewline(t *testing.T) {
	type args struct {
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
			name:       "empty input => EOF",
			args:       args{input: strings.NewReader("")},
			wantErr:    io.EOF,
			wantRemain: []byte{},
		},
		{
			name:       "no match => not matched",
			args:       args{input: strings.NewReader("\r")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'\r'},
		},
		{
			name:       "match => match",
			args:       args{input: strings.NewReader("\n")},
			wantMatch:  '\n',
			wantRemain: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Newline()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.ErrorIs(t, err, tt.wantErr)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}

func TestLineEnding(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []byte
		wantRemain []byte
		wantErr    error
	}{
		{
			name:    "empty input => EOF",
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "missing linefeed => EOF",
			args:       args{input: strings.NewReader("\r")},
			wantErr:    io.EOF,
			wantRemain: []byte{'\r'},
		},
		{
			name:       "no match => not matched",
			args:       args{input: strings.NewReader("\r\r")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'\r', '\r'},
		},
		{
			name:       "match CRLF => match",
			args:       args{input: strings.NewReader("\r\n")},
			wantMatch:  []byte{'\r', '\n'},
			wantRemain: []byte{},
		},
		{
			name:       "match LF => match",
			args:       args{input: strings.NewReader("\n\n")},
			wantMatch:  []byte{'\n'},
			wantRemain: []byte{'\n'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.LineEnding()
			s, err := p(tt.args.input)

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
