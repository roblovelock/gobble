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

func TestAlpha(t *testing.T) {
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
			name:    "empty input => EOF",
			args:    args{input: strings.NewReader("")},
			wantErr: io.EOF,
		},
		{
			name:       "non alpha => no match",
			args:       args{input: strings.NewReader("8")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'8'},
		},
		{
			name:       "alpha => match",
			args:       args{input: strings.NewReader("a")},
			wantMatch:  'a',
			wantRemain: []byte{},
		},
		{
			name:       "alpha => match with remainder",
			args:       args{input: strings.NewReader("ab")},
			wantMatch:  'a',
			wantRemain: []byte{'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Alpha()
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

func TestAlpha1(t *testing.T) {
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
			name:       "non alpha => no match",
			args:       args{input: strings.NewReader("8")},
			wantErr:    errors.ErrNotMatched,
			wantRemain: []byte{'8'},
		},
		{
			name:       "alpha => match",
			args:       args{input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "alpha => match with remainder",
			args:       args{input: strings.NewReader("ab8")},
			wantMatch:  []byte{'a', 'b'},
			wantRemain: []byte{'8'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Alpha1()
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

func TestAlpha0(t *testing.T) {
	type args struct {
		input parser.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantMatch  []byte
		wantRemain []byte
	}{
		{
			name:       "empty input => EOF",
			args:       args{input: strings.NewReader("")},
			wantMatch:  []byte{},
			wantRemain: []byte{},
		},
		{
			name:       "non alpha => no match",
			args:       args{input: strings.NewReader("8")},
			wantMatch:  []byte{},
			wantRemain: []byte{'8'},
		},
		{
			name:       "alpha => match",
			args:       args{input: strings.NewReader("a")},
			wantMatch:  []byte{'a'},
			wantRemain: []byte{},
		},
		{
			name:       "alpha => match with remainder",
			args:       args{input: strings.NewReader("ab8")},
			wantMatch:  []byte{'a', 'b'},
			wantRemain: []byte{'8'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ascii.Alpha0()
			s, err := p(tt.args.input)

			assert.Equal(t, tt.wantMatch, s)
			assert.NoError(t, err)

			remain, err := io.ReadAll(tt.args.input)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRemain, remain)
		})
	}
}
